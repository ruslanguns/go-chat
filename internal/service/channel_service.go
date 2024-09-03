package service

import (
	"github.com/ruslanguns/go-chat/internal/domain"
	"github.com/ruslanguns/go-chat/internal/domain/model"
	"github.com/ruslanguns/go-chat/internal/repository"
)

type ChannelService interface {
	CreateChannel(name, description string) (*model.Channel, error)
	GetChannelByID(id domain.EntityID) (*model.Channel, error)
	GetChannelByName(name string) (*model.Channel, error)
	UpdateChannel(channel *model.Channel) error
	DeleteChannel(id domain.EntityID) error
	ListChannels(offset, limit int) ([]*model.Channel, error)
	AddUserToChannel(channelID, userID domain.EntityID) error
	RemoveUserFromChannel(channelID, userID domain.EntityID) error
	GetChannelUsers(channelID domain.EntityID, offset, limit int) ([]*model.User, error)
}

type channelService struct {
	channelRepo repository.ChannelRepository
	userRepo    repository.UserRepository
}

func NewChannelService(channelRepo repository.ChannelRepository, userRepo repository.UserRepository) ChannelService {
	return &channelService{
		channelRepo: channelRepo,
		userRepo:    userRepo,
	}
}

func (s *channelService) CreateChannel(name, description string) (*model.Channel, error) {
	channel := &model.Channel{
		Name:        name,
		Description: description,
	}

	err := s.channelRepo.Create(channel)
	if err != nil {
		return nil, err
	}

	return channel, nil
}

func (s *channelService) GetChannelByID(id domain.EntityID) (*model.Channel, error) {
	return s.channelRepo.GetByID(id)
}

func (s *channelService) GetChannelByName(name string) (*model.Channel, error) {
	return s.channelRepo.GetByName(name)
}

func (s *channelService) UpdateChannel(channel *model.Channel) error {
	return s.channelRepo.Update(channel)
}

func (s *channelService) DeleteChannel(id domain.EntityID) error {
	return s.channelRepo.Delete(id)
}

func (s *channelService) ListChannels(offset, limit int) ([]*model.Channel, error) {
	return s.channelRepo.List(offset, limit)
}

func (s *channelService) AddUserToChannel(channelID, userID domain.EntityID) error {
	_, err := s.channelRepo.GetByID(channelID)
	if err != nil {
		return err
	}

	_, err = s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}

	return s.channelRepo.AddUser(channelID, userID)
}

func (s *channelService) RemoveUserFromChannel(channelID, userID domain.EntityID) error {
	return s.channelRepo.RemoveUser(channelID, userID)
}

func (s *channelService) GetChannelUsers(channelID domain.EntityID, offset, limit int) ([]*model.User, error) {
	return s.channelRepo.GetUsers(channelID, offset, limit)
}
