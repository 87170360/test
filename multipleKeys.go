/*
优化了less
*/

package main

import (
	"fmt"
	"sort"
)

// A Change is a record of source code changes, recording user, language, and delta size.
type Change struct {
	user     string
	language string
	lines    int
}

//type lessFunc func(p1, p2 *Change) bool
type lessFunc func(p1, p2 *Change) int

// multiSorter implements the Sort interface, sorting the changes within.
type multiSorter struct {
	changes []Change
	less    []lessFunc
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *multiSorter) Sort(changes []Change) {
	ms.changes = changes
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.changes)
}

// Swap is part of sort.Interface.
func (ms *multiSorter) Swap(i, j int) {
	ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that is either Less or
// !Less. Note that it can call the less functions twice per call. We
// could change the functions to return -1, 0, 1 and reduce the
// number of calls for greater efficiency: an exercise for the reader.
/*
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.changes[i], &ms.changes[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}
*/

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.changes[i], &ms.changes[j]
	for k := 0; k < len(ms.less); k++ {
		less := ms.less[k]
		ret := less(p, q)
		if ret > 0 {
			return true
		} else if ret < 0 {
			return false
		}
	}
	return false
}

var changes = []Change{
	{"gri", "Go", 100},
	{"ken", "C", 150},
	{"glenda", "Go", 200},
	{"rsc", "Go", 200},
	{"r", "Go", 100},
	{"ken", "Go", 200},
	{"dmr", "C", 100},
	{"r", "C", 150},
	{"gri", "Smalltalk", 80},
}

// ExampleMultiKeys demonstrates a technique for sorting a struct type using different
// sets of multiple fields in the comparison. We chain together "Less" functions, each of
// which compares a single field.
func main() {
	// Closures that order the Change structure.
	user := func(c1, c2 *Change) int {
		switch {
		case c1.user < c2.user:
			return 1
		case c1.user > c2.user:
			return -1
		case c1.user == c2.user:
			return 0
		}
		return 0
	}
	language := func(c1, c2 *Change) int {
		switch {
		case c1.language < c2.language:
			return 1
		case c1.language > c2.language:
			return -1
		case c1.language < c2.language:
			return 0
		}
		return 0
	}
	increasingLines := func(c1, c2 *Change) int {
		switch {
		case c1.lines < c2.lines:
			return 1
		case c1.lines > c2.lines:
			return -1
		case c1.lines < c2.lines:
			return 0
		}
		return 0
	}
	decreasingLines := func(c1, c2 *Change) int {
		switch {
		case c1.lines > c2.lines:
			return 1
		case c1.lines < c2.lines:
			return -1
		case c1.lines == c2.lines:
			return 0
		}
		return 0
	}
	/*
		user := func(c1, c2 *Change) bool {
			return c1.user < c2.user
		}
		language := func(c1, c2 *Change) bool {
			return c1.language < c2.language
		}
		increasingLines := func(c1, c2 *Change) bool {
			return c1.lines < c2.lines
		}
		decreasingLines := func(c1, c2 *Change) bool {
			return c1.lines > c2.lines // Note: > orders downwards.
		}
	*/

	// Simple use: Sort by user.
	OrderedBy(user).Sort(changes)
	fmt.Println("By user:", changes)

	// More examples.
	OrderedBy(user, increasingLines).Sort(changes)
	fmt.Println("By user,<lines:", changes)

	OrderedBy(user, decreasingLines).Sort(changes)
	fmt.Println("By user,>lines:", changes)

	OrderedBy(language, increasingLines).Sort(changes)
	fmt.Println("By language,<lines:", changes)

	OrderedBy(language, increasingLines, user).Sort(changes)
	fmt.Println("By language,<lines,user:", changes)

}
