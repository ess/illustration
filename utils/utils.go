package utils

import (
	"fmt"
)

func PackageMethod() {
	fmt.Println("I am a package method. I was called either as 'PackageMethod()' within the package or 'utils.PackageMethod()' outside the package.")
}
