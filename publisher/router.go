package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	pubctrl "../controllers"
	"../helpers"
	"github.com/go-chi/chi"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {

	basePath := "/"

	r := chi.NewRouter()
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("1"))
	})

	requestUrls := helpers.GenerateUrls()

	pubCtrl := pubctrl.PubController{}

	for _, requestURL := range requestUrls {
		r.Post(basePath+requestURL, pubCtrl.Handle)
	}

	fmt.Println("Route list :\n", strings.Join(requestUrls, "\n"))

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
