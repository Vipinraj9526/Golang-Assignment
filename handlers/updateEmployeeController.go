package handlers

import (
	"data-manage/business"
	"data-manage/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateEmployeeDetailsController struct {
	service *business.UpdateEmployeeDetailsService
}

func NewUpdateEmployeeDetailsController(service *business.UpdateEmployeeDetailsService) *UpdateEmployeeDetailsController {
	return &UpdateEmployeeDetailsController{
		service: service,
	}
}
func (controller *UpdateEmployeeDetailsController) UpdateEmployeeDetailsHandler(ctx *gin.Context) {
	var request models.UpdateEmployeeDetails
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	err = controller.service.UpdateEmployeeDetails(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, "Updated Successfully")
}
