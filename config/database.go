package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	MySql  *sqlx.DB
	GormDB *gorm.DB
}

var (
	DbConn = &DB{}
)

func Init() *DB {

	err_host := godotenv.Load(".env")
	if err_host != nil {
		fmt.Println(err_host)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PSWD")
	dbname := os.Getenv("DB_NAME")

	psql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sqlx.Open("pgx", psql)
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxOpenConns(7)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Duration(300 * time.Second))

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	DbConn.MySql = db

	config, err := pgx.ParseConfig(psql)
	if err != nil {
		log.Fatalln(err)
	}

	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: stdlib.OpenDB(*config),
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatalln(err)
	}

	DbConn.GormDB = gormDb

	return DbConn
}
