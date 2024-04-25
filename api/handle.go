package api

import (
	"api-starterV2/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
//	@Summary		Health check
//	@Description	always returns OK
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	string
//	@Success		200
//	@Failure		500
//	@Router			/health [get]
func handleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"code":   http.StatusOK,
	})
}

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
