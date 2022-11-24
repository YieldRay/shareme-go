package routes

import (
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"shareme/db"
)

func StaticMiddleware(g *gin.Engine, db db.IDB, efs fs.FS) {
	readFile := func(name string) []byte {
		f, _ := fs.ReadFile(efs, "public/"+name)
		return f
	}
	sendFile := func(ctx *gin.Context, name string) {
		if buf := readFile(name); len(buf) > 0 {
			ct := getContentType(name)
			if len(ct) == 0 {
				ct = http.DetectContentType(buf)
			}
			ctx.Data(http.StatusOK, ct, buf)
		} else {
			ctx.String(http.StatusOK, "404")
		}
	}

	g.GET("/", func(ctx *gin.Context) {
		sendFile(ctx, "index.html")
	})

	g.GET("/:namespace", func(ctx *gin.Context) {
		ua := ctx.Request.Header.Get("user-agent")
		if notBrowser(ua) {
			namespace, isNamespaceValid := getNamespace(ctx)
			if !isNamespaceValid {
				return
			}
			respondGetNamespace(ctx, namespace, db)
		} else {
			namespace := ctx.Param("namespace")
			if isNamespaceValid(namespace) {
				sendFile(ctx, "index.html")
			} else {
				sendFile(ctx, namespace)
			}
		}
	})
}
