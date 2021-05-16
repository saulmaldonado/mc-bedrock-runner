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

	"github.com/saulmaldonado/mc-bedrock-runner/pkg/signal"
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

	stop := signal.SetupSignalHandler(log)

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

	cmdStop := serverExitHandler(cmd, log)

	l, err := rcon.ListenRCON(":" + strconv.Itoa(int(*port)))
	if err != nil {
		log.Fatal(err)
	}

	rconStop := make(chan struct{})

	go mainExitHandler(stop, cmdStop, stdin, log, l, rconStop)

	log.Info(fmt.Sprintf("Listening on port %d...", *port))

	conns := make(chan rcon.RCONServerConn)

	// Start a goroutine to accept new connnection and push then into a chan
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				select {
				case <-rconStop:
					// if the rcon listener has been closed, silently break out of the goroutine,
					// otherwise log the error
					return
				default:
					log.Error(err)
				}
			} else {
				conns <- conn
			}
		}
	}()

	for {
		select {
		case <-rconStop:
			stopRCON(l, log)
			log.Info("done")
			os.Exit(0)
		case conn := <-conns:
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

// Sets up a an exit handler for the forked bedrock server process in a goroutine.
// Returns a chan that will get sent an empty struct when the process exits
func serverExitHandler(cmd *exec.Cmd, log *zap.SugaredLogger) chan struct{} {
	cmdStop := make(chan struct{}, 1)

	go func() {
		err := cmd.Wait()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				exitCode := exitErr.ExitCode()
				log.Errorw("bedrock server exited with error", "exitCode", exitCode, "error", err.Error())
			} else {
				log.Errorw("bedrock server process failed", "error", err.Error())
			}
		}
		close(cmdStop)
	}()

	return cmdStop
}

// Will wait for one of the two exit handler channels to unblock. If the main process stops unblocks a "stop"
// command will be sent to the bedrock server and will wait until the server exits.
// If the bedrock process exits the RCON listener will be closed and the main process will exit
func mainExitHandler(stop, cmdStop chan struct{}, stdin io.WriteCloser, log *zap.SugaredLogger, rcon *rcon.RCONListener, rconStop chan struct{}) {
	for {
		select {
		case <-stop:
			stopServer(stdin, log)
		case <-cmdStop:
			close(stop)
			close(rconStop)
			return
		}
	}
}

// Sends "stop" command to bedrock server stdin
func stopServer(stdin io.WriteCloser, log *zap.SugaredLogger) {
	log.Info("writing stop to bedrock server...")
	if _, err := stdin.Write([]byte("stop\n")); err != nil {
		log.Error("failed to write \"stop\" to server")
	}
}

// Closes RCON listener
func stopRCON(rcon *rcon.RCONListener, log *zap.SugaredLogger) {
	if err := rcon.Close(); err != nil {
		log.Errorw("error closing RCON listener", "error", err.Error())
	} else {
		log.Info("RCON listener closed")
	}
}
