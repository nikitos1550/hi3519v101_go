// +build openapi

package openapi

import (
    "log"
    "net/http"
)


// Middleware function, which will be called for each request
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("X-Session-Token")
        log.Println("X-Session-Token:", token)
        next.ServeHTTP(w, r)

        /*
        if user, found := amw.tokenUsers[token]; found {
            // We found the token in our map
            log.Printf("Authenticated user %s\n", user)
            // Pass down the request to the next middleware (or final handler)
            next.ServeHTTP(w, r)
        } else {
            // Write an error and stop the handler chain
            http.Error(w, "Forbidden", http.StatusForbidden)
        }
        */
    })
}
