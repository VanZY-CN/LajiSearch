package main

import (
	"go-dirsearch/check"
	"go-dirsearch/myflag"
	"go-dirsearch/request"
	"flag"
	"os"
	"fmt"
	"bufio"
	"sync"
	"github.com/panjf2000/ants/v2"
)

var URL = flag.String("url", "", "input url")
var Urllist = flag.String("file", "", "input path to urllist")
var UA []string
type taskFunc func()
var Wg sync.WaitGroup

func main(){
	myflag.Banner()
	defer ants.Release()
	UA=request.Get_ua()
	flag.Parse()
	if *URL != "" {
		check.Check(*URL,UA)
	}
	if *Urllist != "" {
		file, err := os.Open(*Urllist)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		
		p, _ := ants.NewPool(10)
		defer p.Release()
		for _, line := range lines {
			Wg.Add(1)
			p.Submit(taskFuncWrapper(line,UA,&Wg))
		}
	}
	Wg.Wait()
}

func taskFuncWrapper(line string, UA []string, wg *sync.WaitGroup) taskFunc {
  return func() {
		check.Check(line,UA)
    	wg.Done()
  }
}