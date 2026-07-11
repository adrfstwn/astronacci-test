package main

import (
    "log"
    "os"

    "backend/internal/db"
    "backend/internal/handler"
    "backend/internal/repository"
    "backend/internal/router"
    "backend/internal/service"
)

func main() {
    dbPath := os.Getenv("DB_PATH")
    if dbPath == "" {
        dbPath = "vouchers.db"
    }

    conn, err := db.Connect(dbPath)
    if err != nil {
        log.Fatalf("failed to connect db: %v", err)
    }
    defer conn.Close()

    repo := repository.NewVoucherRepository(conn)
    svc := service.NewVoucherService(repo)
    h := handler.NewVoucherHandler(svc)

    r := router.Setup(h)
    r.Run(":8080")
}