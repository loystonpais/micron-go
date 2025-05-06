package main

type Display interface {
	Repr() string
}

type DisplayCode interface {
	Code() string
}
