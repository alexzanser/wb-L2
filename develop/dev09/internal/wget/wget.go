package wget

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"log"
	"github.com/gocolly/colly"
)

type Wget struct {
	BasePath		string
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
	return path.Base(url)
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

	createDir(wget.BasePath + "/"+ getPath(url))

	file, err  := os.Create(wget.BasePath + "/" + getPath(url) + "/" + getFileName(url))

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
 
 func (wget *Wget) Start() {
	if len(os.Args) < 2 {
		log.Fatal("wget: missing URL")
	}

	url := os.Args[1]

	if len(os.Args) == 3 {
		wget.BasePath = os.Args[2]
	}

	wget.VisitAndGet(url)
 }