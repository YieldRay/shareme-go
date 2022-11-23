package route

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"shareme/db"
	"strings"
)

func getNamespace(ctx *gin.Context) (namespace string, ok bool) {
	namespace = ctx.Param("namespace")
	if len(namespace) == 0 {
		ctx.String(http.StatusBadRequest, `[ShareMe]: namespace cannot be empty`)
		return namespace, false
	}
	if !isNamespaceValid(namespace) {
		ctx.String(http.StatusBadRequest, `[ShareMe]: Namespace is invalid, is can only contains letters and numbers`)
		return "", false
	}
	return namespace, true
}

// After calling this method, should not access ctx anymore
func respondGetNamespace(ctx *gin.Context, namespace string, db db.IDB) {
	content, err := db.Get(namespace)
	if err != nil {
		ctx.String(http.StatusInternalServerError, `[ShareMe]: Server failed to load data`)
	} else {
		ctx.String(http.StatusOK, content)
	}
}

func getPostValue(c *gin.Context) (content string, err error) {
	ct := strings.ToLower(c.ContentType())
	if strings.Contains(ct, "application/json") { // for json
		var data []byte
		data, err = c.GetRawData()
		var body map[string]string
		err = json.Unmarshal(data, &body)

		keys := make([]string, 0, len(body))
		for k := range body {
			keys = append(keys, k)
		}

		if contains(keys, "t") {
			content = body["t"]
		} else {
			err = errors.New("json body has no `t`, just send content from db")
		}
		return
	}
	if strings.Contains(ct, "application/x-www-form-urlencoded") { // for x-www-form-urlencoded
		var has bool
		content, has = c.GetPostForm("t")
		if !has {
			err = errors.New("form body has no `t`, just send content from db")
		}
		return
	}

	return "", errors.New("content-type is not set or can be recognized")
}

func APIMiddleware(g *gin.Engine, db db.IDB) {
	g.POST("/:namespace", func(ctx *gin.Context) {
		namespace, ok := getNamespace(ctx)
		if !ok {
			return // response has been sent, do nothing
		}
		content, err := getPostValue(ctx)

		if err != nil {
			// get from db
			respondGetNamespace(ctx, namespace, db)
		} else {
			// set to db
			if db.Set(namespace, content) {
				ctx.String(http.StatusOK, `[ShareMe]: Ok`)
			} else {
				ctx.String(http.StatusInternalServerError, `[ShareMe]: Server failed to save data`)
			}
		}
	})

}
