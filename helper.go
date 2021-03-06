// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.

package daemon

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Service constants
const (
	success = "\t\t\t\t\t[  \033[32mOK\033[0m  ]" // Show colored "OK"
	failed  = "\t\t\t\t\t[\033[31mFAILED\033[0m]" // Show colored "FAILED"
)

var (
	ErrUnsupportedSystem = errors.New("Unsupported system")
	ErrRootPriveleges    = errors.New("You must have root user privileges. Possibly using 'sudo' command should help")
	ErrAlreadyInstalled  = errors.New("Service has already been installed")
	ErrNotInstalled      = errors.New("Service is not installed")
	ErrAlreadyStopped    = errors.New("Service has already been stopped")
	ErrAlreadyRunning    = errors.New("Service is alredy running")
)

// Lookup path for executable file
func executablePath(name string) (string, error) {
	if path, err := exec.LookPath(name); err == nil {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			return execPath()
		}
		return path, nil
	}
	return execPath()
}

// Check root rights to use system service
func checkPrivileges() (bool, error) {

	if output, err := exec.Command("id", "-g").Output(); err == nil {
		if gid, parseErr := strconv.ParseUint(strings.TrimSpace(string(output)), 10, 32); parseErr == nil {
			if gid == 0 {
				return true, nil
			}
			return false, ErrRootPriveleges
		}
	}
	return false, ErrUnsupportedSystem
}
