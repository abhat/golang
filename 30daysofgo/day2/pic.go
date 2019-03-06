package main
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	dyarr := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
	  dyarr[y] = make([]uint8, dx)
	  for x:= 0; x < dx; x++ {
		  dyarr[y][x] = uint8(x*x + y*y*y)
	  }
	}
	return dyarr
}

func main() {
	pic.Show(Pic)
}

