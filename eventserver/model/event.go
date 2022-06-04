package model

import (
	"gorm.io/gorm"
	pb "memo-RPC/eventserver/ecommerce"
)

type Event struct {
	gorm.Model

	Name      string `json:"name"`
	Content   string `json:"content"`
	EndTimeAt int64  `json:"endTime_at"`
	State     int    `json:"state,omitempty"`
	Type      int    `json:"type,omitempty"`
}

func BuildEventResponse(event *Event) *pb.Event {
	return &pb.Event{
		Id:        uint32(event.ID),
		Name:      event.Name,
		Content:   event.Content,
		CreateAt:  event.CreatedAt.Unix(),
		UpdateAt:  event.UpdatedAt.Unix(),
		EndTimeAt: event.EndTimeAt,
		Type:      int32(event.Type),
		State:     int32(event.State),
	}
}

func BuildEventModel(event *pb.Event) *Event {
	return &Event{
		Name:      event.Name,
		Content:   event.Content,
		EndTimeAt: event.EndTimeAt,
		Type:      int(event.Type),
		State:     int(event.State),
	}
}
