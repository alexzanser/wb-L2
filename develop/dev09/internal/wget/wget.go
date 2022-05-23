package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gocolly/colly"
)

type Wget struct {
	basePath		string
	links			[]string
	client			*http.Client
	collector		*colly.Collector
}

func New() *Wget {
	return &Wget{
		links: []string{},
		client: &http.Client{},
		collector : colly.NewCollector(),
	}
}

func getPath(url string) string {
	dir := strings.TrimPrefix(url, "https://")
	dir = strings.TrimPrefix(dir, "http://")
	fmt.Println(dir)
	return  dir
}

func getFileName(url string) string {
	fname := path.Base(url)
	if strings.Contains(fname, ".") {
		if strings.Index(fname, "?v=") != -1 {
			return fname[:strings.Index(fname, "?v=")]
		}
	}
	return fname
}

func createDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (wget *Wget) GetPage(url string) error {
	c := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error %v when create request", err)
	} 

	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("error when send request: %v", err)
	}

	createDir(getPath(url))

	file, err  := os.Create(getPath(url) + "/" + getFileName(url))

	if err != nil {
		return fmt.Errorf("error when create file: %v", err)
	}
	io.Copy(file, resp.Body)

	return nil
}

func (wget *Wget) VisitAndGet(url string) {
	c := colly.NewCollector()
  
	c.OnHTML("a", func(e *colly.HTMLElement) {
	  e.Request.Visit(e.Attr("href"))
	})
  
	c.OnRequest(func(r *colly.Request) {
	 wget.GetPage(r.URL.String())
	})
 
	c.Visit(url)
 }
 
func main() {
	wget := New()
	wget.VisitAndGet("https://stackoverflow.com/questions/23166468/how-can-i-get-stdin-to-exec-cmd-in-golang")
}
