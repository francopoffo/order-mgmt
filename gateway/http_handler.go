package main

import (
	"errors"
	"net/http"

	"github.com/francopoffo/common"
	pb "github.com/francopoffo/common/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client: client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerId}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")

	var items []*pb.ItemWithQuantity

	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateItems(items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	o, err := h.client.ProcessOrder(r.Context(), &pb.CreateOrderRequest{CustomerId: customerID, Items: items})

	rStatus := status.Convert(err)

	if rStatus != nil {
		if rStatus.Code() == codes.InvalidArgument {
			common.WriteError(w, http.StatusBadRequest, rStatus.Message())
			return
		}
	}

	common.WriteJSON(w, http.StatusOK, o)
}

func validateItems(items []*pb.ItemWithQuantity) error {
	if len(items) == 0 {
		return common.ErrNoItems
	}

	for _, item := range items {
		if item.Quantity <= 0 {
			return errors.New("quantity cannot be zero or lower")
		}

		if item.ID == "" {
			return errors.New("item id cannot be empty")
		}
	}

	return nil
}
