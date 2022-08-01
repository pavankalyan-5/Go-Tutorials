package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://leetcode.com/problemset/algorithms/?status=NOT_STARTED&page=1"

func main() {
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	//fmt.Println(result.Port())

	qparams := result.Query()

	for key, val := range qparams {
		fmt.Printf("%v\t%v\n", key, val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "leetcode.com",
		Path:     "/problemset/algorithms",
		RawQuery: "status=NOT_STARTED",
	}
	AnotherUrl := partsOfUrl.String()
	fmt.Println(AnotherUrl)
}
