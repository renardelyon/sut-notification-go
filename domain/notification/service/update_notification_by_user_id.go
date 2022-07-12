package service

import (
	"context"
	"log"
	"net/http"
	"strings"
	notifpb "sut-notification-go/pb/notification"
)

func (s *Service) UpdateNotificationByUserId(ctx context.Context, req *notifpb.UpdateNotificationRequest) (*notifpb.UpdateNotificationResponse, error) {
	status := strings.ToLower(req.Status)
	qty := req.Quantity

	res := s.H.DB.Exec(`
		update notification 
		set ? = ? + ? 
		where user_id = '?'`,
		status, status, qty, req.UserId)
	if res.Error != nil {
		log.Println(res.Error)
		return &notifpb.UpdateNotificationResponse{
			Error:  res.Error.Error(),
			Status: http.StatusInternalServerError,
		}, nil
	}

	return &notifpb.UpdateNotificationResponse{
		Status: http.StatusOK,
	}, nil
}
