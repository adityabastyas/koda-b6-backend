package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// connection pool. bisa banyak koneksi ke banyak DB
var DB *pgxpool.Pool

var Conn *pgx.Conn

func ConnectDB() {
	//dari .env

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname,
	)

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic("gagal konek ke database: " + err.Error())
	}

	if err := pool.Ping(context.Background()); err != nil {
		panic("data base tidak merespon: " + err.Error())
	}

	DB = pool
	fmt.Println("database berhasil terhubung")

}

func ConnectConn() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic("gagal konek ke database: " + err.Error())
	}

	Conn = conn
	fmt.Println("conn berhasil terhubung!")
}
