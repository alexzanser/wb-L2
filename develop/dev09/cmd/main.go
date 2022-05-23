package main

import (
	"wget/internal/wget"
)
func main() {
	w := wget.New()
	w.basePath = "data/"
	w.VisitAndGet("https://stackoverflow.com/questions/23166468/how-can-i-get-stdin-to-exec-cmd-in-golang")
}
