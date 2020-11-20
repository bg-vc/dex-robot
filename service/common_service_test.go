package service

import (
	"github.com/vincecfl/dex-robot/pkg"
	"testing"
)

func TestCancel(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	InitSmart()
	Cancel("TPdBHYrTDiop2fgsmZGDEfNN5SucJADCf4", "514bfc62a1f84b69a46ba6478f991eacb136ef1a2f63a16a66e7f42c14c1de07", 696457)
}
