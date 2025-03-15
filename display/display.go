package display

import (
	"github.com/veandco/go-sdl2/sdl"
)

const pixelSizeScaleX = 5
const pixelSizeScaleY = 5

const pixelWidth = 5
const pixelHeight = 5

type DisplayInstance struct {
	width         int
	height        int
	displayRects  [][]sdl.Rect
	window        *sdl.Window
	windowSurface *sdl.Surface
}

func (display *DisplayInstance) Init() {
	display.width = 160
	display.height = 144

	display.window, _ = sdl.CreateWindow("Jameboy by JinnFighter", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, int32(display.width*pixelSizeScaleX), int32(display.height*pixelSizeScaleY), sdl.WINDOW_SHOWN)
	display.windowSurface, _ = display.window.GetSurface()
	display.windowSurface.FillRect(nil, 0)

	display.displayRects = make([][]sdl.Rect, display.width)
	for i := range display.width {
		display.displayRects[i] = make([]sdl.Rect, display.height)
		for j := range display.height {
			display.displayRects[i][j] = sdl.Rect{X: int32(i * pixelSizeScaleX), Y: int32(j * pixelSizeScaleY), W: pixelWidth, H: pixelHeight}
		}
	}
}

func (display *DisplayInstance) Terminate() {
	display.window.Destroy()
}
