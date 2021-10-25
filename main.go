package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {

			remote, err := url.Parse("http://google.com")
			if err != nil {
				panic(err)
			}
			log.Println(r.URL)
			r.Host = remote.Host
			p.ServeHTTP(w, r)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
