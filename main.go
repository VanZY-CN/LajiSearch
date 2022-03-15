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
var Dict = flag.String("dict" , "" , "input path of your DIY dict")
var File = flag.Bool("tofile" , false , "see the result in File or powershell or file")
var UA []string
type taskFunc func()
var Wg sync.WaitGroup

func main(){
	myflag.Banner()
	defer ants.Release()
	UA=request.Get_ua()
	flag.Parse()
	if *URL != "" {
		check.Check(*URL,UA,*Dict,*File)
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
			p.Submit(taskFuncWrapper(line,UA,*Dict,*File,&Wg))
		}
	}
	Wg.Wait()
}

func taskFuncWrapper(line string , UA []string , Dict string , File bool , wg *sync.WaitGroup) taskFunc {
  return func() {
		check.Check(line,UA,Dict,File)
    	wg.Done()
  }
}