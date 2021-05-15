package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"

	rcon "github.com/Tnze/go-mc/net"
	"go.uber.org/zap"
)

func main() {
	port := flag.Uint("port", 25575, "RCON server port")
	password := flag.String("password", "minecraft", "RCON password")
	flag.Parse()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	log := logger.Sugar()
	defer log.Sync()

	cmd := exec.Command(flag.Arg(0))
	cmd.Env = append(cmd.Env, "LD_LIBRARY_PATH=.")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	// binds mc-becrock-runner stdin to bedrock server's stdin
	go func() {
		io.Copy(stdin, os.Stdin)
	}()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	l, err := rcon.ListenRCON(":" + strconv.Itoa(int(*port)))
	if err != nil {
		log.Fatal(err)
	}

	log.Info(fmt.Sprintf("Listening on port %d...", *port))

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Error(err)
		} else {
			go handleConenction(conn, log, *password, stdin)
		}
	}
}

// Handles RCON connections. Commands writtern to connection will be piped to stdin of the running bedrock server.
// Connections will be kept open until the client closes or an error occurs writing to bedrock server stdin pipe
func handleConenction(conn rcon.RCONServerConn, log *zap.SugaredLogger, password string, stdin io.WriteCloser) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Debug(err)
		}
	}()
	if err := conn.AcceptLogin(password); err != nil {
		log.Debug(err)
		return
	}

	for {
		cmd, err := conn.AcceptCmd()
		if err != nil {
			log.Debug(err)
			if err := conn.RespCmd(err.Error()); err != nil {
				log.Debug(err)
			}
			return
		}

		if _, err := stdin.Write([]byte(cmd + "\n")); err != nil {
			log.Debug(err)
			if err := conn.RespCmd("command failed"); err != nil {
				log.Debug(err)
			}
			return
		}

		if err := conn.RespCmd(cmd + " command recieved"); err != nil {
			log.Debug(err)
			return
		}
	}

}
