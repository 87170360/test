package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%s:%d", p.Name, p.Age)
}

type PersonSorter struct {
	by      func(p1, p2 *Person) bool
	persons []*Person
}

func (ps PersonSorter) Len() int {
	return len(ps.persons)
}

func (ps PersonSorter) Swap(i, j int) {
	ps.persons[i], ps.persons[j] = ps.persons[j], ps.persons[i]
}

func (ps PersonSorter) Less(i, j int) bool {
	return ps.by(ps.persons[i], ps.persons[j])
}

type By func(p1, p2 *Person) bool

func (by By) Sort(persons []*Person) {
	ps := PersonSorter{
		by:      by,
		persons: persons,
	}
	sort.Sort(ps)
}

func main() {
	s := []*Person{
		&Person{"Job", 60},
		&Person{"Jack", 90},
		&Person{"Bill", 80},
	}

	ByName := func(p1, p2 *Person) bool { return p1.Name < p2.Name }
	ByAge := func(p1, p2 *Person) bool { return p1.Age < p2.Age }

	By(ByName).Sort(s)
	fmt.Println(s)
	By(ByAge).Sort(s)
	fmt.Println(s)
}
