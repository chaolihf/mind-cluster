package npu_exporter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chaolihf/mind-cluster/component/ascend-common/common-utils/hwlog"
	"github.com/chaolihf/mind-cluster/component/ascend-common/devmanager"
)

var HwLogConfig = &hwlog.LogConfig{
	LogFileName:   "npu-exporter.log",
	ExpiredTime:   hwlog.DefaultExpiredTime,
	CacheSize:     hwlog.DefaultCacheSize,
	MaxBackups:    1,
	MaxAge:        7,
	MaxLineLength: 1024,
}

func NpuServer(server *http.Server) {
	if err := hwlog.InitRunLogger(HwLogConfig, context.Background()); err != nil {
		fmt.Printf("hwlog init failed, error is %v\n", err)
		return
	}
	dmgr, err := devmanager.AutoInit("")
	if err != nil {
		fmt.Printf("new npu collector failed, error is %v", err)
		return
	}
	cardNum, cardIDList, err := dmgr.DcMgr.DcGetCardList()
	if err != nil {
		fmt.Printf("error on init %s\n", err)
		return
	}
	fmt.Printf("cardNum %d, cardIDList %v\n", cardNum, cardIDList)
}
