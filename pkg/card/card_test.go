package card_test

import (
	"fmt"
	"github.com/ArtemBond13/hw1.3/pkg/card"
)

func ExampleTranslateMCC() {
	fmt.Println(card.TranslateMCC("4012"))
	fmt.Println(card.TranslateMCC("4013"))
	fmt.Println(card.TranslateMCC("4156"))
	fmt.Println(card.TranslateMCC("3456"))
	// Output:
	// Категория не указана
	// Men's and Boys' Clothing
	// Women's Accessory and Special Stores
	// Категория не указана
}
