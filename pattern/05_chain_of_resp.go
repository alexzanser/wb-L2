package main

import "fmt"

type handler interface {
	sendRequest(string) string
}

type handlerA struct {
	nextHandler handler
}

func (ha *handlerA) sendRequest(msg string) string {
	if msg == "A" {
		return fmt.Sprintf("Handler A done with %s", msg)
	} else if ha.nextHandler != nil {
		return ha.nextHandler.sendRequest(msg)
	} else {
		return fmt.Sprintf("Failed to handle request %s", msg)
	}
}

type handlerB struct {
	nextHandler handler
}

func (ha *handlerB) sendRequest(msg string) string {
	if msg == "B" {
		return fmt.Sprintf("Handler B done with %s", msg)
	} else if ha.nextHandler != nil {
		return ha.nextHandler.sendRequest(msg)
	} else {
		return fmt.Sprintf("Failed to handle request %s", msg)
	}
}

type handlerC struct {
	nextHandler handler
}

func (ha *handlerC) sendRequest(msg string) string {
	if msg == "C" {
		return fmt.Sprintf("Handler C done with %s", msg)
	} else if ha.nextHandler != nil {
		return ha.nextHandler.sendRequest(msg)
	} else {
		return fmt.Sprintf("Failed to handle request %s", msg)
	}
}

func main() {
	a := &handlerA{}
	b := &handlerB{}
	c := &handlerC{}
	a.nextHandler = b
	b.nextHandler = c

	req := "B"
	fmt.Println(a.sendRequest(req))

	req = "C"
	fmt.Println(a.sendRequest(req))

	req = "P"
	fmt.Println(a.sendRequest(req))
}
