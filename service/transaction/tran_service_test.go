package transaction

import (
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/dex-robot/service"
	"testing"
)

func TestTranHandle(t *testing.T) {
	pkg.Init("../../conf/config.yaml")
	service.InitSmart()
	TranHandle()
}
