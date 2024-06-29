package cron

import (
	"fmt"
	"log"
	"time"

	"github.com/containrrr/shoutrrr"
	"github.com/mtricht/remedio/internal/config"
	"github.com/mtricht/remedio/internal/database"
	"gorm.io/gorm"
)

func RunCron() {
	now := time.Now().Format("15:04:00")
	var medication []database.Medication
	database.DB.Debug().Where(gorm.Expr("time(time) = ? AND supply > 0", now)).Find(&medication)
	if len(medication) == 0 {
		return
	}
	log.Printf("sending notification(s) for %d medication(s)", len(medication))
	for _, m := range medication {
		err := shoutrrr.Send(config.Config.NotificationURL, fmt.Sprintf("Time to take '%s'", m.Name))
		if err != nil {
			log.Println("failed to send notification %w", err)
			return
		}
		database.DB.Model(&m).Update("last_notification", time.Now())
	}
}
