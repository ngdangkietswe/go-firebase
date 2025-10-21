/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package util

import (
	"fmt"
	"log"
	"runtime/debug"
)

func RecoverPanic() {
	if r := recover(); r != nil {
		var errMsg string
		switch v := r.(type) {
		case error:
			errMsg = v.Error()
		default:
			errMsg = fmt.Sprintf("%v", v)
		}

		log.Printf("[PANIC RECOVERED] %s\n--- Stack Trace ---\n%s", errMsg, debug.Stack())
	}
}
