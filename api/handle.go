package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary		Health check
// @Description	always returns OK
// @Tags			health
// @Produce		json
// @Success		200	{object}	string
// @Failure		500
// @Router			/health [get]
func handleHealthCheck(c *gin.Context) {
	if true {
		c.Error(errors.New("non fatal error"))
		c.AbortWithError(http.StatusInternalServerError, errors.New("a fatal error"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"code":   http.StatusOK,
	})
}

// HealthCheck godoc
// @Summary		Gets something public from the database
// @Description	Using postgres Gets a public entity from the db and returns it plain
// @Tags			get
// @Produce		json
// @Success		200	{object}	string
// @Failure		500
// @Router			/v1/public/get-something/:id [get]
func handleGetPublicSomething(c *gin.Context, app App) {
	pg := app.DB().PgStore

	result := pg.GetSomethingByID(c, 1)
	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"code":     http.StatusOK,
		"response": result,
	})
}

func handleGetPrivateSomething(c *gin.Context, app App) {}
