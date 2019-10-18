package storage

import (
	"context"
	"time"
)

// NotifStorage is a database-agnostic interface for notificaion data
type NotifStorage interface {
	NewNotification(context.Context, string, string, time.Time) error            //pass uuid, activityID, timestamp | get error
	ScrollNotifications(context.Context, string, string) (string, string, error) //pass uuid, scrollID | get sessionID, expiration, error
	ChangeNotifSeen(context.Context, string, bool) error                         //pass notifID, newValue | get error
	RemoveNotif(context.Context, string, string) error                           //pass sessionID, expiration | get error

	ErrNoNotifsFound() error
	ErrNotifNotFound() error
}

// ActivityStorage is a database-agnostic interface for activity data
type ActivityStorage interface {
	NewActivity(context.Context, string, string, string, time.Time) (string, error) //pass actorUUID, objectID, action, timestamp | get activityID, error
	ScrollActivities(context.Context, string, string) (string, string, error)       //pass uuid, scrollID | get sessionID, expiration, error

	ErrNoActivitiesFound() error
	ErrActivityNotFound() error
}
