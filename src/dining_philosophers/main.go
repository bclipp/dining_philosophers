package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick struct {sync.Mutex}

type Philosopher struct {
	leftChopstick, rightChopstick *Chopstick
	id  int
}
// eatChannel is used to request the ability to eat, the requester sends the ID
// startChannel is used to let the eater know they can eat now.
func (p Philosopher) eat(startChannel chan bool,eatChannel chan int, doneChannel chan bool) {
	defer eatingWaitGroup.Done()
	start := false
	// loop is setup to force the eater to wait till its the eater's turn.
	for !start {
		eatChannel <- p.id
		start = <- startChannel
	}
	// each eater eats 3 times, but it would be nice to have each time they eat they ask for permission.
	for i := 0; i < 3; i++ {
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()
		fmt.Printf("Philosopher #%d is starting to eat.\n", p.id)
		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()
		fmt.Printf("Philosopher #%d is finished eating.\n", p.id)
		time.Sleep(2*time.Second)
	}
	doneChannel <- true
}

var eatingWaitGroup sync.WaitGroup

func main() {
	counter := 5
	chopsticks := make([]*Chopstick,counter)
	// doneChannel is used to tell the host the eater is done
	doneChannel := make(chan bool)
	// eatChannel is used to request the ability to eat, the requester sends the ID
	eatChannel := make(chan int)
	// startChannel is used to let the eater know they can eat now.
	var startChannels [5]chan bool
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
			doneChannel,
			)
	}
	eaters := 0
	for  i := 0; i < 5; {

		if eaters < 3 {
			eater := <-eatChannel
			i++
			startChannels[eater] <- true
		}
		done := <- doneChannel
		if done {
			eaters--
		}
	}

	eatingWaitGroup.Wait()

}
