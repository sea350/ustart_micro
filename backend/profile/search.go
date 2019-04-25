package profile

import (
	"context"
	"strings"

	"github.com/sea350/ustart_micro/backend/profile/profilepb"
)

func split(r rune) bool {
	return r == ',' || r == ' '
}

// Search retreives a list of minimal profile data based off search queries
func (profile *Profile) Search(ctx context.Context, req *profilepb.SearchRequest) (*profilepb.SearchResponse, error) {

	queryArr := strings.FieldsFunc(req.Query, split)
	// filterArr := strings.FieldsFunc(req.Filters, split)

	//TODO:
	//enable filters
	//set up a real must map

	var searchUsername bool
	var searchName bool
	var searchTags bool

	searchUsername = true
	searchName = true
	searchTags = true

	ids, err := profile.strg.Search(ctx, queryArr, searchName, searchUsername, searchTags, make(map[string][]string), req.Scroll)
	if err != nil {
		return nil, err
	}

	var res []*profilepb.LiteResponse
	var resErr error
	for _, id := range ids {
		liteRes, err := profile.Lite(ctx, &profilepb.LiteRequest{UUID: id})
		if err != nil {
			res = append(res, liteRes)
		} else {
			resErr = ErrProblemLoadingProfile
		}
	}
	return &profilepb.SearchResponse{Results: res}, resErr
}
