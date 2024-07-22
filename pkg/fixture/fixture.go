package fixture

import (
	"errors"

	"github.com/google/uuid"
	"github.com/johnfercher/go-hexagonal/internal/core/consts/userstatus"
	"github.com/johnfercher/go-hexagonal/internal/core/models"
)

func Error() error {
	id, _ := uuid.NewRandom()
	return errors.New(id.String())
}

func User() *models.User {
	id, _ := uuid.NewRandom()
	citizenID, _ := uuid.NewRandom()
	return &models.User{
		ID:        id.String(),
		CitizenID: citizenID.String(),
		Status:    userstatus.Pending,
		Name:      "name",
		Address:   Address(),
		Info: map[string]string{
			"key": "value",
		},
	}
}

func UserCreation() *models.UserCreation {
	id, _ := uuid.NewRandom()
	return &models.UserCreation{
		CitizenID: id.String(),
		Name:      "name",
		Address:   Address(),
	}
}

func Address() *models.Address {
	return &models.Address{
		Street:  "street",
		City:    "city",
		State:   "RJ",
		Country: "Brasil",
	}
}

func Info(qtd int) map[string]string {
	m := make(map[string]string)

	for i := 0; i < qtd; i++ {
		key, _ := uuid.NewRandom()
		value, _ := uuid.NewRandom()
		m[key.String()] = value.String()
	}

	return m
}
