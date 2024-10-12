// controller/vehicle_controller.go
package controller

import (
	"carhub/internal/model"
	"carhub/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VehicleController struct {
	service service.VehicleService
}

func NewVehicleController(service service.VehicleService) *VehicleController {
	return &VehicleController{service: service}
}

func (c *VehicleController) CreateVehicle(ctx *gin.Context) {
	var vehicle model.Vehicle
	if err := ctx.ShouldBindJSON(&vehicle); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := c.service.CreateVehicle(&vehicle); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, vehicle)
}

func (c *VehicleController) GetVehicleByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	vehicle, err := c.service.GetVehicleByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	ctx.JSON(http.StatusOK, vehicle)
}

func (c *VehicleController) UpdateVehicle(ctx *gin.Context) {
	var vehicle model.Vehicle
	if err := ctx.ShouldBindJSON(&vehicle); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := c.service.UpdateVehicle(&vehicle); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, vehicle)
}

func (c *VehicleController) DeleteVehicle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.service.DeleteVehicle(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
