package service

import (
	"context"
	"log"
	"net/http"
	"sut-notification-go/domain/notification/model"
	notifpb "sut-notification-go/pb/notification"
)

func (s *Service) ResetNotificationStatusQty(ctx context.Context, req *notifpb.ResetNotificationRequest) (*notifpb.ResetNotificationResponse, error) {
	res := s.H.DB.Model(model.Notification{}).
		Where("user_id = '?'", req.UserId).
		Updates(model.Notification{Rejected: 0, Accepted: 0})
	if res.Error != nil {
		log.Println(res.Error)
		return &notifpb.ResetNotificationResponse{
			Error:  res.Error.Error(),
			Status: http.StatusInternalServerError,
		}, nil
	}

	return &notifpb.ResetNotificationResponse{
		Status: http.StatusOK,
	}, nil
}
