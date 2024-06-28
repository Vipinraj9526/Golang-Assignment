package business

import (
	"data-manage/models"
	"data-manage/repositories"
	"data-manage/utils/mysql"
	"data-manage/utils/redis"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateEmployeeDetailsService struct {
	updateEmployeeDetailsServiceRegistry repositories.EmployeeRepository
}

func NewUpdateEmployeeDetailsService(updateEmployeeDetailsServiceRegistry repositories.EmployeeRepository) *UpdateEmployeeDetailsService {
	return &UpdateEmployeeDetailsService{
		updateEmployeeDetailsServiceRegistry: updateEmployeeDetailsServiceRegistry,
	}
}

func (service *UpdateEmployeeDetailsService) UpdateEmployeeDetails(ctx *gin.Context, request models.UpdateEmployeeDetails) error {
	// Fetch employee details from the database if not cached
	db := mysql.GetMySQLClient().GormDB
	conditions := map[string]interface{}{
		"email": request.Email,
	}
	var employee models.Employee
	err := service.updateEmployeeDetailsServiceRegistry.UpdateEmployee(ctx, db, &employee, conditions, request)
	if err != nil {
		return err
	}

	// Store fetched employees in cache
	redis.GetRedisClient().Client.Del(ctx, request.Email)
	redis.GetRedisClient().Client.Set(ctx, request.Email, employee, 5*time.Minute)
	return nil
}
