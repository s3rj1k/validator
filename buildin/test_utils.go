package buildin

import (
	"net"
	"os"
	"syscall"
)

const (
	notExists   = "/tmp/testNotExists"
	regularFile = "/tmp/testRegularFile"

	dir        = "/tmp/testDir"
	anotherDir = "/tmp/testAnotherDir"

	notDir       = "/tmp/testNotDir"
	notStickyDir = "/tmp/testNotStickyDir"
	writableDir  = "/tmp/testWritableDir"
	stickyDir    = "/tmp/testStickyDir"

	notWritableFile = "/tmp/testNotWritableFile"
	writableFile    = "/tmp/testWritableFile"

	fileWithSetGID = "/tmp/testFileWithSetGID"
	fileWithSetUID = "/tmp/testFileWithSetUID"

	namedPipe  = "/tmp/testNamedPipe"
	unixSocket = "/tmp/testUnixSocket"

	executableFile    = "/tmp/testExecutableFile"
	notExecutableFile = "/tmp/testNotExecutableFile"

	symlink        = "/tmp/testSymlink"
	anotherSymlink = "/tmp/testAnotherSymlink"
	notSymlink     = "/tmp/testNotSymlink"

	symlinkTargetNotExist = "/tmp/testSymlinkTargetNotExist"
	symlinkTargetIsFile   = "/tmp/testSymlinkTargetIsFile"
	symlinkTargetIsDir    = "/tmp/testSymlinkTargetIsDir"
)

// creates but panics if already exists and on any errors
func _createfile(path string) {
	_, err := os.Stat(path)
	if os.IsExist(err) {
		panic(err)
	}

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}

// creates but panics if already exists and on any errors
func _createfiles(paths ...string) {
	for _, p := range paths {
		_createfile(p)
	}
}

// creates path with mode
func _createfilewithmode(path string, mode os.FileMode) { // nolint: unused,deadcode
	_createfile(path)
	_setfilemode(path, mode)
}

// silently attempts to delete path
func _silentdelete(path string) {
	_ = os.Remove(path)
	_ = syscall.Unlink(path)
}

// silently attempts to delete paths
func _silentdeleteMany(paths ...string) {
	for _, p := range paths {
		_silentdelete(p)
	}
}

// creates but panics if already exists
func _createdir(path string) {
	_, err := os.Stat(path)
	if os.IsExist(err) {
		panic(err)
	}

	err = os.Mkdir(path, 0777)
	if err != nil {
		panic(err)
	}
}

// creates but panics if already exists
func _createdirs(paths ...string) {
	for _, p := range paths {
		_createdir(p)
	}
}

// sets file mode and panics on error
func _setfilemode(path string, mode os.FileMode) {
	err := os.Chmod(path, mode)
	if err != nil {
		panic(err)
	}
}

// creates named pipe at path
func _createnamedpipe(path string) { // nolint: unparam
	err := syscall.Mkfifo(path, 0644)
	if err != nil {
		panic(err)
	}
}

func _creatunixsocket(path string) { // nolint: unparam
	_, err := net.Listen("unix", path)
	if err != nil {
		panic(err)
	}
}
