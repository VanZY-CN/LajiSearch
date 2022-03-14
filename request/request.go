package request

import (
	"github.com/go-resty/resty/v2"
	"go-dirsearch/assets"
	"bufio"
	"math/rand"
	//"fmt"
)



func Got_statuscode(url string,UA []string) (statuscode int){
	client := resty.New()
	resp, _ := client.R().
		EnableTrace().
		SetHeader("User_Agent", get_random_ua(UA)). 
  		Get(url)
	statuscode = resp.StatusCode()
	//fmt.Println(statuscode)
	return statuscode
}

func Get_ua() []string {
	ass := assets.Dict
	file1, _ := ass.Open("ua.txt")
	defer file1.Close()
	scanner := bufio.NewScanner(file1)
	scanner.Split(bufio.ScanLines)
	var UA []string
	for scanner.Scan() {
		UA = append(UA, scanner.Text())
	}
	return UA
}

func get_random_ua(UA []string) string {
	length := len(UA)
	index := rand.Intn(length)
	//fmt.Println(UA[index])
	return UA[index]
}