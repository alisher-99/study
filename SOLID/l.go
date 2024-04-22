package SOLID

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func PrintArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	PrintArea(Rectangle{Height: 10, Width: 10}) // Area: 100
	PrintArea(Circle{Radius: 10})               // Area: 314.1592653589793
}

//////////////////////////////////////////////////////////////////////

// Animal представляет базовый тип животного
type Animal interface {
	Speak() string
}

// Dog представляет собаку, являющуюся подтипом Animal
type Dog struct{}

// Speak возвращает звук, издаваемый собакой
func (d Dog) Speak() string {
	return "Гав-гав!"
}

// Cat представляет кошку, являющуюся подтипом Animal
type Cat struct{}

// Speak возвращает звук, издаваемый кошкой
func (c Cat) Speak() string {
	return "Мяу!"
}

// Zoo представляет зоопарк с животными
type Zoo struct {
	Animals []Animal
}

// MakeNoise играет звук всех животных в зоопарке
func (z Zoo) MakeNoise() {
	for _, animal := range z.Animals {
		fmt.Println(animal.Speak())
	}
}

func main() {
	dog := Dog{}
	cat := Cat{}

	zoo := Zoo{
		Animals: []Animal{dog, cat},
	}

	zoo.MakeNoise()
}

//////////////////////////////////////////////////////////////////////

// Vehicle определяет интерфейс для всех транспортных средств
type Vehicle interface {
	CalculateCost() float64
}

// Truck представляет грузовик
type Truck struct {
	Weight int
}

// CalculateCost вычисляет стоимость перевозки груза грузовиком
func (t Truck) CalculateCost() float64 {
	// Предположим, стоимость зависит от веса груза и расстояния
	return float64(t.Weight) * 0.1 // Пример: стоимость 0.1 за каждый кг груза
}

// Ship представляет корабль
type Ship struct {
	Containers int
}

// CalculateCost вычисляет стоимость перевозки груза кораблем
func (s Ship) CalculateCost() float64 {
	// Предположим, стоимость зависит от количества контейнеров и расстояния
	return float64(s.Containers) * 100 // Пример: стоимость 100 за каждый контейнер
}

// TotalCost вычисляет общую стоимость перевозки груза различными транспортными средствами
func TotalCost(vehicles []Vehicle) float64 {
	totalCost := 0.0
	for _, vehicle := range vehicles {
		totalCost += vehicle.CalculateCost()
	}
	return totalCost
}

func main() {
	truck := Truck{Weight: 5000}
	ship := Ship{Containers: 10}

	vehicles := []Vehicle{truck, ship}

	fmt.Printf("Общая стоимость перевозки груза: $%.2f\n", TotalCost(vehicles))
}
