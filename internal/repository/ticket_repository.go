package repository

import "database/sql"

type TicketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) InsertTicket(
	tenantID int64,
	title string,
	description string,
	priority string,
	customerID int64,
) error {

	_, err := r.db.Exec(`
		INSERT INTO tickets
		(tenant_id, title, description, status, priority, customer_id, created_at)
		VALUES (?, ?, ?, 'open', ?, ?, NOW())
	`,
		tenantID,
		title,
		description,
		priority,
		customerID,
	)

	return err
}
func (r *TicketRepository) GetTicketsByTenant(tenantID int64) ([]map[string]interface{}, error) {
	rows, err := r.db.Query(`
		SELECT id, title, description, status, priority, created_at
		FROM tickets
		WHERE tenant_id = ?
		ORDER BY created_at DESC
	`, tenantID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []map[string]interface{}

	for rows.Next() {
		var id int64
		var title, description, status, priority string
		var createdAt string

		err := rows.Scan(&id, &title, &description, &status, &priority, &createdAt)
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, map[string]interface{}{
			"id":          id,
			"title":       title,
			"description": description,
			"status":      status,
			"priority":    priority,
			"created_at":  createdAt,
		})
	}

	return tickets, nil
}
