package model

type User struct {
	ID       int64
	Email    string
	Password string
	Role     string // admin, agent, customer
	TenantID int64
}
