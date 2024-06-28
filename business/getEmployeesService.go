package business

import (
	"data-manage/models"
	"data-manage/repositories"
	"data-manage/utils/mysql"
	"data-manage/utils/redis"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type GetEmployeesService struct {
	getEmployeesServiceRegistry repositories.EmployeeRepository
}

func NewGetEmployeesService(getEmployeesServiceRegistry repositories.EmployeeRepository) *GetEmployeesService {
	return &GetEmployeesService{
		getEmployeesServiceRegistry: getEmployeesServiceRegistry,
	}
}

func (service *GetEmployeesService) GetEmployees(ctx *gin.Context, request models.GetEmployeeDetails) (*models.Employee, error) {
	redisClient := redis.GetRedisClient()
	cachedEmployees, err := redisClient.Client.Get(ctx, request.Email).Result()
	if err == nil {
		var employee models.Employee
		err := json.Unmarshal([]byte(cachedEmployees), &employee)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling cached data: %v", err)
		}
		return &employee, nil
	}

	db := mysql.GetMySQLClient().GormDB
	var employee models.Employee
	condition := map[string]interface{}{
		"email": request.Email,
	}
	employeeDetails, err := service.getEmployeesServiceRegistry.ReadEmployeeWithCondition(ctx, db, employee, condition)
	if err != nil {
		return nil, err
	}

	// Store fetched employees in cache
	employeesJSON, err := json.Marshal(employee)
	if err != nil {
		return nil, fmt.Errorf("error marshalling employees data: %v", err)
	}
	err = redisClient.Client.Set(ctx, request.Email, employeesJSON, 5*time.Minute).Err()
	if err != nil {
		return nil, fmt.Errorf("error storing employees data in cache: %v", err)
	}

	return employeeDetails, nil
}
