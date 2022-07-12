package service

import (
	"context"
	"log"
	"net/http"
	"sut-notification-go/domain/notification/model"
	notifpb "sut-notification-go/pb/notification"
)

func (s *Service) SubscribeNotificationByUserId(ctx context.Context, req *notifpb.SubscribeNotificationRequest) (*notifpb.SubscribeNotificationResponse, error) {
	newNotif := model.Notification{
		UserId:   req.UserId,
		Rejected: 0,
		Accepted: 0,
	}

	if res := s.H.DB.Create(&newNotif); res.Error != nil {
		log.Println(res.Error)
		return &notifpb.SubscribeNotificationResponse{
			Error:  res.Error.Error(),
			Status: http.StatusInternalServerError,
		}, nil
	}

	return &notifpb.SubscribeNotificationResponse{
		Status: http.StatusOK,
	}, nil
}
