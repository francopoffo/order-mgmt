package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "github.com/francopoffo/common/api"
	"github.com/francopoffo/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service     OrderService
	amqpChannel *amqp.Channel
}

func NewGrpcHandler(grpcServer *grpc.Server, service OrderService, ch *amqp.Channel) {
	handler := &grpcHandler{service: service, amqpChannel: ch}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) ProcessOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	q, err := h.amqpChannel.QueueDeclare(broker.OrderCreatedEvent, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	tr := otel.Tracer("amqp")

	amqpContext, messageSpan := tr.Start(ctx, fmt.Sprintf("AMQP - publish - %s", q.Name))

	defer messageSpan.End()

	o := &pb.OrderResponse{}

	marshalledOrder, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	// inject the headers
	headers := broker.InjectAMQPHeaders(amqpContext)

	h.amqpChannel.PublishWithContext(amqpContext, "", q.Name, false, false, amqp.Publishing{
		ContentType:  "application/json",
		Body:         marshalledOrder,
		DeliveryMode: amqp.Persistent,
		Headers:      headers,
	})

	return o, nil
}
