package main

import (
	"flag"
	"fmt"
	"os"
)

func displayIntro() {
	fmt.Println("")
	fmt.Println(" d888888b d888888b d888888b db      d88888b")
	fmt.Println(" `~~88~~'   `88'   `~~88~~' 88      88'")
	fmt.Println("    88       88       88    88      88ooooo")
	fmt.Println("    88       88       88    88      88~~~~~")
	fmt.Println("    88      .88.      88    88booo. 88.")
	fmt.Println("    YP    Y888888P    YP    Y88888P Y88888P")
	fmt.Println("")
	fmt.Println("")
}

func displayHelp() {
	fmt.Println("    _   _  ____  __    ____    ")
	fmt.Println("   ( )_( )( ___)(  )  (  _ \\   ")
	fmt.Println("    ) _ (  )__)  )(__  )___/   ")
	fmt.Println("   (_) (_)(____)(____)(__)     ")
	fmt.Println("  ===========================  ")
	fmt.Println("")
	flag.PrintDefaults()
	fmt.Println("")
	fmt.Println("")
}

func init() {
	// DISPLAYS INTRO MESSAGE
	displayIntro()
}

func main() {
	// SET FLAGS ACCEPTED
	help := flag.Bool("h", false, "View help message")

	// PARSE FLAGS
	flag.Parse()

	// DISPLAY HELP IF CALLED
	if *help {
		displayHelp()
	}

	// EXIT WITH SUCCESS
	os.Exit(0)
}