package storage

import (
	"context"

	"github.com/sea350/ustart_micro/backend/project/auth/types"
)

// Storage is a database-agnostic interface for persisting auth data
type Storage interface {
	//New Member functions
	AddMember(context.Context, string, string, string, string) error                   //pass uuid, projectid, role, joindate | get error
	AddRole(context.Context, string, string, bool, bool, bool, bool, bool, bool) error //pass role, projectid, permissions: member, icon, banner, entries, links, tags | get error

	//Modifying existing members
	ModifyRoleManageMembers(context.Context, string, string, bool) error //pass projectid, role name, new value | get error
	ModifyRoleChangeIcon(context.Context, string, string, bool) error
	ModifyRoleChangeBanner(context.Context, string, string, bool) error
	ModifyRoleManageEntries(context.Context, string, string, bool) error
	ModifyRoleManageLinks(context.Context, string, string, bool) error
	ModifyRoleManageTags(context.Context, string, string, bool) error
	ModifyMemberRole(context.Context, string, string, string) error //pass user id, project id, role name | get error

	//Remove member functions
	RemoveMember(context.Context, string, string) error //pass user id, project id | get error
	RemoveRole(context.Context, string, string) error   //pass project id, role name | get error

	//Retreive Member functions
	GetMemberRole(context.Context, string, string) (string, error)      //pass user id, project id | get role name, error
	GetMembers(context.Context, string) ([]types.Member, error)         //pass project id | get a list of members, error
	GetRoleProfile(context.Context, string, string) (types.Role, error) //pass project id, profile name | get role profile, error
	GetProjectRoles(context.Context, string) ([]types.Role, error)      //pass project id | get list of role profiles, error

	//Member Request functions
	AddRequest(context.Context, string, string, string) error //pass user id, project id, and request date
	RemoveRequest(context.Context, string, string) error      //pass user id and project id
	GetRequest(context.Context, string)                       //pass project id
	AcceptRequest(context.Context, string, string, string)    //pass project id, newMember id, and user id

	//Errors
	ErrNoResultsFound() error
	ErrTooManyResults() error

	//TODO:
	//AggregateUserProjects(context.Context, string)([]string, error)
}
