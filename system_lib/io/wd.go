package io

import (
	"io"
	"sync"
	"time"
)

const (
	writeInterval = 300
	readInterval = 1000
)

type Generator interface {
	GetOne() byte
}

type NumberStacker struct {
	i int
}

func (t *NumberStacker) GetOne() byte{
	t.i++
	return byte(t.i)
}

type WallDoor struct {
	sync.RWMutex
	buffer []byte
	r int
	w int
	ch chan byte
	closer chan struct{}
}

func NewWallDoor() *WallDoor{
	wd := &WallDoor{
		buffer: make([]byte, 100),
		ch: make(chan byte),
		closer: make(chan struct{}),
	}
	g := &NumberStacker{}
	wd.Run(g)
	return wd
}

func (t *WallDoor) generate(g Generator) {
	b := g.GetOne()
	//full
	if (t.w+1)%len(t.buffer) == t.r{
		return
	}
	t.buffer[t.w] = b
	t.w++
	if t.w >= len(t.buffer) {
		t.w = 0
	}
}

func (t *WallDoor) filling() {
	//empty
	if t.w == t.r{
		return
	}
	t.ch <- t.buffer[t.r]
	t.r++
	if t.r >= len(t.buffer) {
		t.r = 0
	}
}

func (t *WallDoor) Pop() (byte, error){
	select {
	case b := <-t.ch:
		return b, nil
	case <-t.closer:
		return 0, io.EOF
	}
}

func (t *WallDoor) Run(g Generator) {
	go func() {
		for {
			select {
			case <-t.closer:
				return
			default:
				t.generate(g)
				time.Sleep(writeInterval * time.Millisecond)
			}
		}
	}()
	go func() {
		for {
			select {
			case <-t.closer:
				return
			default:
				t.filling()
			}
		}
	}()
}

func (t *WallDoor) Close() {
	close(t.closer)
}

type WDStream struct {
	src *WallDoor
}

func NewWDStream(wd *WallDoor) *WDStream{
	return &WDStream{
		src: wd,
	}
}

func (t *WDStream) Read(b []byte) (n int, err error){
	for i:=0;i< len(b);i++ {
		item, err := t.src.Pop()
		if err != nil {
			return i, err
		}
		b[i] = item
	}
	return len(b), nil
}
