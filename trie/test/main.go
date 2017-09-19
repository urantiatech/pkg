package main

import (
	"fmt"

	"github.com/urantiatech/pkg/trie"
)

func main() {

	t := trie.New()

	t.Insert("", "Null Key")
	t.Insert(" ", "Space")
	t.Insert("  ", "Multiple Spaces")
	t.Insert("Spirit", "This is spirit")
	t.Insert("Spirit of Truth", "This is Spirit of Truth")
	t.Insert("Spirit of Wisdom", "This is Spirit of Wisdom")
	t.Insert("Spirit of Wisdom is great", "This is Spirit of Wisdom is great")
	t.Insert("Spiritual", "Spiritual values")

	fmt.Println("Values inserted successfully")

	t.Traverse("\t", "........")

	keys := []string{"", " ", "  ", "Spirit", "Spirit of Truth", "Spirit of", "Spirit of Wisdom is Great"}
	for _, k := range keys {
		v, err := t.Value(k)
		if err == nil {
			fmt.Printf("[%s] : [%s]\n", k, v)
		} else {
			fmt.Printf("[%s] : %s\n", k, err.Error())
		}
	}
}
