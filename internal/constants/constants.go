package constants

import "fmt"

const (
	// Name is the program name
	Name = "gnm"
	// Usage is for simple description
	Usage   = "Go Networks Manager: manage your ssh/vpn connections"
	Version = "0.0.1"

	// CheckSymbol is the code for check symbol
	CheckSymbol = "\u2714 "
	// CrossSymbol is the code for check symbol
	CrossSymbol = "\u2716 "

	Logo = `
  _______ .__   __. .___  ___. 
 /  _____||  \ |  | |   \/   | 
|  |  __  |   \|  | |  \  /  | 
|  | |_ | |  .    | |  |\/|  | 
|  |__| | |  |\   | |  |  |  | 
 \______| |__| \__| |__|  |__| 
	`
)

func PrintLogo() {
	fmt.Println(Logo)
}
