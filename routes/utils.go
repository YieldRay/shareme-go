package routes

import (
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

func isNamespaceValid(namespace string) bool {
	ok, _ := regexp.Match("^[a-zA-Z0-9]{1,16}$", []byte(namespace))
	return ok
}

// check the UA
func notBrowser(ua string) bool {
	if strings.HasPrefix(ua, "curl") {
		return true
	}
	if strings.HasPrefix(ua, "Mozilla") {
		return false
	}
	if len(ua) < 10 {
		return true
	}

	return false
}

func contains(slice []string, s string) bool {
	for _, value := range slice {
		if value == s {
			return true
		}
	}
	return false
}

func getMimeType(name string) string {
	if strings.HasSuffix(name, ".js") {
		return "application/javascript"
	}
	if strings.HasSuffix(name, ".css") {
		return "text/css"
	}
	return ""
}

func getOrigin(c *gin.Context) string {
	if host := c.Request.Host; len(host) > 0 {
		scheme := c.Request.URL.Scheme
		if len(scheme) == 0 {
			scheme = "http"
		}
		return scheme + "://" + host
	}
	if origin := c.Request.Header.Get("Origin"); len(origin) > 0 {
		return origin
	}
	return "${domain}"
}
