package storage

import (
	"context"

	notifpb "github.com/sea350/ustart_micro/backend/notification/notificationpb"
)

// NotifStorage is a database-agnostic interface for notificaion data
type NotifStorage interface {
	NewNotif(context.Context, string, string, string) error                               //pass uuid, activity, timestamp | get error
	ScrollNotifs(context.Context, string, string) ([]notifpb.Notification, string, error) //pass uuid, scrollID | get notifications, scrollID, error
	ChangeSeen(context.Context, string, bool) error                                       //pass notifID, newValue | get error
	DeleteNotif(context.Context, string) error                                            //pass notifID | get error

	ErrNoNotifsFound() error
	ErrNotifNotFound() error
}

// ActivityStorage is a database-agnostic interface for activity data
type ActivityStorage interface {
	NewActivity(context.Context, string, string, string, string) (string, error)                 //pass actorUUID, objectID, action, timestamp | get activityID, error
	ScrollNotifActivities(context.Context, []string, string) ([]notifpb.Activity, string, error) //pass ALLactivityIDs, scrollID | get activities, scrollID, error
	ScrollUserActivities(context.Context, string, string) ([]notifpb.Activity, string, error)    //pass uuid, scrollID | get activities, scrollID, error

	ErrActivityNotFound() error
}
