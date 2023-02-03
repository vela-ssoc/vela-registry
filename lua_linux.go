package registry

import (
	"github.com/vela-ssoc/vela-kit/vela"
)

func WithEnv(env vela.Environment) {
	env.Error("not support registry with linux")
}
