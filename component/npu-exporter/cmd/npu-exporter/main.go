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
	deviceManager := dmgr.DcMgr
	cardNum, cardIDList, err := deviceManager.DcGetCardList()
	if err != nil {
		fmt.Printf("error on init %s\n", err)
		return
	}
	fmt.Printf("cardNum %d, cardIDList %v\n", cardNum, cardIDList)
	for _, cardID := range cardIDList {
		deviceNum, err := deviceManager.DcGetDeviceNumInCard(cardID)
		if err != nil {
			fmt.Printf("error on get device id %s\n", err)
			return
		}
		fmt.Printf("cardID %d, deviceNum %d\n", cardID, deviceNum)
		for deviceID := 0; deviceID < int(deviceNum); deviceID++ {
			logicID, err := deviceManager.DcGetDeviceLogicID(cardID, int32(deviceID))
			if err != nil {
				fmt.Printf("error on get device logic id %s\n", err)
				return
			}
			fmt.Printf("cardID %d, deviceID %d,logicId %d\n", cardID, deviceID, logicID)
			newCardId, newDeviceId, err := deviceManager.DcGetCardIDDeviceID(logicID)
			if err != nil {
				fmt.Printf("error on get device real device id %s\n", err)
				return
			}
			fmt.Printf("new cardID %d, deviceID %d,logicId %d\n", newCardId, newDeviceId, logicID)

			voltageInfo, err := deviceManager.DcGetDeviceVoltage(cardID, int32(deviceID))
			if err != nil {
				fmt.Printf("error on get device voltageInfo id %s\n", err)
				return
			}
			fmt.Printf("cardID %d, deviceID %d,voltageInfo %v\n", cardID, deviceID, voltageInfo)

			powerInfo, err := deviceManager.DcGetDevicePowerInfo(cardID, int32(deviceID))
			if err != nil {
				fmt.Printf("error on get device powerInfo id %s\n", err)
				return
			}
			fmt.Printf("cardID %d, deviceID %d,powerInfo %v\n", cardID, deviceID, powerInfo)

			temperatureInfo, err := deviceManager.DcGetDeviceTemperature(cardID, int32(deviceID))
			if err != nil {
				fmt.Printf("error on get device temperatureInfo id %s\n", err)
				return
			}
			fmt.Printf("cardID %d, deviceID %d,temperatureInfo %v\n", cardID, deviceID, temperatureInfo)

			highBandwidthMemoryInfo, err := deviceManager.DcGetHbmInfo(cardID, int32(deviceID))
			if err != nil {
				fmt.Printf("error on get device highBandwidthMemoryInfo id %s\n", err)
				return
			}
			fmt.Printf("cardID %d, deviceID %d,highBandwidthMemoryInfo %v\n", cardID, deviceID, highBandwidthMemoryInfo)

			devProcessInfo, err := deviceManager.DcGetDevProcessInfo(cardID, int32(deviceID))
			if err != nil {
				fmt.Printf("error on get device devProcessInfo id %s\n", err)
				return
			}
			fmt.Printf("cardID %d, deviceID %d,devProcessInfo %v\n", cardID, deviceID, devProcessInfo)

			memoryInfo, err := deviceManager.DcGetMemoryInfo(cardID, int32(deviceID))
			if err != nil {
				fmt.Printf("error on get device memory id %s\n", err)
				//return
			}
			fmt.Printf("cardID %d, deviceID %d,memory %v\n", cardID, deviceID, memoryInfo)
		}
	}
}
