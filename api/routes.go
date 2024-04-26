package api

import (
	"api-starterV2/middleware"

	_ "api-starterV2/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initializes and returns a new Gin engine with configured routes.
func InitRouter(app App) *gin.Engine {
	gin.SetMode(app.GinMode())
	router := gin.Default()
	router.Use(middleware.ErrorManager())
	router.Use(middleware.CorsConfig())
	defineRoutes(router, app)
	return router
}

// defineRoutes sets up routes and their corresponding groups.
func defineRoutes(router *gin.Engine, app App) {
	registerHealthAndSwaggerRoutes(router)
	registerAPIRoutes(router, app)
}

// registerHealthAndSwaggerRoutes registers routes for health checks and Swagger UI.
func registerHealthAndSwaggerRoutes(router *gin.Engine) {
	router.GET("/health", handleHealthCheck)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// registerAPIRoutes sets up API versioning and their specific routes.
func registerAPIRoutes(router *gin.Engine, app App) {
	router.Use(middleware.CheckAPIKey(app.ClientKey())) // API key check middleware

	v1 := router.Group("/v1")
	defineV1Routes(v1, app)
}

// defineV1Routes registers routes for API version 1.
func defineV1Routes(rg *gin.RouterGroup, app App) {
	public := rg.Group("/public")
	definePublicRoutesV1(public, app)

	private := rg.Group("/private")
	private.Use(middleware.ValidateAndSetToken())
	definePrivateRoutesV1(private, app)
}

// definePublicRoutes registers public routes. /v1/public/...
func definePublicRoutesV1(rg *gin.RouterGroup, app App) {
	rg.GET("/get-something/:id", func(c *gin.Context) { handleGetPublicSomething(c, app) })
}

// definePrivateRoutes registers private routes requiring authentication.
func definePrivateRoutesV1(rg *gin.RouterGroup, app App) {
	rg.GET("/user/home", func(c *gin.Context) { handleGetPrivateSomething(c, app) })
}
