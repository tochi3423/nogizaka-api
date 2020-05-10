package response

import (
	"github.com/nogizaka-api/entity"
	"time"
)

type (
	Member struct {
		ID            int64
		Name          string
		Birthday      string
		Blood         string
		Constellation string
		Height        string
		Status        string
		Description   string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	Members []Member
)
func NewMember(entity *entity.Member) *Member {
	if entity == nil {
		return nil
	} else {
		return &Member{
			ID:            entity.ID,
			Name:          entity.Name,
			Birthday:      entity.Birthday,
			Blood:         entity.Blood,
			Constellation: entity.Constellation,
			Height:        entity.Height,
			Status:        entity.Status,
			Description:   entity.Description,
			CreatedAt:     entity.CreatedAt,
			UpdatedAt:     entity.UpdatedAt,
		}
	}
}

func NewMembers(entities *entity.Members) *Members {
	members := Members{}
	for _, v := range *entities {
		members = append(members, *NewMember(&v))
	}
	return &members
}
