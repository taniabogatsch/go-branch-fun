//go:build static_lib

package gobf

/*
#cgo CPPFLAGS: -DDUCKDB_STATIC_BUILD
#include <duckdb.h>
*/
import "C"

// FIXME: maybe, instead of passing static_lib, we want this as the default?
