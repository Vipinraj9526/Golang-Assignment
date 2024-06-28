package business

import (
	"data-manage/models"
	"data-manage/repositories"
	"data-manage/utils/mysql"
	"data-manage/utils/redis"

	"github.com/gin-gonic/gin"
)

type DeleteEmployeeDetailsService struct {
	deleteEmployeeDetailsServiceRegistry repositories.EmployeeRepository
}

func NewDeleteEmployeeDetailsService(deleteEmployeeDetailsServiceRegistry repositories.EmployeeRepository) *DeleteEmployeeDetailsService {
	return &DeleteEmployeeDetailsService{
		deleteEmployeeDetailsServiceRegistry: deleteEmployeeDetailsServiceRegistry,
	}
}

func (service *DeleteEmployeeDetailsService) DeleteEmployeeDetails(ctx *gin.Context, request models.DeleteEmployeeDetails) error {
	// Fetch employee details from the database if not cached
	db := mysql.GetMySQLClient().GormDB
	conditions := map[string]interface{}{
		"email": request.Email,
	}
	err := service.deleteEmployeeDetailsServiceRegistry.DeleteEmployee(ctx, db, &models.Employee{}, conditions)
	if err != nil {
		return err
	}

	// Delete record from redis
	redis.GetRedisClient().Client.Del(ctx, request.Email)
	return nil
}
