package models

type RaceVariant struct {
	ID          uint
	Name        string
	Description string

	RaceID uint
	Race   *Race
}
