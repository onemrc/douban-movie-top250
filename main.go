package main

import(
	"time"
	"strconv"
	"fmt"
)

var idx int

func main()  {
	start := time.Now()

	idx = 1

	// 遍历
	for i := 0;i < 10;i++{
		parseUrl("https://movie.douban.com/top250?start=" + strconv.Itoa(25 * i))
	}

	elapsed := time.Since(start)

	fmt.Printf("Took %s",elapsed)
}