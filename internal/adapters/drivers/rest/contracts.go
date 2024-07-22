package rest

type UserCreationRequest struct {
	CitizenID string   `json:"citizen_id"`
	Name      string   `json:"name"`
	Address   *Address `json:"address"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type UserCreationResponse struct {
	ID        string   `json:"id"`
	CitizenID string   `json:"citizen_id"`
	Name      string   `json:"name"`
	Address   *Address `json:"address"`
}
