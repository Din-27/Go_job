package main

import (
	"github.com/Din-27/Go_job/internal/api/routes"
	"github.com/Din-27/Go_job/internal/middleware"
	"github.com/Din-27/Go_job/internal/utils"
	"github.com/gin-gonic/gin"
	// "time"
)

func main() {
	utils.LoadEnv()
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	routes.Services(router)
}
