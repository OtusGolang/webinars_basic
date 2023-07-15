package main

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
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func calculateArea(s any) (float64, error) {
	sTyped, ok := s.(Shape)
	if !ok {
		return 0, fmt.Errorf("Ошибка: переданный объект не является фигурой.")
	}
	return sTyped.Area(), nil
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 10, Height: 5}
	t := Triangle{Base: 8, Height: 6}

	area1, err1 := calculateArea(c)
	if err1 == nil {
		fmt.Printf("Круг: радиус %.2f\n", c.Radius)
		fmt.Printf("Площадь: %.2f\n", area1)
	} else {
		fmt.Println(err1)
	}

	area2, err2 := calculateArea(r)
	if err2 == nil {
		fmt.Printf("Прямоугольник: ширина %.2f, высота %.2f\n", r.Width, r.Height)
		fmt.Printf("Площадь: %.2f\n", area2)
	} else {
		fmt.Println(err2)
	}

	area3, err3 := calculateArea(t)
	if err3 == nil {
		fmt.Printf("Треугольник: основание %.2f, высота %.2f\n", t.Base, t.Height)
		fmt.Printf("Площадь: %.2f\n", area3)
	} else {
		fmt.Println(err3)
	}

	s := "Не фигура"
	_, err4 := calculateArea(s)
	if err4 != nil {
		fmt.Println(err4)
	}
}
