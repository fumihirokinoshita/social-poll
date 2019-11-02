package main

import "net/http"

func main() {}

func withAPIKey(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isValidAPIKey(r, URL.Query().Get("key")) {
			respondErr(w, r, http.StatusUnauthorized, "不正なAPIです")
			return
		}
		fn(w, r)
	}
}

func isValidAPIKey(key string) bool {
	return key == "abc123"
}

func withVars(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		OpenVars(r)
		defer CloseVars(r)
		fn(w, r)
	}
}
