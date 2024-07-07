package httphandlers

import (
	"encoding/json"
	"net/http"

	"github.com/johnfercher/go-hexagonal/internal/core/models"
	"github.com/johnfercher/go-hexagonal/internal/core/ports"
)

type UserRegisterHandler struct {
	template     *HTTPTemplate
	userRegister ports.UserRegister
}

func NewUserRegisterHandler(userRegister ports.UserRegister) *UserRegisterHandler {
	return &UserRegisterHandler{
		template:     NewHTTPTemplate(),
		userRegister: userRegister,
	}
}

func (u *UserRegisterHandler) Method() string {
	return http.MethodPost
}

func (u *UserRegisterHandler) Pattern() string {
	return "/users"
}

func (u *UserRegisterHandler) Func(w http.ResponseWriter, r *http.Request) {
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
