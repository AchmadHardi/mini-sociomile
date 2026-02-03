package model

import "time"

type Ticket struct {
	ID              int64
	Title           string
	Description     string
	Status          string
	Priority        string
	AssignedAgentID int64
	CustomerID      int64
	TenantID        int64
	CreatedAt       time.Time
}
