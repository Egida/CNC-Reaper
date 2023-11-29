package util

import "net"

func Contains(slice []*net.Conn, value *net.Conn) bool {
	for _, val := range slice {
		if *val == *value {
			return true
		}
	}
	return false
}
