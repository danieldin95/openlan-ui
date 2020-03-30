package api

import (
	"github.com/danieldin95/openlan-ui/backend/schema"
	"github.com/danieldin95/openlan-ui/backend/service"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
)

type User struct {
}

func (u User) Router(router *mux.Router) {
	router.HandleFunc("/api/user", u.GET).Methods("GET")
}

func (u User) GET(w http.ResponseWriter, r *http.Request) {
	us := make([]schema.VSwitch, 0, 32)
	for h := range service.SERVICE.Users.List() {
		if h == nil {
			break
		}
		us = append(us, *h)
	}
	sort.SliceStable(us, func(i, j int) bool {
		return us[i].Name < us[j].Name
	})
	ResponseJson(w, us)
}

func (u User) POST(w http.ResponseWriter, r *http.Request) {
	ResponseMsg(w, 0, "")
}

func (u User) PUT(w http.ResponseWriter, r *http.Request) {
	ResponseMsg(w, 0, "")
}

func (u User) DELETE(w http.ResponseWriter, r *http.Request) {
	ResponseMsg(w, 0, "")
}
