// https://go.dev/tour/moretypes/18

/*
package main

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		res[y] = row
		for x := 0; x < dx; x++ {
			res[y][x] = uint8((x ^ y))
		}
	}

	return res
}

func main() {
	pic.Show(Pic)
}
*/