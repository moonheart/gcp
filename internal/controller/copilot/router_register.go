package copilot

import (
	"github.com/gin-gonic/gin"
	"ripper/internal/middleware"
)

func GinApi(g *gin.RouterGroup) {
	// Copilot 请求代理

	g.GET("/models", models)
	g.GET("/_ping", _ping)
	g.POST("/telemetry", postTelemetry)
	g.GET("/user", middleware.AccessTokenCheckAuth(), getLoginUser)
	g.GET("/user/orgs", middleware.AccessTokenCheckAuth(), getUserOrgs)
	g.GET("/api/v3/user", middleware.AccessTokenCheckAuth(), getLoginUser)
	g.GET("/api/v3/user/orgs", middleware.AccessTokenCheckAuth(), getUserOrgs)

	g.GET("/copilot_internal/v2/token", middleware.AccessTokenCheckAuth(), getCopilotInternalV2Token)
	g.POST("/v1/engines/copilot-codex/completions", codeCompletions)
	g.POST("/v1/engines/copilot-codex", codeCompletions)

	g.POST("/chat/completions", chatCompletions)

	g.GET("/api/v3/meta", v3meta)
	g.GET("/api/v3/", cliv3)
	g.GET("/agents", agents)
	g.GET("/teams/:teamID/memberships/:username", getMembership)
	g.GET("/", cliv3)

}
