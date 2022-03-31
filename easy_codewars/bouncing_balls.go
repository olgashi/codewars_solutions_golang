package main

import (
	"fmt"
)

func BouncingBall(h, bounce, window float64) int {
	if h <= 0 || bounce <= 0 || bounce >= 1 || window >= h {
		return -1
	}

	numberOfBouncingBalls := 1
	bounceHeight := h

	for bounceHeight > window {
		if bounceHeight = bounceHeight * bounce; bounceHeight > window {
			numberOfBouncingBalls += 2
		}
	}

  return numberOfBouncingBalls
}

func main() {
	fmt.Println(BouncingBall(3, 0.66, 1.5)) //3
	fmt.Println(BouncingBall(3, 1, 1.5)) // -1
	fmt.Println(BouncingBall(-1, 0.5, 1.7)) // -1
	fmt.Println(BouncingBall(2, 1.5, 0.3)) // -1
	fmt.Println(BouncingBall(1, 0.5, 5)) // -1
	fmt.Println(BouncingBall(2, 0.5, 0.7)) // 3
	fmt.Println(BouncingBall(5, 0.3, 0.3)) // 5
}
