package dbd_sqlite

import (
    "http"
    "os"
    "github.com/thomaslee/go-dbi"
    _sqlite "gosqlite.googlecode.com/hg/sqlite"
)

type SqliteDriver struct {
}

func (self *SqliteDriver) Connect(dsn *http.URL) (conn dbi.Connection, err os.Error) {
    c, err := _sqlite.Open(dsn.Path[1:])
    if err != nil {
        return
    }
    conn = newConnection(c)
    return
}

func init() {
    dbi.AddDriver("sqlite", new(SqliteDriver))
}

