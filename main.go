package main

import "fmt"

func isDividable(cand int, number int) bool {
	if cand % number == 0 {
		return true
	} else {
		return false
	}
}

func evaluate(cand int, i int, stream chan bool) {
	stream <- isDividable(cand, i)
}

func worker(cand int, howmuch int, stream chan bool) {
	for i := 2; i <= howmuch + 1; i++ {
		go evaluate(cand, i, stream)
	}
}

func main() {
	oldPrime := 2
	newPrime := 3
	count := 2
	sum := newPrime - oldPrime
	oldPrime = newPrime
	a := 4
	for a < 100 {
		dividableStream := make(chan bool)
		howmuch := a/2
		worker(a, howmuch, dividableStream)
		it := 0
		for {
			value := <-dividableStream
			if(value) {
				close(dividableStream)
				break;
			}
			it = it + 1
			if it == howmuch {
				newPrime = a
				sum = sum + newPrime - oldPrime
				count = count + 1
				oldPrime = newPrime
				fmt.Println(newPrime)
				close(dividableStream)
				break;
			}
		}
		a = a + 1
	}
	fmt.Println(count)
	fmt.Println(sum)
}
