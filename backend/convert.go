package backend

import (
	"github.com/sea350/ustart_micro/backend/backendpb"
	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

func convertProfileUserToBackendUser(profU profilepb.Profile) *backendpb.Profile {
	profB := backendpb.Profile{
		UUID:         profU.UUID,
		Username:     profU.Username,
		FirstName:    profU.FirstName,
		LastName:     profU.LastName,
		Description:  profU.Description,
		Organization: profU.Organization,
		Available:    profU.Available,
		Avatar:       profU.Avatar,
		Banner:       profU.Banner,
		DOB:          profU.DOB,
		Tags:         profU.Tags,
	}
	for _, record := range profU.Degrees {
		profB.Degrees = append(profB.Degrees, &backendpb.AcademicRecord{
			School:         record.School,
			Majors:         record.Majors,
			Minors:         record.Minors,
			Graduation:     record.Graduation,
			EducationLevel: record.EducationLevel,
		})
	}
	return &profB
}

func convertLiteResToLiteProfile(liteA *profilepb.LiteResponse) *backendpb.LiteProfile {
	liteB := backendpb.LiteProfile{
		UUID:      liteA.UUID,
		Username:  liteA.Username,
		FirstName: liteA.FirstName,
		LastName:  liteA.LastName,
		Avatar:    liteA.Avatar,
	}
	return &liteB
}
