package main

import (
	"net/http"

	npu_exporter "github.com/chaolihf/mind-cluster/component/npu-exporter/cmd/npu-exporter"
)

func main() {
	server := &http.Server{
		Addr: "127.0.0.1:9100",
	}
	npu_exporter.NpuServer(server)
}
