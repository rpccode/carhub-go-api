package controller

import (
	"carhub/internal/model"
	"carhub/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ModelController maneja las solicitudes relacionadas con la entidad Model
type ModelController struct {
	Service service.ModelService // Suponiendo que tienes un servicio para Model
}

// NewModelController crea un nuevo ModelController
func NewModelController(s service.ModelService) *ModelController {
	return &ModelController{Service: s}
}

// Create crea un nuevo modelo
func (mc *ModelController) Create(c *gin.Context) {
	var model model.Model
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := mc.Service.Create(&model); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, model)
}

// GetAll obtiene todos los modelos
func (mc *ModelController) GetAll(c *gin.Context) {
	models, err := mc.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models)
}
