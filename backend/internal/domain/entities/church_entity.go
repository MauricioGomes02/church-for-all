package entities

// Church representa uma igreja
type Church struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Address string   `json:"address"`
	Phones  []string `json:"phones"`
	Email   string   `json:"email"`
}
