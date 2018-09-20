package main

import (
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

func swapWithDraw(win *pixelgl.Window,cfg *pixelgl.WindowConfig, imd *imdraw.IMDraw, arr []float64, colors *colorMap, i, j int) {
	limit := time.Tick(time.Second / 240)
	colors.special[i] = colornames.Blue
	colors.special[j] = colornames.Red

	arr[i], arr[j] = arr[j], arr[i]
	<- limit
	draw(win, cfg, imd, &arr, colors)
}

func partition (win *pixelgl.Window,cfg *pixelgl.WindowConfig, imd *imdraw.IMDraw, arr []float64, low, high int) int{
	pivot := arr[high]
	i := low - 1
	colMap.clearHighlight()
	for k := i; k < high; k++ {
		colMap.highlight[k] = colornames.Beige
	}
	for j := low; j <= high - 1; j++ {
		if arr[j] <= pivot {
			i++
			colMap.clear()
			swapWithDraw(win, cfg, imd, arr, &colMap, i, j)
		}
	}
	colMap.clear()
	swapWithDraw(win, cfg, imd, arr, &colMap, i + 1, high)
	return i + 1
}

func quickSort(win *pixelgl.Window, cfg *pixelgl.WindowConfig, imd *imdraw.IMDraw, arr []float64, low, high int) {
	if low < high {
		pi := partition(win, cfg, imd, arr, low, high)

		quickSort(win, cfg, imd, arr, low, pi - 1)
		quickSort(win, cfg, imd, arr, pi + 1, high)

	}
}