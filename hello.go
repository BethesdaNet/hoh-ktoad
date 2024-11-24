package main

import "fmt"
import "github.com/BethesdaNet/hoh-ktoad/pkg/things"

func main() {

	// Create a new Ktoad
	k := things.Ktoad{
		Name:        "Ktoad",
		Greeting:    "Hello, World!",
		Catchphrase: "I'm sorry Mario, but your princess is in another castle.",
	}

	fmt.Printf("%s says %q\n", k.Name, k.Catchphrase)
}
