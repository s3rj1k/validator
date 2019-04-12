package buildin

import (
	"net"
	"os"
	"syscall"
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
func _createfilewithmode(path string, mode os.FileMode) {
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
func _createnamedpipe(path string) {
	err := syscall.Mkfifo(path, 0644)
	if err != nil {
		panic(err)
	}
}

func _creatunixsocket(path string) {
	_, err := net.Listen("unix", path)
	if err != nil {
		panic(err)
	}
}
