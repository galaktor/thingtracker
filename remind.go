package main

import (
	"fmt"
	"time"
)

var dailyTicker *time.Ticker
var stopTicker chan bool

func startTimer() {
	dailyTicker := newTicker()
	go func() {
	loop:
		for {
			select {
			case t := <-dailyTicker.C:
				fmt.Printf("ticker expired at %v\n", t)
				remindAll()
				dailyTicker = newTicker()
			case <-stopTicker:
				dailyTicker.Stop()
				dailyTicker = nil
				break loop
			}

		}
	}()
}

func remindAll() {
	for _,t := range things {
		fmt.Printf("reminding: %v\n", t.Id)
		err := remindOne(t)
		if err != nil {
			println(err.Error())
			// TODO HANDLE THIS
		}
	}
}

func remindOne(t *Thing) error {
	if(!t.AllDone()) {
		return t.EmailParticipants()
	}

	return nil
}

func newTicker() *time.Ticker {
	now := time.Now()

	hour := 10
	min  := 30

	next := time.Date(now.Year(), now.Month(), now.Day(), hour, min, 0, 0, time.Local)

	// RAPH HACK, DELETE ME
	next = now.Add(time.Second*15)
	
	if !next.After(now) {
		next = next.Add(time.Hour * 24)
	}
	dur := next.Sub(now)
	return time.NewTicker(dur)
}
