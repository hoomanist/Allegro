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
	config string
	fs     *flag.FlagSet
}
type MigrateCommand struct {
	fs     *flag.FlagSet
	config string
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
	sc.fs.StringVar(&sc.config, "cfg", "./config.ini", "what is the config ?")
	return sc
}

func NewMigrateCommand() *MigrateCommand {
	mc := &MigrateCommand{
		fs: flag.NewFlagSet(
			"migrate",
			flag.ContinueOnError,
		),
	}
	mc.fs.StringVar(&mc.config, "cfg", "./config.ini", "what is the config?")
	return mc
}

func (mc *MigrateCommand) Name() string {
	return mc.fs.Name()
}
func (mc *MigrateCommand) Run() error {
	c := Configure(mc.config)
	return database.Migrate(c.DB)
}
func (mc *MigrateCommand) Init(args []string) error {
	return mc.fs.Parse(args)
}

func (sc *ServeCommand) Name() string {
	return sc.fs.Name()
}
func (sc *ServeCommand) Run() error {
	c := Configure(sc.config)
	server.Serve(c.Port)
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
