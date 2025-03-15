package main

import (
	"main/cpu"
	"main/display"

	"github.com/veandco/go-sdl2/sdl"
)

var cpuInstance cpu.CpuInstance
var displayInstance display.DisplayInstance

func main() {
	cpuInstance.Init()
	displayInstance.Init()
	sdl.Delay(3000)
	displayInstance.Terminate()
	cpuInstance.Terminate()
}
