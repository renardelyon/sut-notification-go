package service

import (
	"context"
	"log"
	"net/http"
	notifpb "sut-notification-go/pb/notification"
)

func (s *Service) ResetNotificationStatusQty(ctx context.Context, req *notifpb.ResetNotificationRequest) (*notifpb.ResetNotificationResponse, error) {
	res := s.H.DB.Exec("update notifications set rejected = 0, accepted = 0 where user_id = ?", req.UserId)
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
