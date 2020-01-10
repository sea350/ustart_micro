package storage

import (
	activstore "github.com/sea350/ustart_micro/backend/notification/storage/elastic/activity"
	notifstore "github.com/sea350/ustart_micro/backend/notification/storage/elastic/notif"
)

// Config determines the runtime behavior of the SQL or backed session manager
type Config struct {
	ESNotifConfig    *notifstore.Config
	ESActivityConfig *activstore.Config
}
