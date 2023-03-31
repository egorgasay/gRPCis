package handler

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-storage/internal/grpc-storage/storage"
	"grpc-storage/internal/grpc-storage/usecase"
	api "grpc-storage/pkg/api/storage"
)

type Handler struct {
	api.UnimplementedStorageServer
	logic *usecase.UseCase
}

func New(logic *usecase.UseCase) *Handler {
	return &Handler{logic: logic}
}

func (h *Handler) Set(ctx context.Context, r *api.SetRequest) (*api.SetResponse, error) {
	memUsage := h.logic.Set(r.Key, r.Value)
	return &api.SetResponse{
		Status:    "ok",
		Total:     memUsage.Total,
		Available: memUsage.Available,
	}, nil
}

func (h *Handler) Get(ctx context.Context, r *api.GetRequest) (*api.GetResponse, error) {
	value, err := h.logic.Get(r.Key)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return &api.GetResponse{
				Value: err.Error(),
			}, status.Error(codes.NotFound, err.Error())
		}
		return &api.GetResponse{
			Value: err.Error(),
		}, err
	}

	return &api.GetResponse{
		Value: value,
	}, nil
}
