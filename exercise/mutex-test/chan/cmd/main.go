package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var adderChan = 0
var adderMutex = 0
var adderSpin = 0
var adderAtomic int32 = 0
var c = make(chan struct{}, 1)
var m = sync.Mutex{}

var exit = make(chan struct{}, 1)
const times = 1000000

type spinLock uint32
var l spinLock = 0

func Init() {
	adderChan = 0
	adderMutex = 0
	adderSpin = 0
	adderAtomic = 0
}


func (sl *spinLock)Lock() {
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		runtime.Gosched()
	}
}

func (sl *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}


func main() {
	for i:= 0 ; i < 1000; i++ {
		//init
		Init()

		fmt.Println("-------------------------", i + 1)
		TChan()
		TMutex()
		TSpinLock()
		TAtomic()
		fmt.Println("chan adder:", adderChan)
		fmt.Println("mux adder:", adderMutex)
		fmt.Println("spin adder:", adderSpin)
		fmt.Println("atomic adder:", adderAtomic)
	}
}

func TChan() {
	defer func(tm time.Time) {
		fmt.Println("chan:", time.Since(tm))
	}(time.Now())

	for i:=0; i < times;i++ {
		go addChan()
	}
	select {
	case <-exit:
		return
	}
}

func TMutex() {
	defer func(tm time.Time) {
		fmt.Println("mutex:", time.Since(tm))
	}(time.Now())

	for i:=0; i < times;i++ {
		go addMutex()
	}
	select {
	case <-exit:
		return
	}
}

func TSpinLock() {
	defer func(tm time.Time) {
		fmt.Println("spin:", time.Since(tm))
	}(time.Now())

	for i:=0; i < times;i++ {
		go addSpin()
	}
	select {
	case <-exit:
		return
	}
}

func TAtomic() {
	defer func(tm time.Time) {
		fmt.Println("atomic:", time.Since(tm))
	}(time.Now())

	for i:=0; i < times;i++ {
		go func() {
			atomic.AddInt32(&adderAtomic, 1)
			if adderAtomic == times {
				exit <- struct{}{}
			}
		}()
	}
	select {
	case <-exit:
		return
	}
}


func addChan() {
	c <- struct{}{}
	adderChan ++
	if adderChan == times {
		exit <- struct{}{}
	}
	<-c
}

func addMutex() {
	m.Lock()
	adderMutex ++
	if adderMutex == times {
		exit <- struct{}{}
	}
	m.Unlock()
}

func addSpin() {
	l.Lock()
	adderSpin ++
	if adderSpin == times {
		exit <- struct{}{}
	}
	l.Unlock()
}