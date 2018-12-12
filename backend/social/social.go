package social

import (
	"github.com/sea350/ustart_mono/backend/social/storage"
)

const ()

type Social struct {
	
	sqlClient  *sqlstore.SqlStore

}

func New(cfg *Config) (*Social, error) {
	sqlStor, err := sqlstore.New(cfg.Config)

	social := &Social{
		sqlClient: sqlStor,
	}
	return social, err
}


