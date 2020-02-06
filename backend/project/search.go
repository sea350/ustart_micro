package project

import (
	"context"
	"strings"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

func split(r rune) bool {
	return r == ',' || r == ' '
}

// Search retreives a list of minimal profile data based off search queries
func (project *Project) Search(ctx context.Context, req *projectpb.SearchRequest) (*projectpb.SearchResponse, error) {

	queryArr := strings.FieldsFunc(req.Query, split)
	// filterArr := strings.FieldsFunc(req.Filters, split)

	//TODO:
	//enable filters
	//set up a real must map

	var searchCustomURL bool
	var searchName bool
	var searchTags bool

	searchCustomURL = true
	searchName = true
	searchTags = true

	ids, err := project.strg.Search(ctx, queryArr, searchName, searchCustomURL, searchTags, make(map[string][]string), req.Scroll)
	if err != nil {
		return nil, err
	}

	var res []*projectpb.LiteResponse
	var resErr error
	for _, id := range ids {
		liteRes, err := project.Lite(ctx, &projectpb.LiteRequest{ProjectID: id})
		if err != nil {
			res = append(res, liteRes)
		} else {
			resErr = ErrProblemLoadingProject
		}
	}
	return &projectpb.SearchResponse{Results: res}, resErr
}
