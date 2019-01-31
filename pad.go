package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type NoPad struct {
	a uint64
	b uint64
	c uint64
}

func (myatomic *NoPad) AddA() {
	myatomic.a++
	myatomic.a++
	myatomic.a++
	myatomic.a++
	myatomic.a++
	myatomic.a++
	myatomic.a++
	myatomic.a++
	myatomic.a++
	myatomic.a++
}

func (myatomic *NoPad) AddB() {
	myatomic.b++
	myatomic.b++
	myatomic.b++
	myatomic.b++
	myatomic.b++
	myatomic.b++
	myatomic.b++
	myatomic.b++
	myatomic.b++
	myatomic.b++
}

type Pad struct {
	a   uint64
	_p1 [8]uint64
	_p2 [8]uint64
	b   uint64
}

func (myatomic *Pad) AddA() {
	myatomic.a++
	myatomic.a++
	myatomic.a++
	myatomic.a++
	myatomic.a++
	myatomic.a++
	myatomic.a++
	myatomic.a++
	myatomic.a++
	myatomic.a++
}

func (myatomic *Pad) AddB() {
	myatomic.b++
	myatomic.b++
	myatomic.b++
	myatomic.b++
	myatomic.b++
	myatomic.b++
	myatomic.b++
	myatomic.b++
	myatomic.b++
	myatomic.b++
}

func main() {
	t := time.Now()
	runtime.GOMAXPROCS(2)
	//var pad = &Pad{}
	var pad = &NoPad{}

	go func() {
		for {
			pad.AddA()
		}
	}()

	go func() {
		for {
			pad.AddB()
		}
	}()

	time.Sleep(time.Minute)
	fmt.Println(uint64(time.Now().Sub(t)), pad.a+pad.b)
	os.Exit(0)
}
