package main

import (
	"sync"
)

func syncWithWaitGroups() {
	var wg sync.WaitGroup
	var varA string
	var varB string

	wg.Add(2)

	go func() {
		defer wg.Done()
		varA = "a"
	}()
	go func() {
		defer wg.Done()
		varB = "b"
	}()

	wg.Wait()

	_, _ = varA, varB
}

func syncWithChannels() {
	var varACh chan string = make(chan string)
	var varBCh chan string = make(chan string)

	go func() {
		varACh <- "a"
	}()
	go func() {
		varBCh <- "b"
	}()

	_, _ = <-varACh, <-varBCh
}
