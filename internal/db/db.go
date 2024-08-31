package db

import (
	"database/sql"
	"fmt"
	"github.com/gabrielvieira/go-api/internal/config"
	"github.com/gabrielvieira/go-api/internal/db/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type DB struct {
	gorm.DB
	config *config.Config
}

func New(config *config.Config) *DB {
	return &DB{config: config}
}

func (d *DB) Open() error {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// sample dsn user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.config.DBUser, d.config.DBPassword, d.config.DBUrl, d.config.DBSchema,
	)

	mysqlConn, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	mysqlConn.SetConnMaxLifetime(time.Second * d.config.DBConnectionMaxLifetimeSeconds)
	mysqlConn.SetMaxOpenConns(d.config.DBMaxConnections)
	mysqlConn.SetMaxIdleConns(d.config.DBMaxIdleConnections)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mysqlConn,
	}), &gorm.Config{})

	err = d.RunMigrations()
	if err != nil {
		return err
	}

	d.DB = *gormDB
	return nil
}

func (d *DB) RunMigrations() error {
	err := d.DB.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	return nil
}
