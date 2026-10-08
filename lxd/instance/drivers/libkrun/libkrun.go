package libkrun

/*
#cgo linux LDFLAGS: -ldl -lpthread
#include <stdlib.h>
#include "libkrun_fwd.h"
*/
import "C"

// Log levels for InitLog.
const (
	LogLevelOff   = uint32(C.KRUN_LOG_LEVEL_OFF)
	LogLevelError = uint32(C.KRUN_LOG_LEVEL_ERROR)
	LogLevelWarn  = uint32(C.KRUN_LOG_LEVEL_WARN)
	LogLevelInfo  = uint32(C.KRUN_LOG_LEVEL_INFO)
	LogLevelDebug = uint32(C.KRUN_LOG_LEVEL_DEBUG)
	LogLevelTrace = uint32(C.KRUN_LOG_LEVEL_TRACE)
)

// Log styles for InitLog.
const (
	LogStyleAuto   = uint32(C.KRUN_LOG_STYLE_AUTO)
	LogStyleAlways = uint32(C.KRUN_LOG_STYLE_ALWAYS)
	LogStyleNever  = uint32(C.KRUN_LOG_STYLE_NEVER)
)

// LogOptionNoEnv disallows environment variables (e.g. RUST_LOG) from overriding the
// level/style passed to InitLog.
const LogOptionNoEnv = uint32(C.KRUN_LOG_OPTION_NO_ENV)

// InitLog initializes libkrun's internal logger. fd is the target file descriptor to write
// log output to, or -1 to use the default target (stderr). It must be called at most once,
// before CreateContext.
func InitLog(fd int, level uint32, style uint32, options uint32) error {
	return check(C.krun_init_log(C.int(fd), C.uint32_t(level), C.uint32_t(style), C.uint32_t(options)))
}

// Context is a libkrun configuration context used to build and start a single microVM.
type Context struct {
	id C.uint32_t
}

// CreateContext creates a new libkrun configuration context.
func CreateContext() (*Context, error) {
	ret := C.krun_create_ctx()
	if ret < 0 {
		return nil, errnoFromRet(ret)
	}

	return &Context{id: C.uint32_t(ret)}, nil
}

// Close releases the libkrun configuration context.
func (c *Context) Close() error {
	return check(C.krun_free_ctx(c.id))
}

// StartEnter starts and enters the microVM.
func (c *Context) StartEnter() error {
	return check(C.krun_start_enter(c.id))
}
