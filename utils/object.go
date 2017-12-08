package utils

import (
	"fmt"
)

type Object struct{}

func (o *Object) Method() {
	fmt.Println("I am a method on an object. I am about to call a package method.")
	PackageMethod()
}
