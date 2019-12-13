# golang 爬取豆瓣top250电影名称



### 1. http抓取整个网页

```go
package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
)

func fetch(url string) string {
	fmt.Println("Fetch url:",url)

	client := &http.Client{}
	
	req,_ := http.NewRequest("GET",url,nil)

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

	resp,err := client.Do(req)

	if err != nil{
		fmt.Println("Http get err: ",err)
		return ""
	}

	if resp.StatusCode != 200 {
		fmt.Println("Http status code:",resp.StatusCode)
		return ""
	}

	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error:",err)
		return ""
	}

	return string(body)
}
```



### 2. 正则表达式解析拿回来的字符串

```go
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

```



### 3. 测试

```go
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
```





