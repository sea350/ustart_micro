package storage

import (
	activstore "github.com/sea350/ustart_micro/backend/notification/storage/elastic/activity"
	notifstore "github.com/sea350/ustart_micro/backend/notification/storage/elastic/notif"
)

// NewNotifES determines the runtime behavior of the ElasticSearch-backed notification tracker
func NewNotifES(config *Config) (NotifStorage, error) {
	strg, err := notifstore.New(config.ESNotifConfig)
	return strg, err
}

// NewActivityES determines the runtime behavior of the ElasticSearch-backed activity tracker
func NewActivityES(config *Config) (ActivityStorage, error) {
	strg, err := activstore.New(config.ESActivityConfig)
	return strg, err
}
