package handlers

import (
	"data-manage/business"
	"data-manage/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetEmployeesController struct {
	service *business.GetEmployeesService
}

func NewGetEmployeesController(service *business.GetEmployeesService) *GetEmployeesController {
	return &GetEmployeesController{
		service: service,
	}
}
func (controller *GetEmployeesController) GetEmployeesHandler(ctx *gin.Context) {
	var request models.GetEmployeeDetails
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	employee, err := controller.service.GetEmployees(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, employee)
}
