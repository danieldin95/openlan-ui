package http

import (
	"context"
	"github.com/danieldin95/lightstar/libstar"
	"github.com/danieldin95/openlan-ui/backend/http/api"
	"github.com/danieldin95/openlan-ui/backend/service"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"strings"
)

type Server struct {
	listen  string
	server  *http.Server
	crtFile string
	keyFile string
	pubDir  string
	router  *mux.Router
	token   string
}

func NewServer(listen, staticDir string) (h *Server) {
	h = &Server{
		listen: listen,
		pubDir: staticDir,
	}
	return
}

func (h *Server) Router() *mux.Router {
	if h.router != nil {
		return h.router
	}
	h.router = mux.NewRouter()
	return h.router
}

func (h *Server) LoadRouter() {
	router := h.Router()
	router.Use(h.Middleware)

	// Api router
	api.VSwitch{}.Router(router)
	api.User{}.Router(router)
	api.Point{}.Router(router)
	api.Link{}.Router(router)
	api.Graph{}.Router(router)
	// static files
	Dist{h.pubDir}.Router(router)
}

func (h *Server) SetCert(keyFile, crtFile string) {
	h.crtFile = crtFile
	h.keyFile = keyFile
}

func (h *Server) Initialize() {
	if h.token == "" {
		h.token = libstar.GenToken(32)
	}
	libstar.Info("Server.Initialize token %s", h.token)
	r := h.Router()
	if h.server == nil {
		h.server = &http.Server{
			Addr:    h.listen,
			Handler: r,
		}
	}
	h.LoadRouter()
}

func (h *Server) IsAuth(w http.ResponseWriter, r *http.Request) bool {
	name, pass, _ := api.GetAuth(r)
	libstar.Print("Server.IsAuth %s:%s", name, pass)

	user, ok := service.SERVICE.Users.Get(name)
	if !ok || user.Password != pass {
		return false
	}
	return true
}

func (h *Server) LogRequest(r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/js") ||
		strings.HasPrefix(r.URL.Path, "/css") ||
		strings.HasPrefix(r.URL.Path, "/dist") ||
		strings.HasPrefix(r.URL.Path, "/fonts") ||
		strings.HasSuffix(r.URL.Path, ".ico") ||
		strings.HasSuffix(r.URL.Path, ".png") ||
		strings.HasSuffix(r.URL.Path, ".gif") {
		return
	}
	path := r.URL.Path
	if q, _ := url.QueryUnescape(r.URL.RawQuery); q != "" {
		path += "?" + q
	}
	libstar.Info("Server.Middleware %s %s %s", r.RemoteAddr, r.Method, path)
}

func (h *Server) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.LogRequest(r)
		if h.IsAuth(w, r) {
			next.ServeHTTP(w, r)
		} else {
			w.Header().Set("WWW-Authenticate", "Basic")
			http.Error(w, "Authorization Required.", http.StatusUnauthorized)
		}
	})
}

func (h *Server) Start() error {
	h.Initialize()
	if h.keyFile == "" || h.crtFile == "" {
		libstar.Info("Server.Start http://%s", h.listen)
		if err := h.server.ListenAndServe(); err != nil {
			libstar.Error("Server.Start on %s: %s", h.listen, err)
			return err
		}
	} else {
		libstar.Info("Server.Start https://%s", h.listen)
		if err := h.server.ListenAndServeTLS(h.crtFile, h.keyFile); err != nil {
			libstar.Error("Server.Start on %s: %s", h.listen, err)
			return err
		}
	}
	return nil
}

func (h *Server) Shutdown() {
	libstar.Info("Server.Shutdown %s", h.listen)
	if err := h.server.Shutdown(context.Background()); err != nil {
		libstar.Error("Server.Shutdown %v", err)
	}
}
