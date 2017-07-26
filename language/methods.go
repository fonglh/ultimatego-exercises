// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

// Add imports.
import "fmt"

// Declare a struct that represents a ball player.
// Include field called name, atBats and hits.
type player struct {
	name   string
	atBats int
	hits   int
}

// Declare a method that calculates the batting average for a player.
func (p *player) average() float32 {
	if p.atBats == 0 {
		return 0.0
	}

	return float32(p.hits) / float32(p.atBats)
}

func main() {

	// Create a slice of players and populate each player
	// with field values.
	players := []player{
		{name: "Abel", atBats: 30, hits: 10},
		{name: "Bill", atBats: 20, hits: 19},
		{name: "Cathy", atBats: 6, hits: 5},
	}

	// Display the batting average for each player in the slice.
	for i := range players {
		fmt.Printf("Address: %p, ", &players[i])
		fmt.Printf("%s: AVG[.%.f]\n", players[i].name, players[i].average()*1000)
	}

	fmt.Println()

	// This form is slower as a copy of the data is created at the variable p.
	for _, p := range players {
		fmt.Printf("Address: %p, ", &p)
		fmt.Printf("%s: AVG[.%.f]\n", p.name, p.average()*1000)
	}
}
