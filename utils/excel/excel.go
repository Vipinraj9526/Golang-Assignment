package excel

import (
	"context"
	"data-manage/models"
	"data-manage/utils/mysql"
	"data-manage/utils/redis"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/xuri/excelize/v2"
)

// ImportDataFromExcel imports data from an Excel file into MySQL and caches it in Redis using email as the primary key
func ImportDataFromExcel(filename string, sheetName string) error {
	xlFile, err := excelize.OpenFile(filename)
	if err != nil {
		return fmt.Errorf("failed to open Excel file: %v", err)
	}

	// Check if the sheet name is provided and exists
	if sheetName == "" {
		return fmt.Errorf("sheet name cannot be blank")
	}

	// Get all rows from the specified sheet
	rows, err := xlFile.GetRows(sheetName)
	if err != nil {
		return fmt.Errorf("failed to get rows from sheet '%s' in Excel file: %v", sheetName, err)
	}

	// Initialize MySQL and Redis clients
	mysqlClient := mysql.GetMySQLClient()
	redisClient := redis.GetRedisClient()

	// Skip the first row as it contains headings
	rows = rows[1:]

	var wg sync.WaitGroup
	for _, row := range rows {
		if len(row) < 10 {
			continue // Skip incomplete rows
		}

		wg.Add(1)
		go func(row []string) {
			defer wg.Done()

			// Parse data from Excel row
			firstName := row[0]
			lastName := row[1]
			companyName := row[2]
			address := row[3]
			city := row[4]
			county := row[5]
			postal := row[6]
			phone := row[7]
			email := row[8]
			web := row[9]

			// Create Data struct and append to slice
			newData := models.Employee{
				FirstName:   firstName,
				LastName:    lastName,
				CompanyName: companyName,
				Address:     address,
				City:        city,
				County:      county,
				Postal:      postal,
				Phone:       phone,
				Email:       email,
				Web:         web,
			}

			// Store in MySQL using email as primary key
			err := mysqlClient.GormDB.Where(models.Employee{Email: email}).Assign(newData).FirstOrCreate(&models.Employee{}).Error
			if err != nil {
				log.Printf("Failed to store data '%+v' in MySQL: %v", newData, err)
				return
			}

			// Convert struct to JSON for Redis caching
			jsonData, _ := json.Marshal(newData)

			// Store in Redis
			err = redisClient.Client.Set(context.Background(), newData.Email, string(jsonData), 5*time.Minute).Err()
			if err != nil {
				log.Printf("Failed to cache data '%+v' in Redis: %v", newData, err)
			}
		}(row)
	}

	wg.Wait()

	log.Println("Imported data from Excel successfully")

	return nil
}
