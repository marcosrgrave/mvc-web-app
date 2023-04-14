package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConn() (*sql.DB, error) {
	// TODO: .env
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/tests?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Creates application tables and adds data to them, if they don't exist.
func Init() {
	db, err := DBConn()
	if err != nil {
		panic("Database connection error.\n" + err.Error())
	}
	defer db.Close()

	createTableProducts(db)
	createTableCategories(db)
	createTableUsers(db)
}

func createTableProducts(db *sql.DB) {
	sql_command := `
		CREATE TABLE IF NOT EXISTS tb_products (
			id VARCHAR(255) NOT NULL,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			price DECIMAL(10, 2),
			quantity INT,
			created_at DATETIME,
			creator VARCHAR(255),
			PRIMARY KEY (id)
		);
	`
	_, err := db.Exec(sql_command)
	if err != nil {
		panic("Products table creation error.\n" + err.Error())
	}
}

func createTableCategories(db *sql.DB) {}

func createTableUsers(db *sql.DB) {}
