package main

import (
	"fmt"
	"log"
	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(fmt.Errorf("Error while getting time: %v", err))
	}
	fmt.Println(time)
}
