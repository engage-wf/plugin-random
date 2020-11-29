package main

import (
	"fmt"
	"strconv"

	random "github.com/engage-wf/plugin-random"

	"github.com/engage-wf/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	envPrefix = "ENGAGE_RANDOM"
)

var (
	version = "none"
	commit  = "none"
	app     = &cobra.Command{
		Use:   "random",
		Short: "Generate random stuff",
	}
	cmdHexadecimal = &cobra.Command{
		Use:     "hexadecimal",
		Aliases: []string{"h", "hex"},
		Short:   "Generate uppercase Hex strings",
		Args:    cobra.MinimumNArgs(1),
		Run:     executeHexadecimal,
	}
	cmdString = &cobra.Command{
		Use:     "string",
		Aliases: []string{"s", "str", "alnum", "alphanum"},
		Short:   "Generate Alphanumerical strings",
		Args:    cobra.MinimumNArgs(1),
		Run:     executeString,
	}
	cmdPrintable = &cobra.Command{
		Use:     "printable",
		Aliases: []string{"p", "prn", "print"},
		Short:   "Generate printable strings",
		Args:    cobra.MinimumNArgs(1),
		Run:     executePrintable,
	}
	cmdDigits = &cobra.Command{
		Use:     "digits",
		Aliases: []string{"d", "dig", "digit"},
		Short:   "Generate Digit strings",
		Args:    cobra.MinimumNArgs(1),
		Run:     executeDigits,
	}
	cmdURLSafe = &cobra.Command{
		Use:     "url-safe",
		Aliases: []string{"u", "url"},
		Short:   "Generate Digit strings",
		Args:    cobra.MinimumNArgs(1),
		Run:     executeURLSafe,
	}
	cmdUUID = &cobra.Command{
		Use:     "uuid",
		Aliases: []string{"u"},
		Short:   "Generate UUIDs",
		Args:    cobra.MinimumNArgs(1),
		Run:     executeUUID,
	}
)

func init() {
	core.DefaultCLI(app, version, commit, envPrefix)
	app.PersistentFlags().BoolP("secure", "s", false, "Use secure random source for generation")
	viper.BindPFlag("secure", app.PersistentFlags().Lookup("secure"))
	app.PersistentFlags().BoolP("text", "t", false, "Whether to output Text (instead of the default JSON)")
	viper.BindPFlag("text", app.PersistentFlags().Lookup("text"))
	app.AddCommand(cmdURLSafe, cmdDigits, cmdHexadecimal, cmdString, cmdPrintable)
}

func main() {
	if err := app.Execute(); err != nil {
		panic(err)
	}
}

func createGenerator() random.RNG {
	if viper.GetBool("secure") {
		fmt.Println("Using secure RNG")
		return random.NewSecureRNG()
	} else {
		return random.NewDefaultRNG()
	}
}

func executeHexadecimal(cmd *cobra.Command, args []string) {
	generateForArgs(random.Hex(), args)
}

func executeString(cmd *cobra.Command, args []string) {
	generateForArgs(random.String(), args)
}

func executePrintable(cmd *cobra.Command, args []string) {
	generateForArgs(random.Printable(), args)
}

func executeDigits(cmd *cobra.Command, args []string) {
	generateForArgs(random.Digits(), args)
}

func executeURLSafe(cmd *cobra.Command, args []string) {
	generateForArgs(random.URLSafe(), args)
}

func executeUUID(cmd *cobra.Command, args []string) {
	fmt.Println(random.GenUUID())
}

func generateForArgs(alphabet random.Alphabet, args []string) {
	gen := createGenerator()
	var result []string
	for _, arg := range args {
		length, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		result = append(result, alphabet.RGen(gen, length))
	}
	if viper.GetBool("text") {
		for _, s := range result {
			fmt.Println(s)
		}
	} else {
		core.PrintJSON(result)
	}
}
