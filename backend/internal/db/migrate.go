package db

import (
    "database/sql"
    "embed"
    "fmt"
    "sort"
    "strings"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func RunMigrations(conn *sql.DB) error {
    _, err := conn.Exec(`
        CREATE TABLE IF NOT EXISTS schema_migrations (
            version TEXT PRIMARY KEY,
            applied_at TEXT DEFAULT CURRENT_TIMESTAMP
        );
    `)
    if err != nil {
        return fmt.Errorf("failed to init schema_migrations: %w", err)
    }

    entries, err := migrationFiles.ReadDir("migrations")
    if err != nil {
        return fmt.Errorf("failed to read migrations dir: %w", err)
    }

    var names []string
    for _, e := range entries {
        if strings.HasSuffix(e.Name(), ".up.sql") {
            names = append(names, e.Name())
        }
    }
    sort.Strings(names)

    for _, name := range names {
        version := strings.TrimSuffix(name, ".up.sql")

        var exists string
        err := conn.QueryRow(`SELECT version FROM schema_migrations WHERE version = ?`, version).Scan(&exists)
        if err == nil {
            continue
        }

        content, err := migrationFiles.ReadFile("migrations/" + name)
        if err != nil {
            return fmt.Errorf("failed to read %s: %w", name, err)
        }

        if _, err := conn.Exec(string(content)); err != nil {
            return fmt.Errorf("failed to apply migration %s: %w", name, err)
        }

        if _, err := conn.Exec(`INSERT INTO schema_migrations (version) VALUES (?)`, version); err != nil {
            return fmt.Errorf("failed to record migration %s: %w", name, err)
        }

        fmt.Printf("applied migration: %s\n", version)
    }

    return nil
}
