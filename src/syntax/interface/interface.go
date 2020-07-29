package main

type geometry interface {
	area() float64
	perim() float64
}

type rect struct{
  width,height float64
}
