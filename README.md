ROCKSDB
=======


## A Go wrapper for Rocksdb

[![Circle CI](https://circleci.com/gh/dckit/rocksdb.svg?style=svg)](https://circleci.com/gh/dckit/rocksdb) [![GoDoc](https://godoc.org/github.com/dckit/rocksdb?status.svg)](https://godoc.org/github.com/dckit/rocksdb)


# Install

All dependencies are `go get-able`

```
go get github.com/cockroachdb/c-snappy
go get github.com/cockroachdb/c-lz4
go get github.com/dckit/c-rocksdb
go get github.com/dckit/rocksdb
```


# Caveats

Comparators and WriteBatch iterators must be written in C in your own
library. This seems like a pain in the ass, but remember that you'll have the
LevelDB C API available to your in your client package when you import levigo.

An example of writing your own Comparator can be found in
<https://github.com/jmhodges/levigo/blob/master/examples>.
