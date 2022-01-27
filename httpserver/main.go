package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/golang/glog"
)

func main()  {

	// 日志级别
	flag.Set("v", "4")
	glog.V(1).Infoln("Starting http server...")

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/", index)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.ListenAndServe(":80", mux)
}


func index (w http.ResponseWriter, r *http.Request) {
	header := r.Header
	for k,v := range header {
		glog.V(2).Info("request header k = ", k, "v = ", v)
		io.WriteString(w, fmt.Sprintf("request header is k = %s, v = %s \n", k, v))
	}
	gopathEnv := os.Getenv("GOPATH")
	io.WriteString(w, fmt.Sprintf("os gopath env is %s, \n", gopathEnv))

	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
}

func healthz (w http.ResponseWriter, r *http.Request) {
	glog.V(2).Info("hello healthz")
	w.WriteHeader(200)
	io.WriteString(w, "ok")
}
