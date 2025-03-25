package elasticsearch

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// Driver is the driver for mysql database.
type Driver struct {
	*gdb.Core
}

const (
	quoteChar = "`"
)

func init() {
	var (
		err         error
		driverObj   = New()
		driverNames = g.SliceStr{"es", "elasticsearch"}
	)
	for _, driverName := range driverNames {
		if err = gdb.Register(driverName, driverObj); err != nil {
			panic(err)
		}
	}
}

// New create and returns a driver that implements gdb.Driver, which supports operations for MySQL.
func New() gdb.Driver {
	return &Driver{}
}

// New creates and returns a database object for mysql.
// It implements the interface of gdb.Driver for extra database driver installation.
func (d *Driver) New(core *gdb.Core, node *gdb.ConfigNode) (res gdb.DB, err error) {
	res = &Driver{
		Core: core,
	}
	return
}

// GetChars returns the security char for this type of database.
func (d *Driver) GetChars() (charLeft string, charRight string) {
	return quoteChar, quoteChar
}
