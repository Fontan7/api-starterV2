package storage

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DB holds the connections to various data stores.
type DB struct {
	PgStore *PgStore
    //Mongo
}

type Storer interface {
	GetSomethingByID(*gin.Context, int) string
}


var dbInstance *DB

// OpenAll initializes all necessary database connections.
func OpenAll() (*DB, error) {
    ctx := context.Background()
	pgConnString := os.Getenv("PG_CONN_STRING")
    if pgConnString == "" {
        return nil, fmt.Errorf("PostgreSQL connection string is empty")
    }

    var wg sync.WaitGroup
    var pgPool *pgxpool.Pool
    errCh := make(chan error, 1)

    wg.Add(1)
    go func() {
        defer wg.Done()
        var err error
        pgPool, err = openPgStore(ctx, pgConnString)
        if err != nil {
            errCh <- fmt.Errorf("failed to open PostgreSQL: %v", err)
            return
        }
    }()

    wg.Wait()
    close(errCh)
    if err := <-errCh; err != nil {
        return nil, err
    }

    db := &DB{
        PgStore: newPgStore(pgPool),
    }
    return db, nil
}

func CloseAll() {
	if dbInstance != nil {
		if dbInstance.PgStore != nil {
			dbInstance.PgStore.pool.Close()
		}
		// Close other storages here
		dbInstance = nil // Reset the instance after closing
	}
}

func openPgStore(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	fmt.Println("Opening pg database pool...")

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	poolConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeExec

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	fmt.Println("Health check...")
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to verify connection to database: %v", err)
	}

	fmt.Println("Successfully connected to pg database")
	return pool, nil
}
