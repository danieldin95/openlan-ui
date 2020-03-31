package service

import (
	"github.com/danieldin95/lightstar/libstar"
	"github.com/danieldin95/openlan-ui/backend/schema"
)

type Service struct {
	Users   Users
	VSwitch VSwitch
}

var SERVICE = Service{
	Users: Users{
		Users: make(map[string]*schema.User, 32),
	},
	VSwitch: VSwitch{
		VSwitch: make(map[string]*schema.VSwitch, 32),
	},
}

func (s *Service) Load(path string) {
	if err := s.Users.Load(path + "/auth.json"); err != nil {
		libstar.Error("Service.Load.Users %s", err)
	}
	libstar.Debug("Service.Load %s", s.Users)
	if err := s.VSwitch.Load(path + "/vswitch.json"); err != nil {
		libstar.Error("Service.Load.VSwitch %s", err)
	}
	libstar.Debug("Service.Load %s", s.VSwitch)
}
