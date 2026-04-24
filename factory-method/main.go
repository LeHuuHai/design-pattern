package main

import "fmt"

type Transport interface {
	Delivery()
}

type Truck struct{}

func (_ Truck) Delivery() {
	fmt.Println("truck delivering ")
}

type Ship struct{}

func (_ Ship) Delivery() {
	fmt.Println("ship delivering ")
}

type Logistic interface {
	CreateTransport() Transport
}

type BaseLogistic struct{}

func (_ BaseLogistic) PlanDelivery(l Logistic) {
	t := l.CreateTransport()
	t.Delivery()
}

type Trucklogistic struct {
	BaseLogistic
}

func (_ Trucklogistic) CreateTransport() Transport {
	return Truck{}
}

type Shiplogistic struct {
	BaseLogistic
}

func (_ Shiplogistic) CreateTransport() Transport {
	return Ship{}
}

func main() {
	t := Trucklogistic{}
	t.PlanDelivery(t)
	s := Shiplogistic{}
	s.PlanDelivery(s)
}
