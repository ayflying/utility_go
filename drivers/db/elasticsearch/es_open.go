package elasticsearch

import (
	"database/sql"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (d *Driver) Open(config *gdb.ConfigNode) (db *sql.DB, err error) {
	var (
		source               string
		underlyingDriverName = "elasticsearch"
	)
	source = config.Host

	cfg := elasticsearch.Config{
		Addresses: []string{
			config.Host,
		},
	}

	es, err = elasticsearch.NewTypedClient(cfg)
	//if err != nil {
	//	fmt.Printf("elasticsearch.NewTypedClient failed, err:%v\n", err)
	//	return
	//}

	if db, err = sql.Open(underlyingDriverName, source); err != nil {
		err = gerror.WrapCodef(
			gcode.CodeDbOperationError, err,
			`sql.Open failed for driver "%s" by source "%s"`, underlyingDriverName, source,
		)
		return nil, err
	}

	return
}
