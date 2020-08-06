package models

type Admin struct {
	Id        int
	Name      string
	Password  string
	Tel       string
	Ip        string
	LoginAt   int64
	CreatedAt int64
}
