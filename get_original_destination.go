package ciphernaut

import (
	"errors"
	"os"
	"unsafe"
)

// #include <errno.h>
// #include <string.h>
// #include <stdlib.h>
// #include "./get_original_destination.h"
import "C"

// Returns the address (ipv4 or ipv6 string) and port, that the connection had before it was re-routed or
// redirected. ipVersion needs to be 4 or 6. socket is the socket that was used to intercept the connection.
func GetOriginalDestination(ipVersion int, socket *os.File) (destination string, port int, err error) {
	var cDst *C.char
	var cPort C.int
	var cErr C.int
	var status C.int

	var fd uintptr = socket.Fd()

	if ipVersion == 4 {
		status = C.get_original_destination_4(C.int(fd), &cDst, &cPort, &cErr)
	} else if ipVersion == 6 {
		status = C.get_original_destination_6(C.int(fd), cDst, &cPort)
	} else {
		return "", -1, errors.New("Unknown IP version")
	}

	if status != 0 {
		if int(cErr) == 0 {
			// this hopefully/probably never happens
			return "", -1, errors.New("Failed to obtain original destination")
		}
		var cError *C.char = C.strerror(cErr)
		var why string = C.GoString(cError)
		return "", -1, errors.New("Failed to obtain original destination: " + why)
	}

	destination = C.GoString(cDst)
	port = int(cPort)
	err = nil

	C.free(unsafe.Pointer(cDst))
	return
}
