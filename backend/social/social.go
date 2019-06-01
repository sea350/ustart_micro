package social

import (
	sqlstore "github.com/sea350/ustart_micro/backend/social/storage"
)

const ()

//Social *insert comment here*
type Social struct {
	sqlClient *sqlstore.SqlStore
}

//New *insert comment here*
func New(cfg *Config) (*Social, error) {
	sqlStor, err := sqlstore.New(cfg.SQLStoreConfig)

	social := &Social{
		sqlClient: sqlStor,
	}
	return social, err
}
