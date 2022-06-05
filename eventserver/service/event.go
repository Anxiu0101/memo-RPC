package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	pb "memo-RPC/eventserver/ecommerce"
	"memo-RPC/eventserver/model"
)

type EventService struct {
	pb.UnimplementedEventServiceServer
}

func (eventService *EventService) ShowEvent(ctx context.Context, req *pb.ShowEventRequest) (*pb.ShowEventResponse, error) {

	// 查询事件，若不存在或查询错误，写入日志并返回错误
	var data model.Event
	if err := model.DB.Model(model.Event{}).Where("id = ?", req.Id).First(&data).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return &pb.ShowEventResponse{
			Item: nil,
		}, err
	}

	// 事件模型序列化为响应信息
	result := model.BuildEventResponse(&data)

	return &pb.ShowEventResponse{
		Item: result,
	}, nil
}

func (eventService *EventService) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	data := model.BuildEventModel(req.Item)
	if err := model.DB.Create(data).Error; err != nil {
		return &pb.CreateEventResponse{
			Id: "0",
		}, err
	}

	return &pb.CreateEventResponse{
		Id: string(data.ID),
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
