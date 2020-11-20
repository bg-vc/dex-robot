package task

import (
	"github.com/vincecfl/dex-robot/service/transaction"
)

func TranHandle() {
	go transaction.TranHandle()
}
