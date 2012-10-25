package main

type NameAge struct {
	name string
	age  int
}

func (na *NameAge) getName() string {
	return na.name
}
