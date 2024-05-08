package main

import (
	"fmt"
	"goBasic/imports"
)

func main() {
	newTicket := imports.Ticket{
		ID:    123,
		Event: "FEM course",
	}
	newTicket.PrintEvent()
	fmt.Println(newTicket)
}
