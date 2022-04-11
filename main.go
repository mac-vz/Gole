package gole

import (
	"fmt"
	"os"
	"time"
)

const VERSION string = "1.2.1"

func Main(args []string) {
	fmt.Printf("Gole v%s\n", VERSION)
	conf := ParseConfig(args)
	switch conf.getMode() {
	case "tcp":
		fmt.Printf("tunnel mode: TCP\n")
		fmt.Printf("operation: %s\n", conf.getOp())
		PrintDbgf("%v\n", conf.(*TCPConfig))
	case "udp":
		fmt.Printf("tunnel mode: UDP\n")
		fmt.Printf("tunnel protocol: %s\n", conf.(*UDPConfig).Proto)
		fmt.Printf("operation: %s\n", conf.getOp())
		PrintDbgf("%v\n", conf.(*UDPConfig))
	default:
		fmt.Printf("not implemented\n")
		os.Exit(1)
	}

	// punch hole
	fmt.Println("====================")
	fmt.Printf("punching holes: [local]%s ---> [remote]%s\n", conf.LocalAddr(), conf.RemoteAddr())
	conn, err := Punch(conf)
	if err != nil {
		perror("Failed to punch hole.", err)
		os.Exit(1)
	}
	fmt.Printf("punched through\n")
	time.Sleep(50)
	fmt.Printf("%s %s\n", conn.LocalAddr(), conf.RemoteAddr())
	if conf.getOp() == "holepunch" {
		os.Exit(0)
	}

	// create tunnel
	fmt.Println("====================")
	fmt.Printf("creating tunnel: [local]%s <--> [remote]%s\n", conn.LocalAddr(), conf.RemoteAddr())

	if conf.getOp() == "client" {
		fmt.Println("starting client ...")
		StartClient(conn, conf)
	} else if conf.getOp() == "server" {
		fmt.Println("starting server ...")
		StartServer(conn, conf)
	}
	fmt.Printf("Done\n")
}
