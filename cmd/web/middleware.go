package main

import "net/http"

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request)){
		fmt.Println("Hit the page")
		next.ServeHTTP(w,r)
	}
}

func NoSruve(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookies(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: false,
		SameSite: http.SameSiteLaxMode
	})
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next);
}