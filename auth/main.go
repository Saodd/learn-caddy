package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/Saodd/giary/giary"
	"github.com/gin-gonic/gin"
	"learn-caddy/common"
	"time"
)

func main() {
	// 实例化一个用于加密解密的封装对象
	var cc = giary.NewClient([]byte(common.Secret))
	app := gin.Default()
	app.GET("/auth", func(context *gin.Context) {
		// 生成Token
		token, _ := json.Marshal(&common.UserToken{Name: "Lewin Lan", Expired: time.Now().Unix() + 60})
		tokenCipher := cc.Seal(token) // 加密时可以前后加盐，这里暂时不折腾了
		tokenCipherB64 := base64.StdEncoding.EncodeToString(tokenCipher)
		// 设置Cookie
		context.SetCookie(common.CookieKey, tokenCipherB64, 3600, "/", "localhost", false, true)
		context.String(200, "Auth Passed.")
	})
	app.Run("0.0.0.0:30000")
}
