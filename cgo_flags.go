package rocksdb

import (
	_ "github.com/dckit/c-rocksdb"
)

// #cgo CPPFLAGS: -I ../c-rocksdb/internal/include
// #cgo CPPFLAGS: -I ../../cockroachdb/c-snappy/internal -I ../../cockroachdb/c-lz4/internal/lib
// #cgo darwin LDFLAGS: -Wl,-undefined -Wl,dynamic_lookup
// #cgo !darwin LDFLAGS: -Wl,-unresolved-symbols=ignore-all
import "C"
