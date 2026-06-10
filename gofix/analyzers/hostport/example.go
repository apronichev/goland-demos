// `hostport`: building an address with `fmt.Sprintf("%s:%d", …)` breaks for
// IPv6 (which needs `[host]:port`). `net.JoinHostPort` handles both families.
package hostport

import (
	"fmt"
	"net"
)

func Dial(host string, port int) (net.Conn, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	return net.Dial("tcp", addr)
}

func DialString(host, port string) (net.Conn, error) {
	return net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
}
