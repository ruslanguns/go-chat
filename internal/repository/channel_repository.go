package repository

import (
	"github.com/ruslanguns/go-chat/internal/domain"
	"github.com/ruslanguns/go-chat/internal/domain/model"
	"github.com/ruslanguns/go-chat/internal/errors"
	"gorm.io/gorm"
)

type ChannelRepository interface {
	Create(channel *model.Channel) error
	GetByID(id domain.EntityID) (*model.Channel, error)
	GetByName(name string) (*model.Channel, error)
	Update(channel *model.Channel) error
	Delete(id domain.EntityID) error
	List(offset, limit int) ([]*model.Channel, error)
	AddUser(channelID, userID domain.EntityID) error
	RemoveUser(channelID, userID domain.EntityID) error
	GetUsers(channelID domain.EntityID, offset, limit int) ([]*model.User, error)
}

type channelRepository struct {
	db *gorm.DB
}

func NewChannelRepository(db *gorm.DB) ChannelRepository {
	return &channelRepository{db: db}
}

func (r *channelRepository) Create(channel *model.Channel) error {
	err := r.db.Create(channel).Error
	if err != nil {
		if r.db.Error != nil && r.db.Error.Error() == "UNIQUE constraint failed: channels.name" {
			return errors.NewAppError(errors.ErrAlreadyExists, "A channel with this name already exists")
		}
		return errors.NewAppError(errors.ErrInternal, "Failed to create channel")
	}
	return nil
}

func (r *channelRepository) GetByID(id domain.EntityID) (*model.Channel, error) {
	var channel model.Channel
	err := r.db.First(&channel, "id = ?", id.String()).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewAppError(errors.ErrNotFound, "Channel not found")
		}
		return nil, errors.NewAppError(errors.ErrInternal, "Failed to get channel")
	}
	return &channel, nil
}

func (r *channelRepository) GetByName(name string) (*model.Channel, error) {
	var channel model.Channel
	err := r.db.Where("name = ?", name).First(&channel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewAppError(errors.ErrNotFound, "Channel not found")
		}
		return nil, errors.NewAppError(errors.ErrInternal, "Failed to get channel")
	}
	return &channel, nil
}

func (r *channelRepository) Update(channel *model.Channel) error {
	err := r.db.Save(channel).Error
	if err != nil {
		return errors.NewAppError(errors.ErrInternal, "Failed to update channel")
	}
	return nil
}

func (r *channelRepository) Delete(id domain.EntityID) error {
	result := r.db.Delete(&model.Channel{}, "id = ?", id.String())
	if result.Error != nil {
		return errors.NewAppError(errors.ErrInternal, "Failed to delete channel")
	}
	if result.RowsAffected == 0 {
		return errors.NewAppError(errors.ErrNotFound, "Channel not found")
	}
	return nil
}

func (r *channelRepository) List(offset, limit int) ([]*model.Channel, error) {
	var channels []*model.Channel
	err := r.db.Offset(offset).Limit(limit).Find(&channels).Error
	if err != nil {
		return nil, errors.NewAppError(errors.ErrInternal, "Failed to list channels")
	}
	return channels, nil
}

func (r *channelRepository) AddUser(channelID, userID domain.EntityID) error {
	err := r.db.Exec("INSERT INTO user_channels (channel_id, user_id) VALUES (?, ?)", channelID.String(), userID.String()).Error
	if err != nil {
		return errors.NewAppError(errors.ErrInternal, "Failed to add user to channel")
	}
	return nil
}

func (r *channelRepository) RemoveUser(channelID, userID domain.EntityID) error {
	result := r.db.Exec("DELETE FROM user_channels WHERE channel_id = ? AND user_id = ?", channelID.String(), userID.String())
	if result.Error != nil {
		return errors.NewAppError(errors.ErrInternal, "Failed to remove user from channel")
	}
	if result.RowsAffected == 0 {
		return errors.NewAppError(errors.ErrNotFound, "User not found in channel")
	}
	return nil
}

func (r *channelRepository) GetUsers(channelID domain.EntityID, offset, limit int) ([]*model.User, error) {
	var users []*model.User
	err := r.db.Table("users").
		Joins("JOIN user_channels ON users.id = user_channels.user_id").
		Where("user_channels.channel_id = ?", channelID.String()).
		Offset(offset).Limit(limit).
		Find(&users).Error
	if err != nil {
		return nil, errors.NewAppError(errors.ErrInternal, "Failed to get users in channel")
	}
	return users, nil
}
