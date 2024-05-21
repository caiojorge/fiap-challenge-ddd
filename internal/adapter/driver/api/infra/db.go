package infra

import (
	"fmt"
	"log"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driven/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	db       *gorm.DB
}

func NewDB(host, port, user, password, dbName string) *DB {
	return &DB{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
	}
}

func (d *DB) GetConnection(dbType string) *gorm.DB {

	if dbType == "sqlite" {
		d.db = d.setupSQLite()
	}

	if dbType == "mysql" {
		d.db = d.setupMysql()
	}

	if d.db == nil {
		log.Fatalf("Database type %s not supported", dbType)
	}

	return d.db
}

func (d *DB) setupSQLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Customer{})
	return db
}

func (d *DB) setupMysql() *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.User, d.Password, d.Host, d.Port, d.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Database schema migrated")

	return db
}
