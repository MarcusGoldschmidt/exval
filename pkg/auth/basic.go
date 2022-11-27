package auth

import (
	"exval/pkg"
	"net/http"
)

func BasicAuthentication(opt *pkg.ExvalOptions) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			username, password, ok := r.BasicAuth()

			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if username == opt.AuthUser && password == opt.AuthPassword {
				return
			}

			w.WriteHeader(http.StatusUnauthorized)
		})
	}
}
