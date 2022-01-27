package main

import "fmt"

func main(){
	interfaces := []IF{}
	h := new(human)
	h.firstName = "lu"
	h.lastName = "shijie"
	interfaces = append(interfaces, h)

	c := new(car)
	c.model = "jepan"
	c.factory = "test"
	interfaces = append(interfaces, c)

	for _,f := range interfaces{
		fmt.Printf("name is %s \n", f.getName())
	}
}

type IF interface {
	getName() string
}

type human struct {
	firstName, lastName string
}

func (h *human) getName() string {
	return h.firstName + "_" + h.lastName
}

type car struct {
	factory, model string
}

func (c *car) getName() string {
	return c.factory + "_" + c.model
}