package libcore

import (
	"io"
	"net"
	"os"

	"golang.org/x/sys/unix"
)

const (
	ProtectFailed byte = iota
	ProtectSuccess
)

func getOneFd(socketFd int) (int, error) {
	buf := make([]byte, unix.CmsgSpace(4))
	_, _, _, _, err := unix.Recvmsg(socketFd, nil, buf, 0)
	if err != nil {
		return 0, err
	}
	msgs, _ := unix.ParseSocketControlMessage(buf)

	if len(msgs) != 1 {
		return 0, newError("invalid msgs count: ", len(msgs))
	}
	fds, _ := unix.ParseUnixRights(&msgs[0])
	if len(fds) != 1 {
		return 0, newError("invalid fds count: ", len(fds))
	}
	return fds[0], nil
}

func ServerProtect(path string, protector Protector) io.Closer {
	_ = os.Remove(path)
	l, err := net.ListenUnix("unix", &net.UnixAddr{Name: path, Net: "unix"})
	if err != nil {
		return nil
	}
	_ = os.Chmod(path, 0o777)
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				return
			}
			go func(conn *net.UnixConn) {
				defer conn.Close()
				rawConn, err := conn.SyscallConn()
				if err != nil {
					return
				}
				var (
					connFd int
					recvFd int
				)
				err = rawConn.Control(func(fd uintptr) {
					connFd = int(fd)
					recvFd, err = getOneFd(connFd)
				})
				if err != nil {
					return
				}
				defer unix.Close(connFd)
				if !protector.Protect(int32(recvFd)) {
					_, _ = conn.Write([]byte{ProtectFailed})
					return
				}
				_, _ = conn.Write([]byte{ProtectSuccess})
			}(conn.(*net.UnixConn))
		}
	}()
	return l
}
