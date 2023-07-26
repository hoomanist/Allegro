package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/hoomanist/allegro-server/pkg/server"
)

type ServeCommand struct {
	port string
	fs   *flag.FlagSet
}

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func NewServeCommand() *ServeCommand {
	sc := &ServeCommand{
		fs: flag.NewFlagSet(
			"serve",
			flag.ContinueOnError,
		),
	}
	sc.fs.StringVar(&sc.port, "port", "8000", "what should be the port")
	return sc
}

func (sc *ServeCommand) Name() string {
	return sc.fs.Name()
}
func (sc *ServeCommand) Run() error {
	server.Serve(sc.port)
	return nil
}
func (sc *ServeCommand) Init(args []string) error {
	return sc.fs.Parse(args)
}
func Root(args []string) error {
	if len(args) < 1 {
		return errors.New("you must pass a sub-command")
	}
	cmds := []Runner{
		NewServeCommand(),
	}
	subcommand := os.Args[1]
	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}
	return fmt.Errorf("unkown sub-command: %s", subcommand)
}
