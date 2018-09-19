package main

import (
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

func insertionSort(win *pixelgl.Window, imd *imdraw.IMDraw, cfg *pixelgl.WindowConfig, arr []float64) {
	limit := time.Tick(time.Second / 1000)
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j + 1] = arr[j]
			colMap.clear()
			colMap.special[j - 1] = colornames.Red
			j = j - 1
			<- limit
			draw(win, cfg, imd, &arr, &colMap)
		}
		arr[j + 1] = key
	}
}
