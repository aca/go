package main

import (
	"github.com/WAY29/icecream-go/icecream"
	"github.com/aca/go/q"
)

func main(){
	x := []int{1,4}
	q.Q("hello", x)

    // icecream.ConfigureIncludeContext(true)
	icecream.ConfigurePrefix("")
	icecream.Ic(x)
    icecream.Ic(1, "asd")
}
