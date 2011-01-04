package dbd_sqlite

import (
    "os"
    _sqlite "gosqlite.googlecode.com/hg/sqlite"
)

type SqliteResultSet struct {
    stmt *_sqlite.Stmt
    numRows uint64
}

func newResultSet(stmt *_sqlite.Stmt) *SqliteResultSet {
    rs := new(SqliteResultSet)
    rs.stmt = stmt
    rs.numRows = 0

    //
    // hack to work around sqlite API
    //
    // XXX Reset() not officially supported by gosqlite
    //
    //for stmt.Next() {
    //    rs.numRows++
    //}
    //stmt.Reset()
    return rs
}

func (rs *SqliteResultSet) RowCount() uint64 {
    // XXX broken
    return rs.numRows
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

