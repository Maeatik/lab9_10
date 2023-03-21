package main

import (
	"fmt"
	"math"
)

//23

var (
	cashes            float64 = 3
	minutesOnOne      float64 = 3
	numOfBuyersPerMin float64 = 7
	limit             float64 = 5
)

func RecursiveFactorial(number int) float64 {
	if number >= 1 {
		return float64(number) * RecursiveFactorial(number-1)
	} else {
		return 1
	}
}

func main() {
	intCashes := int(cashes)
	ro := float64(numOfBuyersPerMin / minutesOnOne)
	var probability float64 = 0
	for i := 0; i <= intCashes; i++ {
		probability = probability + math.Pow(ro, float64(i))/(RecursiveFactorial(i))
	}

	probability = probability +
		((math.Pow(ro, cashes+1)) / ((cashes - ro) * RecursiveFactorial(intCashes)) *
			(1 - (math.Pow(ro/cashes, limit))))

	probability = math.Pow(probability, -1)
	fmt.Println("Вероятность того, что обслуживанием не занят не один:", probability, "\n")

	probabilityOfFailure := math.Pow(ro, cashes+limit) / (math.Pow(cashes, limit) * RecursiveFactorial(intCashes)) * probability
	fmt.Println("Вероятность отказа:", probabilityOfFailure)

	numberOfApplications := (math.Pow(ro, cashes+1) / (cashes * RecursiveFactorial(intCashes))) *
		(((1 - math.Pow(ro/cashes, limit)) * (limit + 1 - limit*ro/limit)) / math.Pow((1-ro/cashes), 2)) *
		probability

	fmt.Println("Среднее число заявок, находящихся в очереди:", numberOfApplications, "\n")

	time := numberOfApplications / numOfBuyersPerMin

	fmt.Println("Среднее время ожидания в очереди:", time, "\n")

	capacity := 1 - probabilityOfFailure
	fmt.Println("Относительная пропускная способность системы:", capacity, "\n")

	absoluteCapacity := numOfBuyersPerMin * capacity
	fmt.Println("Абсолютная пропускная способность системы:", absoluteCapacity, "\n")

	busyChan := absoluteCapacity / minutesOnOne
	fmt.Println("Среднее число занятых касс", busyChan, "\n")

	freeChan := cashes - busyChan
	fmt.Println("Среднее число простаивающих касс касс", freeChan, "\n")

	if numberOfApplications < limit {
		fmt.Println("Процесс обслуживания эффективен")
	} else {
		fmt.Println("Процесс обслуживания не эффективен")
	}
}
