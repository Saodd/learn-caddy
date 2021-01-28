package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/Saodd/giary/giary"
	"github.com/gin-gonic/gin"
	"learn-caddy/common"
	"net/http"
	"time"
)

func main() {
	// 实例化一个用于加密解密的封装对象
	var cc = giary.NewClient([]byte(common.Secret))
	app := gin.Default()
	app.GET("/business/1", func(ctx *gin.Context) {
		// 获取Cookie
		tokenCipherB64, err := ctx.Cookie(common.CookieKey)
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Token Not Found.") // 在正式产品中请不要给出这么详细的错误提示
			return
		}

		// 解密Token
		tokenCipher, err := base64.StdEncoding.DecodeString(tokenCipherB64)
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Base64 Decode Failed.") // 在正式产品中请不要给出这么详细的错误提示
			return
		}
		token, err := cc.Open(tokenCipher)
		if err != nil {
			ctx.String(http.StatusUnauthorized, "AES Decode Failed.") // 在正式产品中请不要给出这么详细的错误提示
			return
		}
		var user common.UserToken
		if err = json.Unmarshal(token, &user); err != nil {
			ctx.String(http.StatusUnauthorized, "JSON Decode Failed.") // 在正式产品中请不要给出这么详细的错误提示
			return
		}

		// 验证用户Token是否有效
		if user.Expired < time.Now().Unix() {
			ctx.String(http.StatusUnauthorized, "Token Expired.") // 在正式产品中请不要给出这么详细的错误提示
			return
		}

		// 执行业务
		ctx.String(200, "Hello, Im business code. 30001")
	})
	app.GET("/_/health", func(ctx *gin.Context) {
		ctx.Status(200)
	})
	app.Run("0.0.0.0:30001")
}
