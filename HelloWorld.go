package main

import (
	"fmt"
)

func main() {
	//	helloWorld()
	//	capacity()
	report()

}

func report() {

	plantCapacities := []float64{30, 30, 30, 60, 60, 100}
	activePlants := []int{0, 1}
	gridLoad := 75.

	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Println("Please choose an option: ")

	// anonymous function definition (not inside a class)
	menuLogic := func() {
		var option string
		fmt.Scanln(&option)
		//	println(option)
		switch option {
		case "1":
			generatePlantCapacityReport(plantCapacities...) // ... spread operator (for variadic arguments
		case "2":
			{
				generatePowerGridReport(activePlants, plantCapacities, gridLoad)
			}
		default:
			{
				println("unknown option")
			}

		}
	}
	menuLogic()  // call anonymous function

}

func generatePowerGridReport(activePlants []int, plantCapacities []float64, gridLoad float64) {
	capacity := 0.
	for _, plantId := range activePlants {
		capacity += plantCapacities[plantId]
	}
	fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
	fmt.Printf("%-20s%.0f\n", "Load:", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization:", gridLoad/capacity*100)
}

func generatePlantCapacityReport(plantCapacities ...float64) {
	for idx, cap := range plantCapacities {
		fmt.Printf("Plant %d capacity: %0.f\n", idx, cap)
	}
}

func capacity() {
	var plantCapacities []float64
	plantCapacities = []float64{30, 30, 30, 60, 60, 100}

	var capacity = plantCapacities[0] + plantCapacities[1] + plantCapacities[2] + plantCapacities[3] + plantCapacities[4] + plantCapacities[5]
	var gridLoad = 75.

	utilization := gridLoad / capacity
	fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
	fmt.Printf("%-20s%.0f\n", "Load:", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization:", utilization*100)
}

func helloWorld() {
	fmt.Println("hello world")
	var myInt int
	myInt = 42
	println(myInt)
	var myFloat float32 = 42
	println(myFloat)

}
