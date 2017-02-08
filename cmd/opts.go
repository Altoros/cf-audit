package cmd

type CommandOpts struct {
	// -----> Global options
	Version func() error `long:"version" short:"v" description:"Show Service Brokere version"`
	Config  string       `long:"config" short:"c" description:"Fetch variables" env:"CONFIG_FILE"`
	Help    bool         `long:"help" short:"h" description:"Show this help message"`
}

// command example
// cf-audit diff -c foundations.yml

// command example output
// Org:Mobile
//   Space:Dev
//     App:core-app
//       Quota[foundation1]:10240
//       Quota[foundation2]:10240

// command example output
// Org:Mobile
//   Space:Dev
//     App:core-app
//       Quota[foundation1]:10240
//       Quota[foundation2]:10240

type CommandOpts struct {
	// -----> Global options
	Version func() error `long:"version" short:"v" description:"Show Service Brokere version"`
	Config  string       `long:"config" short:"c" description:"Fetch variables" env:"CONFIG_FILE"`
	Help    bool         `long:"help" short:"h" description:"Show this help message"`
}
