package dbd_sqlite

import (
    "os"
    dbi "github.com/thomaslee/go-dbi"
    _sqlite "gosqlite.googlecode.com/hg/sqlite"
)

type SqliteConnection struct {
    impl *_sqlite.Conn
}

func newConnection(impl *_sqlite.Conn) *SqliteConnection {
    conn := new(SqliteConnection)
    conn.impl = impl
    return conn
}

func (self *SqliteConnection) Query(sql string, params ...interface{}) (rs dbi.ResultSet, err os.Error) {
    stmt, err := self.impl.Prepare(sql)
    if err != nil {
        return
    }
    err = stmt.Exec(params...)
    if err != nil {
        stmt.Finalize()
        return
    }
    rs = newResultSet(stmt)
    return
}

func (self *SqliteConnection) Execute(sql string, params ...interface{}) os.Error {
    return self.impl.Exec(sql, params...)
}

func (self *SqliteConnection) Prepare(sql string) (dbi.Statement, os.Error) {
    // TODO
    return nil, nil
}

func (self *SqliteConnection) BeginTransaction() os.Error {
    // TODO
    return nil
}

func (self *SqliteConnection) Commit() os.Error {
    // TODO
    return nil
}

func (self *SqliteConnection) Rollback() os.Error {
    // TODO
    return nil
}

func (self *SqliteConnection) Close() os.Error {
    return self.impl.Close()
}


