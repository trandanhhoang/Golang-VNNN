package utils

import (
	"fmt"
	"math/rand"
)

var IDLENGTH = 10

func RandomId() string {
	var letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	idForOrder := make([]byte, IDLENGTH)
	for i := range idForOrder {
		idForOrder[i] = letter[rand.Intn(len(letter))]
	}
	fmt.Println("idForOrder", string(idForOrder))
	return string(idForOrder)
}
