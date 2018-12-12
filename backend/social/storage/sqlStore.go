package sqlstore

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"

)


//sqlStore will be a wrapper for any sql request made to the db
type SqlStore struct {
	DB  *sql.DB
}

// 
func New(cfg *Config) (*SqlStore, error) {

	dbInfo := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port)

	_ = pq.Efatal

	db, err := sql.Open("postgres",dbInfo)
	if err != nil {
		/*log something here*/
	  }
	sqlDB := &SqlStore{
		DB : db,
	}

	// set up initializing functions for both tables

	_,err = db.Exec("CREATE TABLE IF NOT EXISTS Block ( blocker numeric(9,2) NOT NULL, blockertype numeric(9,2) NOT NULL, blocked numeric(9,2) NOT NULL, blockedtype numeric(9,2) NOT NULL, timestamps numeric(9,2) NOT NULL, PRIMARY KEY( blocked ) PRIMARY KEY( blocker ) )")
   if err != nil {
       panic(err)
   }

   _,err = db.Exec("CREATE TABLE IF NOT EXISTS Follow ( follower numeric(9,2) NOT NULL, followertype numeric(9,2) NOT NULL, followed numeric(9,2) NOT NULL, followedtype numeric(9,2) NOT NULL, timestamps numeric(9,2) NOT NULL, PRIMARY KEY( follower ) PRIMARY KEY( followed ))")
if err != nil {
   panic(err)
}

	err = db.Ping()
	return sqlDB, err
}