// +build static

package git

/*
#cgo windows CFLAGS: -I${BUILD_DIR}/include/
#cgo windows LDFLAGS: -L${BUILD_DIR}/lib/ -lgit2 -lwinhttp
#cgo !windows pkg-config: --static ${BUILD_DIR}/lib/pkgconfig/libgit2.pc
#include <git2.h>

#if LIBGIT2_VER_MAJOR != 1 || LIBGIT2_VER_MINOR != 0
# error "Invalid libgit2 version; this git2go supports libgit2 v1.0"
#endif

*/
import "C"
