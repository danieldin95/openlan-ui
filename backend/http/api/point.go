package api

import (
	"github.com/danieldin95/openlan-ui/backend/ctl"
	"github.com/danieldin95/openlan-ui/backend/schema"
	"github.com/danieldin95/openlan-ui/backend/service"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
)

type Point struct {
}

func (p Point) Router(router *mux.Router) {
	router.HandleFunc("/api/point/{id}", p.GET).Methods("GET")
}

func (p Point) GET(w http.ResponseWriter, r *http.Request) {
	id, _ := GetArg(r, "id")
	vs, ok := service.SERVICE.VSwitch.Get(id)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	vc, ok := vs.Ctl.(*ctl.VSwitch)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	ps := make([]schema.Point, 0, 32)
	for h := range vc.ListPoint() {
		if h == nil {
			break
		}
		ps = append(ps, *h)
	}
	sort.SliceStable(ps, func(i, j int) bool {
		return ps[i].Alias < ps[j].Alias
	})
	ResponseJson(w, ps)
}

func (p Point) POST(w http.ResponseWriter, r *http.Request) {
	ResponseMsg(w, 0, "")
}

func (p Point) PUT(w http.ResponseWriter, r *http.Request) {
	ResponseMsg(w, 0, "")
}

func (p Point) DELETE(w http.ResponseWriter, r *http.Request) {
	ResponseMsg(w, 0, "")
}