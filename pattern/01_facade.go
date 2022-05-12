package main

import "fmt"

type gpu struct {
}

func (gpu) StartGPU() {
	fmt.Println("GPU started")
}

type cpu struct {
}

func (cpu) StartCPU() {
	fmt.Println("CPU started")
}

type memory struct {
}

func (g *memory) InitMemory() {
	fmt.Println("Memory init")
}

type computerFacade struct {
	Memory	memory
	GPU		gpu
	CPU		cpu
}

func (c computerFacade) startComputer() {
	c.Memory.InitMemory()
	c.CPU.StartCPU()
	c.GPU.StartGPU()
}

func main() {
	pc := computerFacade{}
	pc.startComputer()
}
