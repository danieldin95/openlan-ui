package api

import (
	"github.com/danieldin95/openlan-ui/backend/ctl"
	"github.com/danieldin95/openlan-ui/backend/schema"
	"github.com/danieldin95/openlan-ui/backend/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Graph struct {
}

func (g Graph) Router(router *mux.Router) {
	router.HandleFunc("/api/graph/{id}", g.GET).Methods("GET")
}

func (g Graph) GET(w http.ResponseWriter, r *http.Request) {
	//id, _ := GetArg(r, "id")
	gs := struct {
		Categories []schema.Category `json:"categories"`
		Nodes      []schema.Node     `json:"nodes"`
		Links      []schema.Link     `json:"links"`
	}{
		Categories: []schema.Category{
			{
				Name: "virtual switch",
			},
			{
				Name: "accessed point",
			},
		},
		Nodes: make([]schema.Node, 0, 32),
		Links: make([]schema.Link, 0, 32),
	}

	i := 0
	idx := make(map[string]int, 32)
	for vs := range service.SERVICE.VSwitch.List() {
		if vs == nil {
			break
		}
		idx[vs.Name] = i
		gs.Nodes = append(gs.Nodes, schema.Node{
			Name:       vs.Name,
			SymbolSize: 20,
			Category:   0,
			Id:         i,
			Label: &schema.Label{
				Show: true,
			},
		})
		i += 1
	}
	for vs := range service.SERVICE.VSwitch.List() {
		if vs == nil {
			break
		}
		vid, ok := idx[vs.Name]
		if !ok {
			continue
		}
		vc, ok := vs.Ctl.(*ctl.VSwitch)
		if !ok {
			continue
		}
		for p := range vc.ListPoint() {
			if p == nil {
				break
			}
			pid, ok := idx[p.Alias]
			if !ok {
				pid = i
				idx[p.Alias] = i
				gs.Nodes = append(gs.Nodes, schema.Node{
					Name:       p.Alias,
					SymbolSize: 10,
					Category:   1,
					Id:         pid,
				})
				i += 1
			}
			gs.Links = append(gs.Links, schema.Link{
				Source: pid,
				Target: vid,
				Weight: i,
			})
		}
		// TODO links.
	}
	ResponseJson(w, gs)
}

func (g Graph) POST(w http.ResponseWriter, r *http.Request) {
	ResponseMsg(w, 0, "")
}

func (g Graph) PUT(w http.ResponseWriter, r *http.Request) {
	ResponseMsg(w, 0, "")
}

func (g Graph) DELETE(w http.ResponseWriter, r *http.Request) {
	ResponseMsg(w, 0, "")
}
