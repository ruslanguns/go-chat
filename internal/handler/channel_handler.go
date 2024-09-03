package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ruslanguns/go-chat/internal/domain"
	"github.com/ruslanguns/go-chat/internal/domain/model"
	"github.com/ruslanguns/go-chat/internal/errors"
	"github.com/ruslanguns/go-chat/internal/service"
)

type ChannelHandler struct {
	channelService service.ChannelService
}

func NewChannelHandler(channelService service.ChannelService) *ChannelHandler {
	return &ChannelHandler{
		channelService: channelService,
	}
}

func (h *ChannelHandler) Create(w http.ResponseWriter, r *http.Request) {
	var channel model.Channel
	if err := json.NewDecoder(r.Body).Decode(&channel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdChannel, err := h.channelService.CreateChannel(channel.Name, channel.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdChannel)
}

func (h *ChannelHandler) Get(w http.ResponseWriter, r *http.Request) {
	channelID, err := domain.ParseEntityID(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid channel ID", http.StatusBadRequest)
		return
	}

	channel, err := h.channelService.GetChannelByID(channelID)
	if err != nil {
		if err, ok := err.(errors.AppError); ok && err.ErrorType() == errors.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(channel)
}

func (h *ChannelHandler) Update(w http.ResponseWriter, r *http.Request) {
	channelID, err := domain.ParseEntityID(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid channel ID", http.StatusBadRequest)
		return
	}

	var channel model.Channel
	if err := json.NewDecoder(r.Body).Decode(&channel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	channel.ID = channelID

	if err := h.channelService.UpdateChannel(&channel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(channel)
}

func (h *ChannelHandler) Delete(w http.ResponseWriter, r *http.Request) {
	channelID, err := domain.ParseEntityID(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid channel ID", http.StatusBadRequest)
		return
	}

	if err := h.channelService.DeleteChannel(channelID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ChannelHandler) List(w http.ResponseWriter, r *http.Request) {
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if limit == 0 {
		limit = 10 // default limit
	}

	channels, err := h.channelService.ListChannels(offset, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(channels)
}

func (h *ChannelHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	channelID, err := domain.ParseEntityID(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid channel ID", http.StatusBadRequest)
		return
	}

	var userID domain.EntityID
	if err := json.NewDecoder(r.Body).Decode(&userID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.channelService.AddUserToChannel(channelID, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ChannelHandler) RemoveUser(w http.ResponseWriter, r *http.Request) {
	channelID, err := domain.ParseEntityID(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid channel ID", http.StatusBadRequest)
		return
	}

	userID, err := domain.ParseEntityID(chi.URLParam(r, "userId"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.channelService.RemoveUserFromChannel(channelID, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ChannelHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	channelID, err := domain.ParseEntityID(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid channel ID", http.StatusBadRequest)
		return
	}

	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if limit == 0 {
		limit = 10 // default limit
	}

	users, err := h.channelService.GetChannelUsers(channelID, offset, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}
