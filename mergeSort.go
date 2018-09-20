package main

import (
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"time"
)

func merge(win *pixelgl.Window, cfg *pixelgl.WindowConfig, imd *imdraw.IMDraw, arr []float64, l, m, r int) {
	limit := time.Tick(time.Second / 240)
	n1 := m - l + 1
	n2 := r - m

	L := make([]float64, n1)
	R := make([]float64, n2)

	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	var i, j, k= 0, 0, l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		<- limit
		draw(win, cfg, imd, &arr, &colMap)
		k++
	}

	for i < n1 {
		arr[k] = L[i]
		i++
		k++
		<- limit
		draw(win, cfg, imd, &arr, &colMap)
	}

	for j < n2 {
		arr[k] = R[j]
		j++
		k++
		<- limit
		draw(win, cfg, imd, &arr, &colMap)
	}
}

func mergeSort(win *pixelgl.Window, cfg *pixelgl.WindowConfig, imd *imdraw.IMDraw, arr []float64, l, r int) {
	if l < r {
		m := l + (r - l) / 2

		mergeSort(win, cfg, imd, arr, l, m)
		mergeSort(win, cfg, imd, arr, m + 1, r)

		merge(win, cfg, imd, arr, l, m ,r)
	}
}