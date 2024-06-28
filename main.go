package main

import (
	"context"
	"data-manage/utils/excel"
	"data-manage/utils/mysql"
	"data-manage/utils/redis"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	// initialize postgres client
	err := mysql.InitMySQL(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer mysql.CloseMySQL()

	err = redis.InitRedisClusterClient(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redis.CloseRedis()

	err = excel.ImportDataFromExcel("data\\Sample_Employee_data_xlsx (1).xlsx", "uk-500")
	if err != nil {
		log.Fatal("Error importing data from Excel:", err)
	}

	// Start the router
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	startRouter(router)
}
func startRouter(router *gin.Engine) {
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start the router: %v", err)
	}
}
