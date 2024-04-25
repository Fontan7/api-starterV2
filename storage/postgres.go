package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgStore struct {
	pool *pgxpool.Pool
}

func newPgStore(pool *pgxpool.Pool) *PgStore{
	return &PgStore{pool: pool}
}
 
func (pg *PgStore) GetSomethingByID(c *gin.Context, id int) string {return "hello bitch"}