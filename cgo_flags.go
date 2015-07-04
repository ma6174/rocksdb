package rocksdb

import (
	_ "github.com/dckit/c-rocksdb"
)

// #cgo CXXFLAGS: -std=c++11
// #cgo CPPFLAGS: -I ../c-rocksdb/internal/include
// #cgo CPPFLAGS: -I ../../cockroachdb/c-snappy/internal -I ../../cockroachdb/c-lz4/internal/lib
// #cgo CPPFLAGS: -DROCKSDB_PLATFORM_POSIX -DNDEBUG -DSNAPPY -DLZ4
// #cgo darwin LDFLAGS: -Wl,-undefined -Wl,dynamic_lookup
// #cgo !darwin LDFLAGS: -Wl,-unresolved-symbols=ignore-all
import "C"
