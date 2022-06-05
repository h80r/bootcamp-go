package internal

type User struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Surname   string  `json:"surname"`
	Email     string  `json:"email"`
	Age       int     `json:"age"`
	Height    float64 `json:"height"`
	Active    bool    `json:"active"`
	CreatedAt string  `json:"created_at"`
}
