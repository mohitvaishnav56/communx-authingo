package router

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"
)

// Define basic static routes mapping for the Gateway
var RouteMap = map[string]string{
	"/api/posts":    "http://localhost:8081",
	"/api/comments": "http://localhost:8082",
}

func SetupGatewayRoutes(r chi.Router) {
	r.Route("/api", func(r chi.Router) {
		// Apply AuthMiddleware to all /api routes
		r.Use(AuthMiddleware)

		r.HandleFunc("/*", func(w http.ResponseWriter, req *http.Request) {
			path := req.URL.Path
			var targetURL *url.URL
			var err error

			for prefix, target := range RouteMap {
				if strings.HasPrefix(path, prefix) {
					targetURL, err = url.Parse(target)
					break
				}
			}

			if targetURL == nil || err != nil {
				http.Error(w, "Service not found for route", http.StatusNotFound)
				return
			}

			proxy := httputil.NewSingleHostReverseProxy(targetURL)

			// Optional: modify request before forwarding
			req.URL.Host = targetURL.Host
			req.URL.Scheme = targetURL.Scheme
			req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
			req.Host = targetURL.Host

			fmt.Println("Proxying request to:", targetURL.String()+req.URL.Path)
			proxy.ServeHTTP(w, req)
		})
	})
}
