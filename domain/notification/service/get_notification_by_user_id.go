package service

import (
	"context"
	"net/http"
	"sut-notification-go/domain/notification/model"
	notifpb "sut-notification-go/pb/notification"
)

func (s *Service) GetNotificationByUserId(ctx context.Context, req *notifpb.GetNotificationRequest) (*notifpb.GetNotificationResponse, error) {
	var notifData model.Notification
	if res := s.H.DB.Where(&model.Notification{UserId: req.UserId}).First(&notifData); res.Error != nil {
		return &notifpb.GetNotificationResponse{
			Status: http.StatusNotFound,
			Error:  "user id doesn't exist",
		}, nil
	}

	return &notifpb.GetNotificationResponse{
		Accepted: int64(notifData.Accepted),
		Rejected: int64(notifData.Rejected),
		Status:   http.StatusOK,
	}, nil
}
