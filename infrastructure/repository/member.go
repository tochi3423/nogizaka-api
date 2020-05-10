package repository

import (
	"github.com/nogizaka-api/db"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nogizaka-api/infrastructure/model"
	"github.com/nogizaka-api/entity"
	"github.com/nogizaka-api/userInterface/form"
	"github.com/jinzhu/gorm"
)

type (
	IMember interface {
		SelectMemberByName(params form.Params) (*entity.Member, error)
		SelectAllMember(params form.Params) (*entity.Members, error)
	}

	Member struct {
		DB *gorm.DB
	}
)

func NewMember() IMember {
	return &Member{
		DB: db.New(),
	}
}


func (repository *Member) SelectAllMember(params form.Params) (*entity.Members, error) {
	memberarray := model.Members{}
	members := &entity.Members{}
	memberrecords := repository.DB.Table("members")
	pagination := entity.NewPagination(params.Page, params.Perpage)
	memberrecords.Count(pagination.TotalCount)
	err := memberrecords.Offset(pagination.Offset()).Limit(pagination.Limit()).Find(&memberarray).Error
	if err != nil {
		return nil, err
	}
	models := memberarray

	for i := 0; i < len(models); i++ {
		member := &entity.Member{
			ID:            memberarray[i].ID,
			Name:          memberarray[i].Name,
			Birthday:      memberarray[i].Birthday,
			Blood:         memberarray[i].Blood,
			Constellation: memberarray[i].Constellation,
			Height:        memberarray[i].Height,
			Status:        memberarray[i].Status,
			Description:   memberarray[i].Description,
			CreatedAt:     memberarray[i].CreatedAt,
			UpdatedAt:     memberarray[i].UpdatedAt,
		}
		*members = append(*members, *member)
	}
	return members, err
}

func (repository *Member) SelectMemberByName(params form.Params) (*entity.Member, error) {
	membermodel := model.Member{}
	memberrecord := repository.DB.Table("members").Where("name=?", params.Name)
	err := memberrecord.Find(&membermodel).Error
	if err != nil {
		return nil, err
	}

	member := &entity.Member{
		ID:            membermodel.ID,
		Name:          membermodel.Name,
		Birthday:      membermodel.Birthday,
		Blood:         membermodel.Blood,
		Constellation: membermodel.Constellation,
		Height:        membermodel.Height,
		Status:        membermodel.Status,
		Description:   membermodel.Description,
		CreatedAt:     membermodel.CreatedAt,
		UpdatedAt:     membermodel.UpdatedAt,
	}
	return member, err

}