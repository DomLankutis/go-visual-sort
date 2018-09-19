package main

import (
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

func bubbleSort(win *pixelgl.Window, imd *imdraw.IMDraw, cfg *pixelgl.WindowConfig, arr []float64) {
	limit := time.Tick(time.Second / 1000)
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - i - 1; j++ {
			colMap.clear()
			if arr[j] > arr[j + 1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				colMap.special[j + 1] = colornames.Red
			}
			<-limit
			draw(win, cfg, imd, &arr, &colMap)
		}
	}
}