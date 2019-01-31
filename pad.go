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
}

func (nopad *NoPad) AddA() {
	nopad.a++
	nopad.a++
	nopad.a++
	nopad.a++
	nopad.a++
	nopad.a++
	nopad.a++
	nopad.a++
	nopad.a++
	nopad.a++
}

func (nopad *NoPad) AddB() {
	nopad.b++
	nopad.b++
	nopad.b++
	nopad.b++
	nopad.b++
	nopad.b++
	nopad.b++
	nopad.b++
	nopad.b++
	nopad.b++
}

type Pad struct {
	a   uint64
	_p1 [8]uint64
	_p2 [8]uint64
	b   uint64
}

func (pad *Pad) AddA() {
	pad.a++
	pad.a++
	pad.a++
	pad.a++
	pad.a++
	pad.a++
	pad.a++
	pad.a++
	pad.a++
	pad.a++
}

func (pad *Pad) AddB() {
	pad.b++
	pad.b++
	pad.b++
	pad.b++
	pad.b++
	pad.b++
	pad.b++
	pad.b++
	pad.b++
	pad.b++
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
