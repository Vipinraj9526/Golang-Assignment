package handlers

import (
	"data-manage/business"
	"data-manage/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteEmployeeDetailsController struct {
	service *business.DeleteEmployeeDetailsService
}

func NewDeleteEmployeeDetailsController(service *business.DeleteEmployeeDetailsService) *DeleteEmployeeDetailsController {
	return &DeleteEmployeeDetailsController{
		service: service,
	}
}
func (controller *DeleteEmployeeDetailsController) DeleteEmployeeDetailsHandler(ctx *gin.Context) {
	var request models.DeleteEmployeeDetails
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	err = controller.service.DeleteEmployeeDetails(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, "Deleted Successfully")
}
