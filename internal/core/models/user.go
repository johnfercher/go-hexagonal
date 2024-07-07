package models

import "github.com/johnfercher/go-hexagonal/internal/core/consts/userstatus"

type User struct {
	ID        string
	CitizenID string
	Status    userstatus.Status
	Name      string
	Address   *Address
	Info      map[string]string
}
