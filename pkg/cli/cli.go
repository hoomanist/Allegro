package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/hoomanist/allegro-server/pkg/database"
	"github.com/hoomanist/allegro-server/pkg/server"
)

type ServeCommand struct {
	fs *flag.FlagSet
}
type MigrateCommand struct {
	fs *flag.FlagSet
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
	return sc
}

func NewMigrateCommand() *MigrateCommand {
	mc := &MigrateCommand{
		fs: flag.NewFlagSet(
			"migrate",
			flag.ContinueOnError,
		),
	}
	return mc
}

func (mc *MigrateCommand) Name() string {
	return mc.fs.Name()
}
func (mc *MigrateCommand) Run() error {
	return database.Migrate()
}
func (mc *MigrateCommand) Init(args []string) error {
	return mc.fs.Parse(args)
}

func (sc *ServeCommand) Name() string {
	return sc.fs.Name()
}
func (sc *ServeCommand) Run() error {
	fmt.Println(os.Getenv("port"))
	server.Serve()
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
		NewMigrateCommand(),
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
