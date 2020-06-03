package api

import (
	"encoding/json"
	"github.com/go-clog/clog"
	"log"
	"net/http"

	"github.com/henson/proxypool/pkg/setting"
	"github.com/henson/proxypool/pkg/storage"
)

// VERSION for this program
const VERSION = "/v2"

// Run for request
func Run() {

	mux := http.NewServeMux()
	mux.HandleFunc(VERSION+"/ip", ProxyHandler)
	mux.HandleFunc(VERSION+"/https", FindHandler)
	log.Println("Starting server", setting.AppAddr+":"+setting.AppPort)
	http.ListenAndServe(setting.AppAddr+":"+setting.AppPort, mux)
}

// ProxyHandler .
func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		b, err := json.Marshal(storage.ProxyRandom())
		if err != nil {
			clog.Info("[api] get ip err %v", err)
			return
		}
		w.Write(b)
	}
}

// FindHandler .
func FindHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		b, err := json.Marshal(storage.ProxyFind("https"))
		if err != nil {
			clog.Info("[api] get https err %v", err)
			return
		}
		w.Write(b)
	}
}
