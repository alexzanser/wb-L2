package main

import "fmt"

type GPU struct {
}

func (g *GPU) StartGPU() {
	fmt.Println("GPU started")
}

type CPU struct {
}

func (g *CPU) StartCPU() {
	fmt.Println("CPU started")
}

type Memory struct {
}

func (g *Memory) InitMemory() {
	fmt.Println("Memory init")
}

type ComputerFacade struct {
	Memory Memory
	GPU	GPU
	CPU	CPU
}

func (c ComputerFacade) StartComputer() {
	c.Memory.InitMemory()
	c.CPU.StartCPU()
	c.GPU.StartGPU()
}

func main() {
	pc := ComputerFacade{}
	pc.StartComputer()
}
