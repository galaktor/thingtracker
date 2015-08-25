package main

import (
	"time"
	"fmt"
)

// thing id -> ticker
var tickers map[int]*time.Ticker

func reset() {
	if tickers != nil {
		for _,t := range tickers {
			t.Stop()
		}
	}
	tickers = make(map[int]*time.Ticker)
}

func refreshReminders(things map[int]*Thing) error {
	reset()
	
	for _,t := range things {
		setOne(t)
	}

	return nil
}

func setOne(t *Thing) error {
	// cancel existing ticker
	ticker, ok := tickers[t.Id]
	if ok {
		ticker.Stop()
		delete(tickers, t.Id)
	}

	// hard-code time of day for now
	expHour := 17
	expMin  := 15
	
	now := time.Now()

	then := time.Date(now.Year(), now.Month(), now.Day(), expHour, expMin, 0, 0, time.Local)
	
	// if past time, move by a day to tomorrow
	if now.After(then) {
		then = then.AddDate(0, 0, 1)
		fmt.Printf("changed then to %v\n", then)
	}

	fmt.Printf("id %v then at %v\n", t.Id, then)
	dur := then.Sub(now)
	fmt.Printf("id %v wait time is %v\n", t.Id, dur)
	
	ticker = time.NewTicker(dur)
	tickers[t.Id] = ticker

	waitFor(ticker)
	
	fmt.Printf("set thing id %v to expire at %v\n", t.Id, then)

	return nil
}

func waitFor(t *time.Ticker) {
	go func(c <-chan time.Time) {
		fmt.Printf("expired: %v\n", <-c)
	}(t.C)
}


// todo: on expire, set anew for next day
