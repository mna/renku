package main

import (
	"fmt"
	"github.com/PuerkitoBio/ghost/handlers"
	"log"
	"net/http"
	"path"
	"time"
)

type logMode int

const (
	lmNormal logMode = iota
	lmQuiet
	lmVerbose

	faviconCacheTTL = 30 * 24 * time.Hour
)

type serverOptions struct {
	Port    int
	Root    string
	LogMode logMode
	Watch   bool
}

func servePage(w http.ResponseWriter, r *http.Request) {
	// TODO : Generate page from template
	http.Error(w, "Teapot", http.StatusTeapot)
}

func listenAndServe(opts serverOptions) {
	mux := http.NewServeMux()
	mux.Handle("/public", http.StripPrefix("/public",
		http.FileServer(http.Dir(path.Join(opts.Root, "public/")))))
	mux.HandleFunc("/", servePage)

	h := handlers.FaviconHandler(
		handlers.PanicHandler(
			mux,
			nil),
		path.Join(opts.Root, "public/favicon.ico"),
		faviconCacheTTL)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", opts.Port), h); err != nil {
		log.Fatal("^", err)
	}
}
