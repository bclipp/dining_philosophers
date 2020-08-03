package main

import (
	"fmt"
	"sync"
)

//each chopstick a mutex
//each phil is a go ruitine
//each is associated with 2 chop l/r

type Chopstick struct {sync.Mutex}

type Philosopher struct {
	leftChopstick, rightChopstick *Chopstick
	id  int
}

func (p Philosopher) eat(startChannel chan bool,eatChannel chan int) {
	for i := 0; i < 3; i++ {
		start := false
		for start != false {
			eatChannel <- p.id
			start = <- startChannel
		}
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()
		fmt.Printf("Philosopher #%d is starting to eat.\n", p.id+1)
		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()
		fmt.Printf("Philosopher #%d is finished eating.\n", p.id+1)
		//time.Sleep(2*time.Second)
		eatingWaitGroup.Done()
	}
}

var eatingWaitGroup sync.WaitGroup


func main() {
	counter := 5
	chopsticks := make([]*Chopstick,counter)

	eatChannel := make(chan int)
	var startChannels []chan bool
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
		startChannels[i] = make(chan bool)
 	}
 	philosophers := make([]*Philosopher,5)
	for i := 0; i < counter; i++ {
		philosophers[i] = &Philosopher{
			leftChopstick: chopsticks[i],
			rightChopstick: chopsticks[(i+1)%counter],
			id: i,
		}
		eatingWaitGroup.Add(1)
		go philosophers[i].eat(
			startChannels[i],
			eatChannel,
			)
	}
	eaters := 0

	for range eatChannel {
		if eaters <= 2 {
			eater := <- eatChannel
			eaters++
			startChannels[eater] <- true
		}

	}

	eatingWaitGroup.Wait()
}
