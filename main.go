package main

import (
	"context"
	"data-manage/commons/constants"
	"data-manage/router"
	"data-manage/utils/excel"
	"data-manage/utils/mysql"
	"data-manage/utils/redis"
	"fmt"
	"log"
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

	startRouter(ctx)
}

func startRouter(ctx context.Context) {
	router := router.GetRouter()
	err := router.Run(fmt.Sprintf(":%d", constants.PortDefaultValue))
	if err != nil {
	}
	fmt.Println(fmt.Sprintf("Server is running on port %d", constants.PortDefaultValue))

}
