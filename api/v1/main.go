package v1

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Database *sql.DB

// Ping godoc
// @Summary ping
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /v1/ping [get]
func Ping(g *gin.Context) {
	g.JSON(http.StatusOK, "pong")
}

func Handler(r *gin.RouterGroup) {
	r.GET("/ping", Ping)

	UserHandler(r.Group("/user"))
	GroupHandler(r.Group("/group"))
}
