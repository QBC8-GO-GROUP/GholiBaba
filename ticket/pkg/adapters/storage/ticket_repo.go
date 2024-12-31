package storage

import (
	"context"

	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/internal/ticket/domain"
	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/internal/ticket/port"
	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/pkg/adapters/storage/mapper"
	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/pkg/adapters/storage/types"
	"github.com/QBC8-GO-GROUP/GholiBaba/ticket/pkg/logger"
	"gorm.io/gorm"
)

type ticketRepo struct {
	db *gorm.DB
}

func NewTicketRepo(db *gorm.DB) port.Repo {
	return &ticketRepo{db}

}

func (tr *ticketRepo) Create(ctx context.Context, ticketDomain domain.Ticket) (domain.TicketID, error) {
	ticket := mapper.TicketDomain2Storage(ticketDomain)
	return domain.TicketID(ticket.ID), tr.db.Table("tickets").WithContext(ctx).Create(ticket).Error
}

func (tr *ticketRepo) GetTickets(ctx context.Context, userID domain.UserID) ([]*domain.Ticket, error) {
	var tickets []types.Ticket
	err := tr.db.Model(&types.Ticket{}).Table("tickets").Where("user_id = ?", userID).WithContext(ctx).Find(&tickets).Error

	if err != nil {
		logger.Error(err.Error(), nil)
		return nil, err
	}

	logger.Info("get tickets based on user has been called succesfuly", nil)
	var mappedTickets []*domain.Ticket
	for _, ticket := range tickets {
		mappedTickets = append(mappedTickets, mapper.TicketStorage2Domain(ticket))
	}
	return mappedTickets, nil
}

func (tr *ticketRepo) Delete(ctx context.Context, ticketID domain.TicketID) error {
	return tr.db.Where("id = ?", ticketID).Delete(&types.Ticket{}).Error
}
