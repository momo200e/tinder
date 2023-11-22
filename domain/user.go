package domain

type Gender int

const (
	Male   Gender = 1
	Female Gender = 2
)

type User struct {
	Name        string
	Height      uint8
	Gender      Gender
	RemainDates uint8
}
