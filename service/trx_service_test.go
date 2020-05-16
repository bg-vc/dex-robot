package service

import (
	"github.com/vincecfl/dex-robot/pkg"
	"testing"
)

func TestTransferTrxHandle(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	InitSmart()
	TransferTrxHandle()
}
