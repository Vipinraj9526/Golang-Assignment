package mysql

import (
	"context"
	"data-manage/utils/configs"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mysqlClient *MySQLClient
)

// MySQLClient holds the MySQL client connections
type MySQLClient struct {
	GormDB *gorm.DB
	SqlDB  *sql.DB
}

// InitMySQL initializes MySQL client
func InitMySQL(ctx context.Context) error {
	// Load MySQL configuration
	mysqlConfig, err := configs.LoadConfig("configs/mysql.yml")
	if err != nil {
		return fmt.Errorf("failed to load MySQL configuration: %v", err)
	}

	// Establish connection to MySQL database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.MySQL.Username, mysqlConfig.MySQL.Password, mysqlConfig.MySQL.Host, mysqlConfig.MySQL.Port, mysqlConfig.MySQL.Database)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open GORM connection: %v", err)
	}

	// Initialize MySQL client
	mysqlClient = &MySQLClient{
		GormDB: gormDB,
	}

	// Ping database to verify connection
	sqlDB, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("failed to get SQL database: %v", err)
	}
	err = sqlDB.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Connected to MySQL database successfully")

	return nil
}

// GetMySQLClient returns the MySQL client instance
func GetMySQLClient() *MySQLClient {
	return mysqlClient
}

// CloseMySQL closes the MySQL client connection
func CloseMySQL() error {
	if mysqlClient == nil {
		return nil
	}

	sqlDB, err := mysqlClient.GormDB.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying SQL database: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		return fmt.Errorf("failed to close MySQL database: %v", err)
	}

	log.Println("Closed MySQL database connection")

	return nil
}
