package dbd_sqlite

import (
    "os"
    _sqlite "gosqlite.googlecode.com/hg/sqlite"
)

type SqliteResultSet struct {
    stmt *_sqlite.Stmt
}

func newResultSet(stmt *_sqlite.Stmt) *SqliteResultSet {
    rs := new(SqliteResultSet)
    rs.stmt = stmt
    return rs
}

func (rs *SqliteResultSet) RowCount() (uint64, os.Error) {
    return 0, os.NewError("ResultSet.RowCount() is unsupported in go-dbd-sqlite")
}

func (rs *SqliteResultSet) Next() bool {
    return rs.stmt.Next()
}

func (rs *SqliteResultSet) Scan(refs ...interface{}) os.Error {
    return rs.stmt.Scan(refs...)
}

func (rs *SqliteResultSet) NamedScan(refs ...interface{}) os.Error {
    // TODO implement me
    return nil
}

func (rs *SqliteResultSet) Close() (err os.Error) {
    err = rs.stmt.Finalize()
    if err != nil {
        return
    }
    rs.stmt = nil
    return
}

