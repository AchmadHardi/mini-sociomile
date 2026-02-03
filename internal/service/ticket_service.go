package service

import "mini-sociomile/internal/repository"

type TicketService struct {
	repo *repository.TicketRepository
}

func NewTicketService(repo *repository.TicketRepository) *TicketService {
	return &TicketService{repo: repo}
}
func (s *TicketService) ListTickets(tenantID int64) ([]map[string]interface{}, error) {
	return s.repo.GetTicketsByTenant(tenantID)
}

func (s *TicketService) CreateTicket(
	tenantID int64,
	title string,
	description string,
	priority string,
	customerID int64,
) error {

	return s.repo.InsertTicket(
		tenantID,
		title,
		description,
		priority,
		customerID,
	)
}
