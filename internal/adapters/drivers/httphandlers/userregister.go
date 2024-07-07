package httphandlers

import (
	"encoding/json"
	"github.com/johnfercher/go-hexagonal/internal/core/models"
	"github.com/johnfercher/go-hexagonal/internal/core/ports"
	"net/http"
)

type UserRegister struct {
	template     *HttpTemplate
	userRegister ports.UserRegister
}

func NewUserRegister(userRegister ports.UserRegister) *UserRegister {
	return &UserRegister{
		template:     NewHttpTemplate(),
		userRegister: userRegister,
	}
}

func (u *UserRegister) Method() string {
	return http.MethodPost
}

func (u *UserRegister) Pattern() string {
	return "/users"
}

func (u *UserRegister) Func(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userCreation := &UserCreationRequest{}
	err := json.NewDecoder(r.Body).Decode(userCreation)
	if err != nil {
		u.template.WriteError(w, ErrUnmarshallRequest.WithInner(err))
		return
	}

	id, err := u.userRegister.Register(ctx, &models.UserCreation{
		CitizenID: userCreation.CitizenID,
		Name:      userCreation.Name,
		Address: &models.Address{
			Street:  userCreation.Address.Street,
			City:    userCreation.Address.City,
			State:   userCreation.Address.State,
			Country: userCreation.Address.Country,
		},
	})
	if err != nil {
		u.template.WriteError(w, mapErr[err])
		return
	}

	response := &UserCreationResponse{
		ID:        id,
		CitizenID: userCreation.CitizenID,
		Name:      userCreation.Name,
		Address:   userCreation.Address,
	}

	u.template.Write(w, response, http.StatusCreated)
}
