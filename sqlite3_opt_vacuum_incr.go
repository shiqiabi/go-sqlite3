// Copyright (C) 2014 Yasuhiro Matsumoto <mattn.jp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
// +build vacuum_incr

package sqlite3

/*
#cgo CFLAGS: -DSQLITE_DEFAULT_AUTOVACUUM=2
#cgo LDFLAGS: -lm
*/
import "C"
