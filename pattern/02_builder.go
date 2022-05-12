package main

import "fmt"

type iHouse interface {
	Open() string
}

type iHouseBuilder interface {
	setWindowType(string) iHouseBuilder
	setDoorType(string) iHouseBuilder
	setNumFloor(int) iHouseBuilder
	Build() iHouse
}

type houseBuilder struct {
	windowType string
	doorType   string
	numFloor   int
}

func (hb *houseBuilder) setWindowType(windowType string) iHouseBuilder {
	hb.windowType = windowType
	return hb
}

func (hb *houseBuilder) setDoorType(doorType string) iHouseBuilder {
	hb.doorType = doorType
	return hb
}

func (hb *houseBuilder) setNumFloor(numFloor int) iHouseBuilder {
	hb.numFloor = numFloor
	return hb
}

func (hb *houseBuilder) Build() iHouse {
	return &house{
		windowType: hb.windowType,
		doorType:   hb.doorType,
		floor:      hb.numFloor,
	}
}

type house struct {
	windowType string
	doorType   string
	floor      int
}

func (h *house) Open() string {
	return fmt.Sprintf("House with \nwindow type: %s; \ndoor type :%s; \nand %d floors \nis ready to move in.",
		h.windowType, h.doorType, h.floor)
}
func new() iHouseBuilder {
	return &houseBuilder{}
}

func main() {
	builder := new()
	house := builder.setWindowType("regular").setDoorType("sliding").setNumFloor(3).Build()

	fmt.Println(house.Open())
}
