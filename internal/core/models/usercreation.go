package models

type UserCreation struct {
	CitizenID string
	Name      string
	Address   *Address
}
