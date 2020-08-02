package dining_philosophers

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
}

func (p Philosopher) eat() {
	for {
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		fmt.Print("eating")

		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()

	}
}

func main() {

	chopsticks := make([]*Chopstick,5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
 	}
 	philosophers := make([]*Philosopher,5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			leftChopstick: chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
		}
	}
}
