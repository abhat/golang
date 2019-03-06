package main

import (
       "golang.org/x/tour/pic"
       "image"
       "image/color"
)

type Image struct{
     w uint8
     h uint8
}

func (i Image) ColorModel() color.Model {
     return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
     return image.Rect(0, 0, (int)(i.w), (int)(i.h))
}

func (i Image) At(x, y int) color.Color {
     return color.RGBA{uint8(x*x), uint8(y*y), uint8(x + 2*y), uint8(y + 2*x)}
}

func main() {
     m := Image{w: 255, h:255}
     pic.ShowImage(m)
}


