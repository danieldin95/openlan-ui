package main

import (
	"flag"
	"github.com/danieldin95/lightstar/libstar"
	"github.com/danieldin95/openlan-ui/backend/ctl"
	"github.com/danieldin95/openlan-ui/backend/http"
	"github.com/danieldin95/openlan-ui/backend/service"
	"os"
)

type Config struct {
	StaticDir string `json:"dir.static"`
	CrtDir    string `json:"dir.crt"`
	ConfDir   string `json:"-"`
	Verbose   int    `json:"log.level"`
	LogFile   string `json:"log.file"`
	Listen    string `json:"listen"`
}

var cfg = Config{
	StaticDir: "static",
	CrtDir:    "ca",
	ConfDir:   "/etc/openlan-ui",
	Listen:    "0.0.0.0:10088",
	LogFile:   "/var/log/lightstar.log",
	Verbose:   2,
}

func main() {
	flag.StringVar(&cfg.Listen, "listen", cfg.Listen, "the address http listen.")
	flag.IntVar(&cfg.Verbose, "log:level", cfg.Verbose, "logger level")
	flag.StringVar(&cfg.CrtDir, "crt:dir", cfg.CrtDir, "the directory X509 certificate file on.")
	flag.StringVar(&cfg.StaticDir, "static:dir", cfg.StaticDir, "the directory to serve files from.")
	flag.StringVar(&cfg.ConfDir, "conf", cfg.ConfDir, "the directory configuration on")
	flag.Parse()

	libstar.Init(cfg.LogFile, cfg.Verbose)
	service.SERVICE.Load(cfg.ConfDir)
	ctl.CTL.Load(&service.SERVICE)

	h := http.NewServer(cfg.Listen, cfg.StaticDir)
	if _, err := os.Stat(cfg.CrtDir); !os.IsNotExist(err) {
		h.SetCert(cfg.CrtDir+"/private.key", cfg.CrtDir+"/crt.pem")
	}
	go h.Start()

	libstar.Wait()
}
