package routes

import (
	"regexp"
	"strings"
)

func isNamespaceValid(namespace string) bool {
	ok, _ := regexp.Match("^[a-zA-Z0-9]{1,16}$", []byte(namespace))
	return ok
}

// check the UA
func notBrowser(ua string) bool {
	if strings.Index(ua, "curl") == 0 {
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


func getContentType(name string) string {
	if strings.HasSuffix(name, ".js") {
		return "application/javascript"
	}
	if strings.HasSuffix(name, ".css") {
		return "text/css"
	}
	return ""
}