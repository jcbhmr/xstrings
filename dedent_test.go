package xstrings_test

import (
	"fmt"
	"testing"

	"github.com/jcbhmr/xstrings"
)

func ExampleDedent() {
	// https://www.poetryfoundation.org/poems/48860/the-raven
	fmt.Printf(xstrings.Dedent(`
		%s
		Over many a quaint and curious volume of forgotten lore—
		    While I nodded, nearly napping, suddenly there came a tapping,
		%s
		“’Tis some visitor,” I muttered, “tapping at my chamber door—
		            Only this and nothing more.”
	`), "Once upon a midnight dreary, while I pondered, weak and weary,", "As of some one gently rapping, rapping at my chamber door.")
	// Output:
	// Once upon a midnight dreary, while I pondered, weak and weary,
	// Over many a quaint and curious volume of forgotten lore—
	//     While I nodded, nearly napping, suddenly there came a tapping,
	// As of some one gently rapping, rapping at my chamber door.
	// “’Tis some visitor,” I muttered, “tapping at my chamber door—
	//             Only this and nothing more.”
}

func TestDedent(t *testing.T) {
	table := []struct{ Input, Want string }{
		{``, ""},
		{`a`, "a"},
		{`
			One
			Two Two
			Three
		`, "One\nTwo Two\nThree"},
		{`

			Two Two
			Three

			Five

		`, "\nTwo Two\nThree\n\nFive\n"},
	}
	for _, row := range table {
		got := xstrings.Dedent(row.Input)
		if got != row.Want {
			t.Errorf("got:\n%q\nwant:\n%q", got, row.Want)
		}
	}
}
