package MeContract

type MeResponseBody struct {
	Id        string   `json:"id"`
	Perms     []string `json:"roles"`
	Email     string   `json:"email"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
}
