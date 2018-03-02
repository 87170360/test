package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name   string
	Age    int
	Gender int
}

func (p *Person) String() string {
	return fmt.Sprintf("name:%s,age:%d,gender:%d", p.Name, p.Age, p.Gender)
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

type By2 func(p1, p2 *Person) int

type PersonSorter2 struct {
	persons []*Person
	by      []By2
}

func (p *PersonSorter2) Sort(persons []*Person) {
	p.persons = persons
	sort.Sort(p)
}

func (p *PersonSorter2) Len() int {
	return len(p.persons)
}

func (p *PersonSorter2) Swap(i, j int) {
	p.persons[i], p.persons[j] = p.persons[j], p.persons[i]
}

func (p *PersonSorter2) Less(i, j int) bool {
	for _, v := range p.by {
		ret := v(p.persons[i], p.persons[j])
		if ret == 1 {
			return true
		} else if ret == -1 {
			return false
		}
	}
	return false
}

func OrderBy(by ...By2) *PersonSorter2 {
	return &PersonSorter2{
		by: by,
	}
}

func main() {
	s := []*Person{
		&Person{"Job", 62, 0},
		&Person{"Job", 60, 0},
		&Person{"Jack", 90, 1},
		&Person{"Bill", 80, 2},
		&Person{"Micheal", 81, 2},
	}

	order1 := func(p1, p2 *Person) int {
		if p1.Name < p2.Name {
			return 1
		} else if p1.Name == p2.Name {
			return 0
		} else {
			return -1
		}
	}
	order2 := func(p1, p2 *Person) int {
		if p1.Age < p2.Age {
			return 1
		} else if p1.Age == p2.Age {
			return 0
		} else {
			return -1
		}
	}
	order3 := func(p1, p2 *Person) int {
		if p1.Gender < p2.Gender {
			return 1
		} else if p1.Gender == p2.Gender {
			return 0
		} else {
			return -1
		}
	}

	OrderBy(order1).Sort(s)
	fmt.Println(s)
	OrderBy(order1, order2).Sort(s)
	fmt.Println(s)
	OrderBy(order3, order2).Sort(s)
	fmt.Println(s)

	/*
		ByName := func(p1, p2 *Person) bool { return p1.Name < p2.Name }
		ByAge := func(p1, p2 *Person) bool { return p1.Age < p2.Age }

		fmt.Println(s)
		By(ByName).Sort(s)
		fmt.Println(s)
		By(ByAge).Sort(s)
		fmt.Println(s)
	*/
}
