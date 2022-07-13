package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	notifpb "sut-notification-go/pb/notification"
)

func (s *Service) UpdateNotificationByUserId(ctx context.Context, req *notifpb.UpdateNotificationRequest) (*notifpb.UpdateNotificationResponse, error) {
	for key, value := range req.StatusQtyMap {
		status := strings.ToLower(value.Status)
		res := s.H.DB.Exec(fmt.Sprintf(
			`update notifications 
			set %s = %s + %d 
			where user_id = '%s'`,
			status, status, value.Quantity, key))
		if res.Error != nil {
			log.Println(res.Error)
			return &notifpb.UpdateNotificationResponse{
				Error:  res.Error.Error(),
				Status: http.StatusInternalServerError,
			}, nil
		}
	}

	return &notifpb.UpdateNotificationResponse{
		Status: http.StatusOK,
	}, nil
}
