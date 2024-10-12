package main

import (
	controller "carhub/internal/controllers"
	"carhub/internal/model"
	"carhub/internal/repository"
	"carhub/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Configuración de conexión a PostgreSQL
	dsn := "host=localhost user=mi_usuario password=mi_contraseña dbname=mi_db port=5432 sslmode=disable TimeZone=America/Los_Angeles"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error al conectar a la base de datos: ", err)
	}

	// Migrar el modelo
	// Migrar los modelos
	db.AutoMigrate(
		&model.Brand{},
		&model.CalendarSettings{},
		&model.FuelType{},
		&model.GeneralSettings{},
		&model.LoanSettings{},
		&model.PaymentPlan{},
		&model.Price{},
		&model.Reservation{},
		&model.User{},
		&model.Vehicle{},
	)

	// Configurar los layers
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	vehicleRepo := repository.NewVehicleRepository(db)
	vehicleService := service.NewVehicleService(vehicleRepo)
	vehicleController := controller.NewVehicleController(vehicleService)

	// Migraciones de la base de datos
	db.AutoMigrate(&model.Brand{}, &model.Model{})

	// Inicialización de repositorios
	brandRepo := repository.NewBrandRepository(db)
	modelRepo := repository.NewModelRepository(db)

	// Inicialización de servicios
	brandService := service.NewBrandService(brandRepo)
	modelService := service.NewModelService(modelRepo)

	// Inicialización de controladores
	brandController := controller.NewBrandController(*brandService)
	modelController := controller.NewModelController(*modelService)

	// Configurar el servidor
	r := gin.Default()
	r.POST("/users", userController.CreateUser)
	r.GET("/users", userController.GetAllUsers)
	r.POST("/vehicles", vehicleController.CreateVehicle)
	r.GET("/vehicles/:id", vehicleController.GetVehicleByID)
	r.PUT("/vehicles/:id", vehicleController.UpdateVehicle)
	r.DELETE("/vehicles/:id", vehicleController.DeleteVehicle)
	r.POST("/brands", brandController.Create)
	r.GET("/brands", brandController.GetAll)
	r.POST("/models", modelController.Create)
	r.GET("/models", modelController.GetAll)

	// Iniciar el servidor
	r.Run(":8080")
}
