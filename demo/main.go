package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/guileen/httpcache"
)

type Handler struct {
	upstream http.Handler
}

func (h Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// req.Headers()['Cache-Control'] = "....max-age..."
	req.Header.Set("X-Cache-Key", "testkey")
	h.upstream.ServeHTTP(rw, req)
	log.Println("rw.header", rw.Header())
	hit := rw.Header().Get("X-Cache") == "HIT"
	// miss := rw.Header().Get("X-Cache") == "MISS"
	if hit {
		log.Println("hit")
	} else {
		log.Println("skip")
	}
}

func main() {
	// proxy := &httputil.ReverseProxy{
	// 	Director: func(r *http.Request) {
	// 	},
	// }
	url, err := url.Parse("http://127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(url)

	handler := httpcache.NewHandler(httpcache.NewMemoryCache(), proxy)
	handler.Shared = true

	listen := ":5000"

	log.Printf("proxy listening on http://%s", listen)
	log.Fatal(http.ListenAndServe(listen, &Handler{
		upstream: handler,
	}))
}
