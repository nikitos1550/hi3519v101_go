// +build !openapi

package openapi

import (
	"net/http"
)

func AddApiRoute(name, pattern, method string, handlerfunc http.HandlerFunc) {}
func AddRoute(name, pattern, method string, handlerfunc http.HandlerFunc) {}
func Init() {}
func Start() {}