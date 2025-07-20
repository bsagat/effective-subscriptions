package httpserver

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"submanager/internal/adapters/http/routers"
	"submanager/internal/core/domain"
	"submanager/internal/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type API struct {
	server *http.Server
}

func New(host, port string, subsService domain.SubsService, log logger.Logger) *API {
	r := gin.New()
	SetSwagger(r)

	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Info("HTTP Request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency", time.Since(start),
		)
	})

	subsHandler := routers.NewSubsHandler(subsService, log)
	subsHandler.RegisterSubsRoutes(r.Group("/subs"))

	return &API{
		server: &http.Server{
			Handler: r,
			Addr:    fmt.Sprintf("%s:%s", host, port),
		},
	}
}

func SetSwagger(r *gin.Engine) {
	// swagger json path
	url := ginSwagger.URL("/swagger-docs/swagger.json")

	r.GET("/swagger-docs/swagger.json", func(ctx *gin.Context) {
		ctx.File("./docs/swagger.json")
	})

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func (a *API) StartServer() {
	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}

func (a *API) Close() error {
	return a.server.Close()
}
