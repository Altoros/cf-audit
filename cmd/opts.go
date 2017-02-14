package cmd

type CommandOpts struct {
	// -----> Global options
	VersionOpt func() error `long:"version" short:"v" description:"Show CLI version"`
	// ConfigFileOpt string       `long:"config" short:"c" description:"Fetch variables" env:"CONFIG_FILE"`
	Help HelpOpts `command:"help" description:"Show this help message"`
	Diff DiffOpts `command:"diff" alias:"d" description:"show diff between several CF installatino"`

	JSONOpt           bool `long:"json"                      description:"Output as JSON"`
	NonInteractiveOpt bool `long:"non-interactive" short:"n" description:"Don't ask for user input"`
	TTYOpt            bool `long:"tty"                       description:"Force TTY-like output"`
	NoColorOpt        bool `long:"no-color"                  description:"Toggle colorized output"`
}

type DiffOpts struct {
	// Args []string `positional-args:"true" required:"false"`
	CloudFoundriesNamesFlags
	VaultFlags
	ConfigFileFlags
	cmd
}

type ConfigFileFlags struct {
	ConfigFileOpt string `long:"config" short:"c" description:"Fetch variables" env:"CONFIG_FILE"`
}

type CloudFoundriesNamesFlags struct {
	CloudFoundriesNamesOpt []string `long:"cloudfoundry" description:"Used with Vault, a comma separated list of foundation names (required, if you use Vault to load foundation creds)" env:"CLOUDFOUNDRIES"`
}

type VaultFlags struct {
	VaultAddrOpt  string `long:"vault-addr" description:"Vault address (optional)" env:"VAULT_ADDR" optional:"true"`
	VaultTokenOpt string `long:"vault-token" description:"Vault token (optional)" env:"VAULT_TOKEN" optional:"true"`
}

// MessageOpts is used for version and help flags
type MessageOpts struct {
	Message string
}

type HelpOpts struct {
	cmd
}

type cmd struct{}

func (c cmd) Execute(_ []string) error { panic("Unreachable") }

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
