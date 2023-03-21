package main

import (
	"fmt"
	"math"
)

const (
	weeklyDemand      = 300
	weeklyOrderAmount = 300
	orderCost         = 20
	storageDailyCost  = 0.03
	timeToDelivery    = 0
)

var (
	storageWeeklyCost = storageDailyCost * 7
)

func getWeeklyCost(orderAmount int64) float64 {
	return orderCost*float64((weeklyDemand/orderAmount)) + storageWeeklyCost*float64((orderAmount/2))
}

func getOptimalOrderAmount() int64 {
	return int64(math.Sqrt(2*orderCost*weeklyDemand) / storageWeeklyCost)
}

func getOrderFrequency(orderAmount int64) float64 {
	return float64(orderAmount / weeklyDemand)
}

func getReserveForOrder() int64 {
	return (weeklyDemand / 7) * timeToDelivery
}

func main() {
	//a
	weeklyCostByCurrStrategy := getWeeklyCost(weeklyOrderAmount)
	fmt.Println("a)  Недельные затраты по существующей стратегии:", weeklyCostByCurrStrategy)

	//b
	optimalOrderAmount := getOptimalOrderAmount()
	weeklyCostByOptimalStrategy := getWeeklyCost(optimalOrderAmount)

	orderFrequency := getOrderFrequency(optimalOrderAmount)

	reserveForOrder := getReserveForOrder()

	fmt.Printf("b) Оптимальный заказ\n"+
		"Объём заказа: %d\n"+
		"Недельные затраты по оптимальной стратегии: %.2f\n "+
		"Новый заказ через: %.2f недели (или %.2f дней)\n"+
		"Заказ подавать при уровне запаса %d фунтов\n",
		optimalOrderAmount, weeklyCostByOptimalStrategy, orderFrequency, (7 * orderFrequency), reserveForOrder)

	//c
	diffOfStrategies := weeklyCostByCurrStrategy - weeklyCostByOptimalStrategy
	fmt.Printf("c) Разница затрат: %.2f", diffOfStrategies)

}
