package main

import (
	"main/cpu"
	"main/display"
)

var cpuInstance cpu.CpuInstance
var displayInstance display.DisplayInstance

func main() {
	cpuInstance.Init()
	displayInstance.Init()
	displayInstance.Terminate()
	cpuInstance.Terminate()
}
