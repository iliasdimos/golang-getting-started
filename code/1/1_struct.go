package main

import "time"

type Employee struct {
	ID       int
	Name     string
	Address  string
	DoB      time.Time
	Position string
	Salary   int
}

func (e *Employee) ChangeAddress(addr string) {
	e.Address = addr
}

func main() {

}
