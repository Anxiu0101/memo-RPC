package service

import "context"
import pb "memo-RPC/server/ecommerce"

type EventService struct {
}

func (eventService *EventService) ShowEvent(ctx context.Context, req *pb.ShowEventRequest) (*pb.ShowEventResponse, error) {

	println("Show Event Info")

	return &pb.ShowEventResponse{
		Item: nil,
	}, nil
}

func (eventService *EventService) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	return &pb.CreateEventResponse{
		Id: "0",
	}, nil
}

func (eventService *EventService) ListEvents(ctx context.Context, req *pb.ListEventsRequest) (*pb.ListEventsResponse, error) {
	return &pb.ListEventsResponse{
		Item: nil,
	}, nil
}

func (eventService *EventService) DeleteEvent(ctx context.Context, req *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {
	return &pb.DeleteEventResponse{
		Id: "0",
	}, nil
}

func (eventService *EventService) UpdateEvent(ctx context.Context, req *pb.UpdateEventRequest) (*pb.UpdateEventResponse, error) {
	return &pb.UpdateEventResponse{
		Item: nil,
	}, nil
}
