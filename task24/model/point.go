package model

import "math"

type Point struct {
	//Так как название полей в нижнем регистре это делает их видимыми только для для этого пакета
	x float64
	y float64
}

// NewPoint Конструктор для точки
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p1 *Point) GetDistance(p2 *Point) float64 {
	//Евклидова метрика
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}
