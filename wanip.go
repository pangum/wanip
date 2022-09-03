package wanip

import (
	"github.com/goexl/wanip"
	"github.com/pangum/http"
	"github.com/pangum/logging"
)

type Agent struct {
	*wanip.Agent
}

func newAgent(_ *http.Client, _ *logging.Logger) (agent *Agent, err error) {
	agent = new(Agent)
	agent.Agent, err = wanip.NewAgent()

	return
}
