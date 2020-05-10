package usecase

import (
  "github.com/nogizaka-api/entity"
  "github.com/nogizaka-api/userInterface/form"
  "github.com/nogizaka-api/infrastructure/repository"
)


type (
  IMember interface {
    CallSelectMemberByName(params form.Params) (*entity.Member, error)
    CallSelectAllMember(params form.Params) (*entity.Members, error)
  }

  Member struct {
    MemberRepo repository.IMember
  }
)

func NewMember() IMember {
  return &Member{
    MemberRepo: repository.NewMember(),
  }
}


func (usecase *Member) CallSelectMemberByName(params form.Params) (*entity.Member, error) {
  entities, err := usecase.MemberRepo.SelectMemberByName(params)
  if err != nil {
    return nil, err
  }

  return entities, nil
}


func (usecase *Member) CallSelectAllMember(params form.Params) (*entity.Members, error) {
  entities, err := usecase.MemberRepo.SelectAllMember(params)
  if err != nil {
    return nil, err
  }

  return entities, nil
}