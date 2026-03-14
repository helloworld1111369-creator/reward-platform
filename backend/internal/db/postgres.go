package db

import (
    "database/sql"
    "log"
    _ "github.com/lib/pq"
)

func Connect() *sql.DB {
    connStr := "user=postgres dbname=rewarddb sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal("Database ping failed:", err)
    }

    log.Println("Database connected successfully")
    return db
}
