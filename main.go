package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

var (
	arr = genListToSort(nil, nil, nil)
	colMap = genColorMap(colornames.White)
)

var (
	frames = 0
	second = time.Tick(time.Second)

)

func draw(win *pixelgl.Window, cfg *pixelgl.WindowConfig, imd *imdraw.IMDraw, arr *[]float64, colors *colorMap) {
	win.Clear(colornames.Black)
	imd.Clear()
	genDrawableList(cfg, imd, *arr, colors)
	imd.Draw(win)
	win.Update()
	frames++
	select {
	case <-second:
		win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
		frames = 0
	default:
	}
}

func update() {
	cfg := pixelgl.WindowConfig{
		Title: "Visual-Sort",
		Bounds: pixel.R(0,0,1920,1080),
		Resizable: false,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	for !win.Closed() {
		for j := 0; j == 0; j--{
			arr = genListToSort(win, &cfg, imd)
			insertionSort(win, imd, &cfg, arr)
			<-second
			arr = genListToSort(win, &cfg, imd)
			bubbleSort(win, imd, &cfg, arr)

		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(update)
}
