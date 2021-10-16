package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
	//	"github.com/k0kubun/pp"
)

// DumpRequest
func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>")
}

// set cookie
func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookieValue := r.FormValue("cp")

	// TODO: Compare same result method
	// cookieValue := r.URL.Query().Get("cp")

	if cookieValue != "" {
		setCookie(w, cookieValue, time.Now().Add(24*time.Hour))
		fmt.Fprintf(w, "Cookie may be added. Please check result.")
	}

	w.WriteHeader(http.StatusOK)
}

// set Cookie func
func setCookie(w http.ResponseWriter, cookieValue string, expireTime time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     "sample_name",
		Value:    cookieValue,
		Domain:   "localhost",
		Path:     "/",
		Expires:  expireTime,
		HttpOnly: true,
	})

}

// dump cookie info
func dumpCookieHandler(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	w.WriteHeader(http.StatusOK)
}

// func handlerDigest(w http.ResponseWriter, r *http.Request) {
// 	pp.Printf("URL: %s\n", r.URL.String())
// }

// This is simple web server.
func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	http.HandleFunc("/set_cookie", setCookieHandler)
	http.HandleFunc("/dump_cookie", dumpCookieHandler)
	//	http.HandleFunc("/", handlerDigest)
	log.Println("start http listening :9090")
	httpServer.Addr = ":9090"
	log.Println(httpServer.ListenAndServe())
}
