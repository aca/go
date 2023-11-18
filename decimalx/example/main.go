package main

import (
	"fmt"

	"github.com/aca/go/decimalx"
)

func main(){
	startx := decimalx.New(12701, -2)
	endx := decimalx.New(12704, -2)
	gapx := decimalx.New(1, -2)

	starty := decimalx.New(3701, -2)
	endy := decimalx.New(3704, -2)
	gapy := decimalx.New(1, -2)

	for curx := startx; curx.IsSmaller(endx); curx = curx.Add(gapx) {
		for cury := starty; cury.IsSmaller(endy); cury = cury.Add(gapy) {
			fmt.Println(curx, cury)
		}
	}
}
