package models

type User struct {
	ID        int     `mapstructure:"id"`
	Name      string  `mapstructure:"name"`
	Surname   string  `mapstructure:"surname"`
	Email     string  `mapstructure:"email"`
	Age       int     `mapstructure:"age"`
	Height    float64 `mapstructure:"height"`
	Active    bool    `mapstructure:"active"`
	CreatedAt string  `mapstructure:"created_at"`
}
