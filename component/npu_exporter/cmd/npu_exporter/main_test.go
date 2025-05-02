package main

import (
	"net/http"
	"testing"

	_ "huawei.com/npu-exporter/v6/plugins/inputs/npu"
)

func TestNpuServer(t *testing.T) {
	server := http.Server{
		Addr: "127.0.0.1:9100",
	}
	NpuServer(server)
}
