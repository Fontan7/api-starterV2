package api

import (
	"api-starterV2/types"
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
func handleGetPublicSomething(c *gin.Context, app types.App) {
	pg := app.DB().PgStore

	result := pg.GetSomethingByID(c, 1)
	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"code":     http.StatusOK,
		"response": result,
	})
}

func handleGetPrivateSomething(c *gin.Context, app types.App) {}
