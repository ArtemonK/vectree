package main

import "fmt"

type Vehicle interface {
	move()
	broke()
}

func drive(vehicle Vehicle) {
	vehicle.move()
	vehicle.broke()
}

type Car struct{}
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}
func (d Car) broke() {
	fmt.Println("Автомобиль сломался")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}
func (b Aircraft) broke() {
	fmt.Println("Самолет сломался")
}

func main() {

	tesla := Car{}
	boing := Aircraft{}
	drive(tesla.broke)

}
