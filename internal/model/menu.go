package model

type TypeMenu string

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      TypeMenu
}

