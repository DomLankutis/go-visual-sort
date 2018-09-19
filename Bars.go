package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"math"
	"math/rand"
	"time"
)

type colorMap struct {
	def			color.RGBA
	special 	map[int]color.RGBA
}

func (c *colorMap) clear() {
	c.special = map[int]color.RGBA{}
}

func genListToSort(win *pixelgl.Window, cfg *pixelgl.WindowConfig, imd *imdraw.IMDraw) []float64 {
	n := 360
	radd := (2 * math.Pi) / float64(n)
	limit := time.Tick(time.Second / 120)
	arr := make([]float64, n)
	for i := range arr{
		arr[i] = float64(i) * float64(radd)
	}
	rand.Seed(time.Now().UnixNano())
	for n := len(arr); n > 0; n-- {
		randIndex := rand.Intn(n)
		arr[n-1], arr[randIndex] = arr[randIndex], arr[n-1]
		if win != nil {
			colMap.clear()
			<-limit
			draw(win, cfg, imd, &arr, &colMap)
		}
	}
	return arr
}

func genColorMap(def color.RGBA) colorMap {
	return colorMap{def, make(map[int]color.RGBA)}
}

func genDrawableList(cfg *pixelgl.WindowConfig, imd *imdraw.IMDraw, arr []float64, colors *colorMap) {
	midPoint := cfg.Bounds.Center()
	centerGap := cfg.Bounds.Size().Y / 100 * math.Sqrt(float64(len(arr)))
	thickness := cfg.Bounds.Size().Y / 100 / math.Sqrt(float64(len(arr)) / 5)
	for i, rad := range arr {
		barLength := rad * cfg.Bounds.Size().Y / math.Sqrt(float64(len(arr)))
		if color, ok := colors.special[i]; ok {
			imd.Color = color
		}else {
			imd.Color = colors.def
		}
		//// Alternative Visualisation
		//px1 := midPoint.X + (centerGap * math.Cos(float64(i) * 0.07))
		//py1 := midPoint.Y + (centerGap * math.Sin(float64(i) * 0.07))
		//px2 := midPoint.X + ((centerGap + barLength) * math.Cos(rad))
		//py2 := midPoint.Y + ((centerGap + barLength) * math.Sin(rad))
		////
		px1 := midPoint.X + (centerGap * math.Cos(rad))
		py1 := midPoint.Y + (centerGap * math.Sin(rad))
		px2 := midPoint.X + ((centerGap + barLength) * math.Cos(float64(i) * 0.07))
		py2 := midPoint.Y + ((centerGap + barLength) * math.Sin(float64(i) * 0.07))
		imd.Push(pixel.V(px1, py1), pixel.V(px2, py2))
		imd.Line(thickness)
	}
}

