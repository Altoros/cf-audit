package cmd

type CommandOpts struct {
	// -----> Global options
	VersionOpt    func() error `long:"version" short:"v" description:"Show CLI version"`
	ConfigFileOpt string       `long:"config" short:"c" description:"Fetch variables" env:"CONFIG_FILE"`
	Help          HelpOpts     `command:"help" description:"Show this help message"`

	// Instance management
	Diff DiffOpts `command:"diff"     description:"show diff between several CF installatino"`

	JSONOpt           bool `long:"json"                      description:"Output as JSON"`
	NonInteractiveOpt bool `long:"non-interactive" short:"n" description:"Don't ask for user input"`
	TTYOpt            bool `long:"tty"                       description:"Force TTY-like output"`
	NoColorOpt        bool `long:"no-color"                  description:"Toggle colorized output"`
}

type DiffOpts struct {
	Args []string `positional-args:"true" required:"false"`
	cmd
}

// MessageOpts is used for version and help flags
type MessageOpts struct {
	Message string
}

type HelpOpts struct {
	cmd
}

type cmd struct{}

// command example
// cf-audit diff -c foundations.yml

// command example output
// Org:Mobile
//   Space:Dev
//     App:core-app
//       Quota[foundation1]:10240
//       Quota[foundation2]:10240

// Org:Mobile
//   Space:Dev
//     App:core-app
//       Instances[founation2]:2
//       Instances[founation2]:3

// Org:Mobile
//   Space:Dev
//     App:core-app
//       Quata[founation2]:2
//       Instances[founation2]:3
