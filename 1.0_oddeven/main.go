package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Perceptron struct {
	weights []float64
	bias    float64
}

func NewPerceptron(inputSize int) *Perceptron {
	rand.Seed(time.Now().UnixNano())
	weights := make([]float64, inputSize)
	for i := range weights {
		weights[i] = rand.Float64() - 0.5
	}
	return &Perceptron{weights: weights, bias: rand.Float64() - 0.5}
}

func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func (p *Perceptron) Predict(inputs []float64) float64 {
	sum := p.bias
	for i := 0; i < len(inputs); i++ {
		sum += inputs[i] * p.weights[i]
	}
	return sigmoid(sum)
}

func (p *Perceptron) Train(inputs [][]float64, labels []float64, epochs int, lr float64) {
	for epoch := 0; epoch < epochs; epoch++ {
		for i := 0; i < len(p.weights); i++ {
			prediction := p.Predict(inputs[i])
			error := labels[i] - prediction

			for j := 0; j < len(p.weights); j++ {
				p.weights[j] += lr * error * inputs[i][j]
			}
			p.bias += lr * error
		}
	}
}

func main() {
	inputs := [][]float64{
		{0, 0, 0, 0}, // 0 -> четное
		{0, 0, 0, 1}, // 1 -> нечетное
		{0, 0, 1, 0}, // 2 -> четное
		{0, 0, 1, 1}, // 3 -> нечетное
		{0, 1, 0, 0}, // 4 -> четное
		{0, 1, 0, 1}, // 5 -> нечетное
		{0, 1, 1, 0}, // 6 -> четное
		{0, 1, 1, 1}, // 7 -> нечетное
		{1, 0, 0, 0}, // 8 -> четное
		{1, 0, 0, 1}, // 9 -> нечетное
	}
	labels := []float64{0, 1, 0, 1, 0, 1, 0, 1, 0, 1}

	p := NewPerceptron(4)
	p.Train(inputs, labels, 1000, 0.1)

	for i := 0; i < len(inputs); i++ {
		result := p.Predict(inputs[i])
		pred := 0
		if result > 0.5 {
			pred = 1
		}
		fmt.Printf("Число %d: предсказание = %d (%.2f), ожидаемое = %.0f\n", i, pred, result, labels[i])
	}
}
