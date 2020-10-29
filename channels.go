package main

import (
	"fmt"
	"net/http"
	"strings"
)

//ChannelsMain func
func ChannelsMain() {
	sites := strings.Split("yahoo,google,facebook,twitter,instagram,whatsapp", ",")
	c := make(chan string)
	for _, site := range sites {
		go checkLink(fmt.Sprintf("https://%s.com", site), c)
	}
	i := 0
	for i < len(sites)-1 {
		fmt.Println(<-c)
		i++
	}
	close(c)
}

func checkLink(url string, c chan string) {
	res, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("Error %v", url)
	} else {
		c <- fmt.Sprintf("A ok %v, %v", url, res.StatusCode)
	}
}
