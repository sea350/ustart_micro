package elasticstore

import (
	"context"

	"github.com/sea350/ustart_micro/backend/project/projectpb"
)

//Register creates a new ES document for a newly created project
func (estor *ElasticStore) Register(ctx context.Context, pid string, customURL string, name string, category string, avatar string, banner string, creationDate string, school string, tags []string, skillsNeeded []string, links []string) error {

	//Lock just to make sure no two people can register the same project url at the same time
	newProjectLock.Lock()
	defer newProjectLock.Unlock()

	_, err := estor.client.Index().
		Index(estor.eIndex).
		// Type(estor.eType).
		BodyJson(projectpb.Project{
			PID:          pid,
			CustomURL:    customURL,
			Name:         name,
			Category:     category,
			Avatar:       avatar,
			Banner:       banner,
			CreationDate: creationDate,
			School:       school,
			Tags:         tags,
			SkillsNeeded: skillsNeeded,
			Links:        links,
		}).
		Id(pid).
		Do(ctx)

	return err
}

// message Project {
//     string PID = 1;
//     string Name = 2;
//     repeated Member ListMembers = 3;
//     string CreationDate = 4;
//     string Category = 5;
//     string Description = 6;
//     string School = 7;
//     Location Location = 8;
//     repeated string Tags = 9;
//     string Avatar = 10;
//     string Banner = 11;
//     string CustomURL = 12;
//     repeated Member ListFollowers = 13;
//     repeated string SkillsNeeded = 14;
//     repeated string Links = 15;
