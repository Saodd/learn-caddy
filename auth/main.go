package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/Saodd/giary/giary"
	"github.com/gin-gonic/gin"
	"learn-caddy/common"
	"time"
)

var cc = giary.NewClient([]byte(common.Secret))

func AuthMiddleware(ctx *gin.Context) {
	// 这里是之前的签发Token的代码
	token, _ := json.Marshal(&common.UserToken{Name: "Lewin Lan", Expired: time.Now().Unix() + 60})
	tokenCipher := cc.Seal(token) // 加密时可以前后加盐，这里暂时不折腾了
	tokenCipherB64 := base64.StdEncoding.EncodeToString(tokenCipher)
	ctx.SetCookie(common.CookieKey, tokenCipherB64, 3600, "/", "", false, true)
	// 执行下一个
	ctx.Next()
}

func main() {
	app := gin.Default()
	app.Use(AuthMiddleware) // 这里把上面自己写的中间件设置到引擎中
	app.GET("/auth", func(context *gin.Context) {
		context.String(200, "Auth Passed.")
	})
	app.Run("0.0.0.0:30000")
}
