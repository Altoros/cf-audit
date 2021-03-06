package cmd

import (
	"bytes"
	// "errors"
	"fmt"
	// "reflect"
	"strings"

	// Should only be imported here to avoid leaking use of goflags through project
	boshcmd "github.com/cloudfoundry/bosh-cli/cmd"
	goflags "github.com/jessevdk/go-flags"
)

type Factory struct {
	deps boshcmd.BasicDeps
}

func NewFactory(deps boshcmd.BasicDeps) Factory {
	return Factory{deps: deps}
}

func (f Factory) New(args []string) (Cmd, error) {
	var cmdOpts interface{}

	commandOpts := &CommandOpts{}

	commandOpts.VersionOpt = func() error {
		return &goflags.Error{
			Type:    goflags.ErrHelp,
			Message: fmt.Sprintf("version %s\n", VersionLabel),
		}
	}

	parser := goflags.NewParser(commandOpts, goflags.HelpFlag|goflags.PassDoubleDash)

	parser.CommandHandler = func(command goflags.Commander, extraArgs []string) error {
		// if opts, ok := command.(*DiffOpts); ok {
		// 	// opts.Args =
		// }

		if len(extraArgs) > 0 {
			errMsg := "Command '%T' does not support extra arguments: %s"
			return fmt.Errorf(errMsg, command, strings.Join(extraArgs, ", "))
		}

		cmdOpts = command

		return nil
	}

	helpText := bytes.NewBufferString("")
	parser.WriteHelp(helpText)

	_, err := parser.ParseArgs(args)

	// --help and --version result in errors; turn them into successful output cmds
	if typedErr, ok := err.(*goflags.Error); ok {
		if typedErr.Type == goflags.ErrHelp {
			cmdOpts = &MessageOpts{Message: typedErr.Message}
			err = nil
		}
	}

	if _, ok := cmdOpts.(*HelpOpts); ok {
		cmdOpts = &MessageOpts{Message: helpText.String()}
	}

	return NewCmd(*commandOpts, cmdOpts, f.deps), err
}
