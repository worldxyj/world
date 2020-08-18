package models

type Admin struct {
	Id        uint
	RoleId    uint
	Name      string
	Password  string
	Tel       string
	Ip        string
	LoginAt   uint
	CreatedAt uint
}
