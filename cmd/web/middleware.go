package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			page := r.URL.Path

			if page == "/" {
				page = "/home"
			}

			fmt.Printf("%s has hit the page:%s\n", ip, page)
			next.ServeHTTP(w, r)
		})
}

func NoSurf(next http.Handler) http.Handler{
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
