package main

import (
	"github.com/yansong2022/thirdpackagetest/mymethod"
	"github.com/yansong2022/thirdpackagetest/mystruct"
)

func main() {
	mytest := mystruct.NewService("first server", "ys", 3)
	mytest.PrintService()
	mymethod.TestMethod()
}
