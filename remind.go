package main

import (
	"time"
	"fmt"
)

type Reminder struct {
	due time.Time
	dur time.Duration
	t *time.Ticker
	done chan bool
}

func NewReminder(due time.Time) *Reminder {
	return &Reminder{due: due, done: make(chan bool)}
}

func (r *Reminder) Stop() {
	if r.t != nil {
		r.done<-true
	}
}

// thing id -> reminder
var reminders map[int]*Reminder


func reset() {
	if reminders != nil {
		for _,r := range reminders {
			r.Stop()
		}
	}
	reminders = make(map[int]*Reminder)
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
	reminder, ok := reminders[t.Id]
	if ok {
		reminder.Stop()
		delete(reminders, t.Id)
	}



	reminder = NewReminder(t.Due)
	reminders[t.Id] = reminder

	reminder.Set()
	reminder.Start()
	
	return nil
}


func (r *Reminder) Set() {
	
	now := time.Now()
	// due date at 10:30
	then := r.due.Add(time.Hour*10).Add(time.Minute*30)

	// TEMP HACK FOR DEV PURPOSES ONLY
//	then = time.Now().Add(time.Second*10)
	
	// if past time, move by a day to tomorrow
	// INEFFICIENT AS HECK BUT WHATEVS!
	for now.After(then) {
		then = then.AddDate(0, 0, 1)
	}

	fmt.Printf("then: %v\n", then)
	r.dur = then.Sub(now)
}

func (r *Reminder) Start() {
	if r.t == nil {
		go func() {

		Loop:
			for {
				r.t = time.NewTicker(r.dur)
				fmt.Printf("wait time is %v\n", r.dur)
				
				select {
				case <- r.t.C:  // expired
					println("expired")
					r.Set()
					continue
				case <- r.done: // closed
					println("closed")
					r.t.Stop()
					break Loop
				}				
			}
			
			r.t = nil
			println("gofunc dead")
		}()
	}
}

// todo: on expire, set anew for next day
