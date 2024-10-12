package controller

import (
	"carhub/internal/service"
	"net/http"

	"carhub/internal/model"

	"github.com/gin-gonic/gin"
)

// BrandController maneja las solicitudes relacionadas con la entidad Brand
type BrandController struct {
	Service service.BrandService // Suponiendo que tienes un servicio para Brand
}

// NewBrandController crea un nuevo BrandController
func NewBrandController(s service.BrandService) *BrandController {
	return &BrandController{Service: s}
}

// Create crea una nueva marca
func (bc *BrandController) Create(c *gin.Context) {
	var brand model.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.Service.Create(&brand); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, brand)
}

// GetAll obtiene todas las marcas
func (bc *BrandController) GetAll(c *gin.Context) {
	brands, err := bc.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, brands)
}
