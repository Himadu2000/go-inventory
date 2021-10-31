package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to Inventory programme! Developed by Himadu.")
	type Inventory struct {
		name string
		qty  int
	}
	inventory := []Inventory{{name: "Item1", qty: 0}}

	getInputValues := func() Inventory {
		fmt.Println("What is the name of the item?")
		name := ""
		fmt.Scanln(&name)
		fmt.Println("What is the qty of the item?")
		qty := 0
		fmt.Scanln(&qty)
		item := Inventory{name, qty}
		return item
	}

	hrLine := func() {
		fmt.Println(strings.Repeat("-", 50))
	}
	listInventory := func() {
		fmt.Println("Index Name Qty")
		for i, j := range inventory {
			fmt.Println(i, j.name, j.qty)
		}
	}

	for {
		hrLine()
		fmt.Println(`
		Press, 
		\n1 to Reload inventory. 
		\n2 to Add item to inventory. 
		\n3 to Update item. 
		\n4 to Remove item. 
		\n0 to Quit the programme.`)
		hrLine()
		listInventory()
		hrLine()
		input := 9
		fmt.Scanln(&input)
		switch input {
		case 0:
			fmt.Println("Thankyou for using the Inventory programme! Developed by Himadu.")
			os.Exit(0)
		case 1:
			listInventory()
		case 2:
			inventory = append(inventory, getInputValues())
		case 3:
			fmt.Println("What index do you want to update?")
			index := 0
			inventory[index] = getInputValues()
			fmt.Println("Index", index, " Updated!")
		case 4:
			fmt.Println("What index do you want to remove?")
			index := 0
			fmt.Scanln(&index)
			fmt.Println("Are you sure you want to remove", index, "? Enter y")
			verify := "n"
			fmt.Scanln(&verify)
			if verify == "y" {
				inventory = append(inventory[:index], inventory[index+1:]...)
				fmt.Println("Removed index ", index)
			}
		default:
			main()
		}
	}
}
