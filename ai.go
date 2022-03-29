package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var firstInput int
var secondInput int

// Creates a neuron struct
type neuron struct {
	weights [4]float64
	bias    float64
	output  float64
}

// Creates a 2D array of neurons
var neurons = [][]neuron{
	{
		neuron{[4]float64{}, 0, 0}, neuron{[4]float64{}, 0, 0},
	},
	{
		neuron{[4]float64{}, 0, 0}, neuron{[4]float64{}, 0, 0}, neuron{[4]float64{}, 0, 0}, neuron{[4]float64{}, 0, 0},
	},
	{
		neuron{[4]float64{}, 0, 0},
	},
}

// Trains the AI
func train(x int, w int, y float64, z float64) {
	firstInput = x
	secondInput = w
	for {
		for i := 0; i < firstInput; i++ {
			for m := 0; m < secondInput; m++ {
				for j := 0; j < len(neurons); j++ {
					for k := 0; k < len(neurons[j]); k++ {
						neurons[j][k].output = 0
						for l := 0; l < len(neurons[j][k].weights); l++ {
							neurons[j][k].output += neurons[j][k].weights[l] * float64(i)
							neurons[j][k].output += neurons[j][k].weights[l] * float64(m)
							// Debug: fmt.Printf("Training: Neuron row %d number %d : %f \n", j, k, neurons[j][k].output)
						}
						neurons[j][k].output += neurons[j][k].bias
					}
				}
			}
		}
		if int(math.Floor(neurons[len(neurons)-1][0].output)) == int(y) || int(math.Floor(neurons[len(neurons)-1][0].output)) == int(z) {
			fmt.Printf("Correct Guess: %v\n", neurons[len(neurons)-1][0].output)
			break
		} else {
			for i := 0; i < len(neurons); i++ {
				for j := 0; j < len(neurons[i]); j++ {
					for k := 0; k < len(neurons[i][j].weights); k++ {
						neurons[i][j].weights[k] = rand.Float64()
					}
				}
			}
		}
	}
}

// Generates training data
func genTrain(x int, y int) {
	for i := 0; i < x; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		x := rand.Intn(y)
		if x == 0 {
			break
		}
		first := x
		rand.Seed(time.Now().UTC().UnixNano())
		x = rand.Intn(y)
		if x == 0 {
			break
		}
		last := x
		max := float64(first + last)
		min := max - 1
		fmt.Printf("Dataset generated: First: %v, Last: %v, Min: %v, Max: %v\n", first, last, min, max)
		train(first, last, max, min)
	}
}
func main() {
	// Seed random number generator
	rand.Seed(time.Now().UTC().UnixNano())

	// Initialize weights
	for i := 0; i < len(neurons); i++ {
		for j := 0; j < len(neurons[i]); j++ {
			for k := 0; k < len(neurons[i][j].weights); k++ {
				neurons[i][j].weights[k] = rand.Float64()
			}
		}
	}

	// Initialize bias
	for i := 0; i < len(neurons); i++ {
		for j := 0; j < len(neurons[i]); j++ {
			for k := 0; k < len(neurons[i][j].weights); k++ {
				neurons[i][j].weights[k] = rand.Float64()
			}
			neurons[i][j].bias = rand.Float64()
		}
	}

	// Train
	genTrain(34, 436)

	// Initialize AI
	firstInput = 2
	secondInput = 5
	for i := 0; i < firstInput; i++ {
		for m := 0; m < secondInput; m++ {
			for j := 0; j < len(neurons); j++ {
				for k := 0; k < len(neurons[j]); k++ {
					neurons[j][k].output = 0
					for l := 0; l < len(neurons[j][k].weights); l++ {
						neurons[j][k].output += neurons[j][k].weights[l] * float64(i)
						neurons[j][k].output += neurons[j][k].weights[l] * float64(m)
						// Debug: fmt.Printf("Final test: Neuron row %d number %d : %f \n", j, k, neurons[j][k].output)
					}
					neurons[j][k].output += neurons[j][k].bias
				}
			}
		}
	}
	fmt.Printf("Final test: %v\n", neurons[len(neurons)-1][0].output)
}
