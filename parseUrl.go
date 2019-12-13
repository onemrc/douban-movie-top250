package main

import(
	"regexp"
	"strings"
	"fmt"
)


func parseUrl(url string)  {
	body := fetch(url)

	body = strings.Replace(body,"\n","",-1)

	rp := regexp.MustCompile(`<div class="hd">(.*?)</div>`)

	titleRe := regexp.MustCompile(`<span class="title">(.*?)</span>`)

	items := rp.FindAllStringSubmatch(body,-1)

	for _,item := range items{
		// fmt.Println(item[1])

		fmt.Println(idx," ",titleRe.FindStringSubmatch(item[0])[1])

		idx += 1
	}
}
