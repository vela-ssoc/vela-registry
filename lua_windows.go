//go:build windows
// +build windows

package registry

import (
	"github.com/vela-ssoc/vela-kit/vela"
	"github.com/vela-ssoc/vela-kit/lua"
)

//rock.registry.LOCAL_MACHINE.create("SOFTWARE\\Hello Go" , "ALL_ACCESS")

var xEnv vela.Environment

func WithEnv(env vela.Environment) {
	xEnv = env

	r := lua.NewUserKV()
	r.Set("USERS", registryL("USERS"))
	r.Set("LOCAL_MACHINE", registryL("LOCAL_MACHINE"))
	r.Set("CLASSES_ROOT", registryL("CLASSES_ROOT"))
	r.Set("CURRENT_CONFIG", registryL("CURRENT_CONFIG"))
	r.Set("CURRENT_USER", registryL("CURRENT_USER"))
	r.Set("PERFORMANCE_DATA", registryL("PERFORMANCE_DATA"))
	xEnv.Set("registry", r)
}
