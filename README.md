# go-dbd-sqlite

go-dbd-sqlite is an sqlite driver implementation for
[go-dbi](http://github.com/thomaslee/go-dbi), a database interface for the
[Go](http://golang.org) programming language.

__WARNING: THIS IS PRE-ALPHA SOFTWARE__

## Status

Pre-Alpha.

## Overview

This package provides an sqlite driver for go-dbi. For more information about
go-dbi, please check out the [go-dbi](http://github.com) project on github.

## Installation

1. Make sure you have a working Go environment. See the
[install instructions](http://golang.org/doc/install.html).
2. Ensure goinstall is on your $PATH.
3. Install go-dbi: *$ goinstall github.com/thomaslee/go-dbi*
4. Install gosqlite: *$ goinstall gosqlite.googlecode.com/hg/sqlite*
5. Install go-dbd-sqlite: *$ goinstall github.com/thomaslee/go-dbd-sqlite*

Alternatively, you can clone the git repository & build it using make.

## Usage

    package main

    import (
        dbi "github.com/thomaslee/go-dbi"
        _ "github.com/thomaslee/go-dbd-sqlite"
        "log"
    )

    func main() {
        conn, err := dbi.Connect("sqlite://./foo.db")
        if err != nil {
            log.Printf("error: unable to connect to the database: %s", err.String())
        }
        defer conn.Close()
    }

## License

This software is licensed under the terms of the [MIT License](http://en.wikipedia.org/wiki/MIT_License).

## Support

Please log defects and feature requests using the issue tracker on github.

## About

go-dbd-sqlite was written by [Tom Lee](http://tomlee.co).

Follow me on [Twitter](http://www.twitter.com/tglee) or
[LinkedIn](http://au.linkedin.com/pub/thomas-lee/2/386/629).

