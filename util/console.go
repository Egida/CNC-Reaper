package util

import (
	"fmt"
	"mirage/global"
	"net"
)

func ReturnBanner(username string) string {
	return fmt.Sprintf("%s\r\n%s\r\n%s\r\n%s\r\n%s\r\n%s\r\n%s\r\n%s\r\n%s\r\n%s\r\n%s%s%s%s%s\r\n%s\r\n%s\r\n%s\r\n",
		`                  |---.\											`,
		"          ___     |    `											",
		"         / .-\\  ./=)												",
		`        |  |"|_/\/|		█▄▄▄▄ ▄███▄   ██   █ ▄▄  ▄███▄   █▄▄▄▄	`,
		"        ;  |-;| /_|		█  ▄▀ █▀   ▀  █ █  █   █ █▀   ▀  █  ▄▀	",
		`       / \_| |/ \ |		█▀▀▌  ██▄▄    █▄▄█ █▀▀▀  ██▄▄    █▀▀▌	`,
		`      /      \/\( |		█  █  █▄   ▄▀ █  █ █     █▄   ▄▀ █  █	`,
		"      |   /  |` ) |		  █   ▀███▀      █  █    ▀███▀     █    ",
		`      /   \ _/    |		▀              █    ▀            ▀  	`,
		`     /--._/  \    |												`,
		"     `/|)    |    /	 	   ╔══════════════════════~		     	",
		"       /     |   |	           ║    Welcome ", username,
		"     .'      |   |		   ║    Graded as ", global.Users_role[username],
		`    /         \  |		   ╚══════════════════════~`,
		"   (_.-.__.__./  /		  				    						",
		"																	")
}

func ClearScreen(conn *net.Conn) {
	(*conn).Write([]byte("\033[2J"))
}

func SetTitle(conn *net.Conn) {
	(*conn).Write([]byte(fmt.Sprintf("\033]0;Welcome to Reaper Botnet | %d zombies\007", len(global.VictimsConnection))))
}
