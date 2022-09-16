package main

import (
	"fmt"
)

func main() {
	var oSergey = Human{}
	// var mSergey = map[string]any{
	// 	"Name":   "Sergey Gonchar",
	// 	"Age":    uint(31),
	// 	"Height": centimeter(174.3),
	// 	"Weight": kilogram(71.5),
	// }
	var mSergey = map[string]any{
		"Name":   "Sergey Gonchar",
		"Age":    31,
		"Height": 174.3,
		"Weight": 71.5,
	}

	if err := assign(&oSergey, mSergey); err != nil {
		fmt.Println(fmt.Errorf("assign error: %w", err))
	}

	fmt.Println()
	fmt.Println("Map:", mSergey)
	fmt.Println("Obj:", oSergey)
}
