package global

import "net"

const (
	CNC_ADDR    = "10.0.0.67"
	MASTER_PORT = 50091
	VICTIM_PORT = 50001
	CONN_TYPE   = "tcp"
)

var (
	Users_role        = map[string]string{"Eddy\r\n": "Owner"}
	Users_cred        = map[string]string{"Eddy\r\n": "1234\r\n"}
	VictimsConnection []*net.Conn
)
