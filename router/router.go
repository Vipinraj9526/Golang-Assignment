package router

import (
	"data-manage/business"
	"data-manage/commons/constants"
	"data-manage/handlers"
	"data-manage/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// GetRouter is used to get the router configured with the middlewares and the routes
func GetRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middlewares...)
	router.Use(gin.Recovery())

	employeeRepository := repositories.NewEmployeesRepository()
	getEmployeesService := business.NewGetEmployeesService(employeeRepository)
	getEmployeesController := handlers.NewGetEmployeesController(getEmployeesService)

	updateEmployeeService := business.NewUpdateEmployeeDetailsService(employeeRepository)
	updateEmployeeController := handlers.NewUpdateEmployeeDetailsController(updateEmployeeService)

	deleteEmployeeService := business.NewDeleteEmployeeDetailsService(employeeRepository)
	deleteEmployeeController := handlers.NewDeleteEmployeeDetailsController(deleteEmployeeService)

	v1Routes := router.Group("v1")
	{
		// Health Check
		v1Routes.GET(constants.HealthCheck, func(c *gin.Context) {
			response := map[string]string{
				"message": "API is up and running",
			}
			c.JSON(http.StatusOK, response)
		})
		v1Routes.POST(constants.GetEmployees, getEmployeesController.GetEmployeesHandler)
		v1Routes.DELETE(constants.DeleteEmployee, deleteEmployeeController.DeleteEmployeeDetailsHandler)
		v1Routes.PATCH(constants.UpdateEmployee, updateEmployeeController.UpdateEmployeeDetailsHandler)
	}

	return router
}
