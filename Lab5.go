package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	p := NewPerson("Виталий", 24)
	PrintPerson(*p)
	Birthday(p)
	PrintPerson(*p)

	c := Circle{2}
	fmt.Println("Площадь круга ", CalcArea(c))

	s := make([]Shape, 5)
	for i := 0; i < 5; i++ {
		r := rand.Intn(2)
		if r == 1 {
			s[i] = Circle{rand.Float32() * 5}
		} else {
			s[i] = Rectangle{float32(rand.Intn(10)), float32(rand.Intn(10))}
		}
	}

	PrintShape(s)

	b := Book{"1984", "Джордж Оруэлл", 1949}

	fmt.Println(b.toString())

}

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	p := Person{name: name, age: age}
	return &p
}

func Birthday(p *Person) {
	p.age = p.age + 1
}

func PrintPerson(p Person) {
	fmt.Println("Имя: ", p.name, " Возраст: ", p.age)
}

type Circle struct {
	radius float32
}

func CalcArea(c Circle) float32 {
	return float32(math.Pow(float64(c.radius), 2)) * math.Pi
}

type Shape interface {
	Area() float32
}

func (c Circle) Area() float32 {
	return CalcArea(c)
}

type Rectangle struct {
	a float32
	b float32
}

func (r Rectangle) Area() float32 {
	return r.a * r.b
}

func PrintShape(s []Shape) {
	for i := 0; i < len(s); i++ {
		fmt.Println("Площадь объекта ", i, " Равна ", s[i].Area())
	}
}

type Book struct {
	name   string
	author string
	date   int
}

type Stringer interface {
	toString()
}

func (b Book) toString() string {
	return fmt.Sprintf("Название %v, Автор %v, Дата выхода %v", b.name, b.author, b.date)
}
