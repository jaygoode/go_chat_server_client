package main

import (
	"bufio"
	"flag"
	"log"

	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/remote"
)

type client struct {
	username string
	serverPID *actor.PID
}

func newClient(username string, serverPID *actor.PID) actor.Producer {
	return func() actor.Receiver {
		return &client{
			username: username,
			serverPID: serverPID,
		}
	}
}

func (c *client) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
	case actor.Stopped:
		_ = msg
	}
}

func main() {
var (
	port = flag.String("port", ":3000", "")
	username = flag.String("username", "", "")
)
flag.Parse()

	e, _ := actor.NewEngine()
	rem := remote.New(e, remote.Config{
		ListenAddr: "127.0.0.1" + *port,
	})
	e.WithRemote(rem)

	serverPID := actor.NewPID("127.0.0.1:4000", "server")
	e.Spawn(newClient(*username, serverPID), "client")

	r := bufio.NewReader(os.stdin)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Errorw("failed to read message from stdin", log.M{"err": err})
		}
		_ = msg
	}
}
