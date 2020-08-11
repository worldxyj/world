package models

type Menu struct {
	Id        uint
	Name      string
	Css       string
	Url       string
	Pid       uint
	Sort      uint16
	Status    uint8
	CreatedAt uint64
	UpdatedAt uint64
}
