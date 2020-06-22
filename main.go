package main

import "github.com/atotto/clipboard"
import "password-generator/packages"

func main() {
	password := packages.RandomString(16)
	clipboard.WriteAll(password)
}