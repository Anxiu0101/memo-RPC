package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	pb "memo-RPC/eventserver/ecommerce"
	"memo-RPC/eventserver/model"
	"strconv"
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
	item := model.BuildEventResponse(&data)

	return &pb.ShowEventResponse{
		Item: item,
	}, nil
}

func (eventService *EventService) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	log.Println("Creating Token")
	data := model.BuildEventModel(req.Item)
	if err := model.DB.Create(data).Error; err != nil {
		return &pb.CreateEventResponse{
			Id: "0",
		}, err
	}

	return &pb.CreateEventResponse{
		Id: strconv.Itoa(int(data.ID)),
	}, nil
}

func (eventService *EventService) ListEvents(ctx context.Context, req *pb.ListEventsRequest) (*pb.ListEventsResponse, error) {

	var items []model.Event
	if err := model.DB.Model(model.Event{}).Find(&items).Error; err != nil {
		return &pb.ListEventsResponse{
			Item: nil,
		}, err
	}

	var result []*pb.Event
	for i, item := range items {
		result[i] = model.BuildEventResponse(&item)
	}

	return &pb.ListEventsResponse{
		Item: result,
	}, nil
}

func (eventService *EventService) DeleteEvent(ctx context.Context, req *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {

	var data model.Event
	if err := model.DB.Model(model.Event{}).Where("id = ?", req.Id).First(&data).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return &pb.DeleteEventResponse{
			Id: "0",
		}, err
	}

	if err := model.DB.Model(model.Event{}).Delete(&data).Error; err != nil {
		return &pb.DeleteEventResponse{
			Id: "0",
		}, err
	}

	return &pb.DeleteEventResponse{
		Id: strconv.Itoa(int(data.ID)),
	}, nil
}

func (eventService *EventService) UpdateEvent(ctx context.Context, req *pb.UpdateEventRequest) (*pb.UpdateEventResponse, error) {

	var data = model.BuildEventModel(req.Item)
	if err := model.DB.Model(model.Event{}).Where("ID = ?", req.Item.Id).Updates(data).Error; err != nil {
		return &pb.UpdateEventResponse{
			Item: nil,
		}, err
	}

	return &pb.UpdateEventResponse{
		Item: model.BuildEventResponse(data),
	}, nil
}
