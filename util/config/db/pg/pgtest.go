package pg

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	pgModel "github.com/phungvandat/life-cafe-backend/model/pg"
)

const (
	dbHost     = "localhost"
	dbPort     = 5439
	dbUser     = "postgres"
	dbPassword = "lifecafetest"
	dbName     = "life-cafe-test"
)

// CreateTestDatabase func create test database
func CreateTestDatabase(t *testing.T) (*gorm.DB, func()) {
	testingPort := fmt.Sprintf("%d", dbPort)
	testingHost := fmt.Sprintf("%s", dbHost)

	if os.Getenv("POSTGRES_TESTING_HOST") != "" {
		testingHost = os.Getenv("POSTGRES_TESTING_HOST")
	}

	if os.Getenv("POSTGRES_TESTING_PORT") != "" {
		testingPort = os.Getenv("POSTGRES_TESTING_PORT")
	}

	connectionString := fmt.Sprintf("host = %s port=%s user=%s password=%s dbname=%s sslmode=disable", testingHost, testingPort, dbUser, dbPassword, dbName)

	db, dbErr := gorm.Open("postgres", connectionString)

	if dbErr != nil {
		t.Fatalf("Fail to create database. %s", dbErr.Error())
	}

	rand.Seed(time.Now().UnixNano())

	schemaName := "test" + strconv.FormatInt(rand.Int63(), 10)

	err := db.Exec("CREATE SCHEMA " + schemaName).Error
	if err != nil {
		t.Fatalf("Fail to create schema. %s", err.Error())
	}

	// set schema for current db connection
	err = db.Exec("SET search_path TO " + schemaName).Error
	if err != nil {
		t.Fatalf("Fail to set search_path to created schema. %s", err.Error())
	}

	return db, func() {
		err := db.Exec("DROP SCHEMA " + schemaName + " CASCADE").Error
		if err != nil {
			t.Fatalf("Fail to drop database. %s", err.Error())
		}
	}
}

// MigrateTables create database table
func MigrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
		pgModel.User{},
		pgModel.Category{},
		pgModel.Product{},
		pgModel.ProductCategory{},
		pgModel.Order{},
		pgModel.ProductOrder{},
	).Error
}
