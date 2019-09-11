package types

import "time"

//Member is all data relevant to the authentication of a project member
type Member struct {
	UUID     string
	RoleName string
	JoinDate time.Time
}

//Role is a complete profile of a role
type Role struct {
	Name          string
	ManageMembers bool
	ChangeIcon    bool
	ChangeBanner  bool
	ManagePosts   bool
	ManageLinks   bool
	ManageTags    bool
}
