package wanip

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newAgent,
	)
}
