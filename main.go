package main

import (
	controllers "architecture/ws/services/IO"
	"architecture/ws/services/alu"
	"architecture/ws/services/bus"
	"architecture/ws/services/controlunit"
	"architecture/ws/services/memory"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize registers, memory, ALU, and bus
	registers := memory.NewRegister()
	memory := memory.NewMemory()
	alu := alu.NewALU()
	dataBus := bus.NewDataBus()

	controlUnit := controlunit.NewControlUnit(dataBus, registers, memory, alu)

	// Initialize Gin router
	r := gin.Default()
	controller := controllers.NewController(controlUnit)

	// Define routes
	r.GET("/fetch", controller.Fetch)
	r.GET("/decode", controller.Decode)
	r.GET("/execute", controller.Execute)
	r.GET("/run_cycle", controller.RunCycle)
	r.POST("/load_instruction", controller.LoadInstruction)
	r.POST("/load_instructions", controller.LoadInstructions)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start the server: ", err)
	}
}

// package main

// import (
// 	"architecture/initializers"
// 	controllers "architecture/ws/services/IO"
// 	"architecture/ws/services/alu"
// 	"architecture/ws/services/bus"
// 	"architecture/ws/services/controlunit"
// 	"architecture/ws/services/memory"
// 	"log"
// 	"os"

// 	"github.com/gin-gonic/gin"
// )

// func init() {
// 	initializers.LoadEnvVariables()
// }

// func main() {

// 	registers := memory.NewRegister()
// 	memory := memory.NewMemory()
// 	alu := alu.NewALU()
// 	dataBus := bus.NewDataBus()

// 	controlUnit := controlunit.NewControlUnit(dataBus, registers, memory, alu)
// 	controller := controllers.NewController(controlUnit)

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "3001"
// 	}

// 	r := gin.Default()

// 	r.GET("/fetch", controller.Fetch)
// 	r.GET("/decode", controller.Decode)
// 	r.GET("/execute", controller.Execute)
// 	r.GET("/run_cycle", controller.RunCycle)
// 	r.POST("/load_instruction", controller.LoadInstruction)
// 	r.POST("/load_instructions", controller.LoadInstructions)
// 	// r := gin.Default()
// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	// io.BulkWriteInMemoryRoutes(router)
// 	if err := router.Run(":" + port); err != nil {
// 		log.Fatalf("Error starting server: %v", err)
// 	}

// }
