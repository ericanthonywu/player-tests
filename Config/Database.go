package Config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
)

var (
	DB *gorm.DB
)

func InitDB() {
	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"))

	logLevel, err := strconv.Atoi(os.Getenv("DATABASE_LOG_LEVEL"))

	if err != nil {
		panic("env DATABASE_LOG_LEVEL is not an integer")
	}

	con, err := gorm.Open(postgres.Open(DSN), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.LogLevel(logLevel)),
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	DB = con

	g := gen.NewGenerator(gen.Config{
		OutPath:      "Model/Database",
		OutFile:      "dto",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		WithUnitTest: true,
		ModelPkgPath: "Database",
	})

	g.UseDB(DB)

	// generate Database if the Database is used only
	// don't generate unused Database

	// Generates All Table in Database
	//g.GenerateAllTable()

	// Generate Specify Table in Database
	//g.GenerateModel("master_type_schedules")

	//g.Execute()
}
