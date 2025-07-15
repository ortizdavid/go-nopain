package main

import (
	"fmt"
	"log"
	"github.com/ortizdavid/go-nopain/httputils/client"
)

func main() {
	client := httputils.NewHttpClient()
	client.SetHeader("X-USER-ID", "fdeb15cb-3431-47e1-aed6-72d6aea822f0")
	client.SetHeader("X-API-KEY", "4--GyiF-grNz7vpeD1u1WfkMcC_nsnkNiK3xQMgUskgIyFH6IMvCSSS_CaSkzYtyUIZi7AKKtHJGRdAJsTiXiMgayza0UGkRTPWTuxf9OW5BmApXJZbWnL-kM8_rLFWMiAlWsw")
	
	url := "http://127.0.0.1:4003/api/configurations/basic-configurations"
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(string(resp.Body))
}