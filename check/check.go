package check

import (
	"go-dirsearch/request"
	"go-dirsearch/assets"
	"bufio"
	"fmt"
	"strings"
)

var statuscode int


func Getline()(lines []string){
	ass := assets.Dict
	file, err := ass.Open("minwpdict.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		//fmt.Println(lines)
		return lines
}

func Check(url string ,UA []string){
	lines := Getline()
	for _, line := range lines {
		strHaiCoder := `%EXT%`
		StrContainers := strings.Contains(line , strHaiCoder)
		if StrContainers == true {
			continue
		}
		statuscode = request.Got_statuscode(url+line,UA)
		if statuscode == 200 {
			fmt.Println(url+line)
		}
	}
}
