package check

import (
	"go-dirsearch/request"
	"go-dirsearch/assets"
	"bufio"
	"fmt"
	"strings"
	"os"
	"strconv"
	"time"
)

var statuscode int
type Multimap map[string][]string

type keyValues struct {
       key    string
       values []string
}

func (multimap Multimap) Add(key, value string) {
	if len(multimap[key]) == 0 {
		   multimap[key] = []string{value}
	} else {
		   multimap[key] = append(multimap[key], value)
	}
}

func (multimap Multimap) Get(key string) []string {
	if multimap == nil {
		   return nil
	}
	values := multimap[key]
	return values
}

func Getline(Dict string)(lines []string){
	if Dict == "" {
		ass := assets.Dict
		file, err := ass.Open("dict.txt")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
			scanner1 := bufio.NewScanner(file)
			scanner1.Split(bufio.ScanLines)
			for scanner1.Scan() {
				lines = append(lines, scanner1.Text())
			}
			//fmt.Println(lines)
			return lines
	} else {
		file, err := os.Open(Dict)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		scanner2 := bufio.NewScanner(file)
		scanner2.Split(bufio.ScanLines)
		var lines []string
		for scanner2.Scan() {
			lines = append(lines, scanner2.Text())
		}
		return lines
	}
}

func Check(url string ,UA []string ,Dict string,File bool){
	lines := Getline(Dict)
	var myMap Multimap
	myMap = make(Multimap)
	for _, line := range lines {
		strHaiCoder := `%EXT%`
		StrContainers := strings.Contains(line , strHaiCoder)
		if StrContainers == true {
			continue
		}
		statuscode = request.Got_statuscode(url+line,UA)
		if statuscode == 200 {
			//fmt.Println(url+line)
			myMap.Add(url,url+line)
		}
	}
	if File == false {
		for key := range myMap {
			fmt.Printf("\n%s:\n",key)
			for j := 0; j < len(myMap[key]); j++ {
				fmt.Println(myMap[key][j])
			}
		}
	}	else {
		year := time.Now().Year()
		years := strconv.Itoa(int(year))
		month := time.Now().Month()
		months := strconv.Itoa(int(month))
		day := time.Now().Day()
		days := strconv.Itoa(int(day))
		var filepath string
		filepath = years + "." + months + "." + days + ".txt"
		file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("文件打开失败", err)
		}
		defer file.Close()
		write := bufio.NewWriter(file)
		for key := range myMap {
			write.WriteString("\n"+key+":\n")
			for j := 0; j < len(myMap[key]); j++ {
				write.WriteString(myMap[key][j]+"\n")
			}
		}
		write.Flush()
		fmt.Printf("Done!,please see the reselt in %s",filepath)
	}	
}
