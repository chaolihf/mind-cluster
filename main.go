package main

import (
	"net/http"

	"github.com/chaolihf/mind-cluster/component/npu_exporter/cmd/npu_exporter"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9100",
	}
	npu_exporter.NpuServer(server)
}
