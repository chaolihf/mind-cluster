package main

/*
#cgo LDFLAGS: -L. -lnpumonitor
#include "libnpumonitor.h"

void NpuServer();
*/
import "C"

func main() {
	C.NpuServer()
}
