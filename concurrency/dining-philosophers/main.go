package main

type Philosopher struct {
	leftHand  bool
	rightHand bool
	status    int
	name      string
}

const (
	EATING = iota
	THINKING
)

func mkPhilosopher(name string) Philosopher {
	return Philosopher{
		name:      name,
		leftHand:  true,
		rightHand: true,
		status:    THINKING,
	}
}

func main() {
	philosophers = []Philosopher{
		mkPhilosopher("Kant"),
		mkPhilosopher("Turing"),
		mkPhilosopher("Kierkegaard"),
		mkPhilosopher("Descartes"),
		mkPhilosopher("Wittgenstein"),
	}

	forkDown, forkUp := make(chan bool), make(chan bool)

	evaluate = func() {
		for {
			select {
			case <-forkUp:
				// Philosophers think.
			case <-forkDown:
				// Next philosopher eats in a round robin fashion.
			}
		}
	}
}
