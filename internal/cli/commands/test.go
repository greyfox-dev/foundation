package commands

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	f "github.com/uplatform-ai/foundation"
	h "github.com/uplatform-ai/foundation/internal/cli/helpers"
)

var Test = &cobra.Command{
	Use:     "test",
	Aliases: []string{"t"},
	Short:   "Run tests",
	Run: func(cmd *cobra.Command, _ []string) {
		if !h.BuiltOnFoundation() {
			log.Fatal("This command must be run from inside a Foundation service")
		}

		if f.IsProductionEnv() {
			log.Fatal("You're trying to run tests in production environment")
		}

		env := os.Environ()
		env = append(env, "FOUNDATION_ENV=test")

		opts := []string{
			"test",
		}
		if cmd.Flag("opts") != nil {
			opts = append(opts, strings.Split(cmd.Flag("opts").Value.String(), " ")...)
		}
		opts = append(opts, h.AtServiceRoot("..."))

		test := exec.Command("go", opts...)
		test.Stdout = os.Stdout
		test.Stderr = os.Stderr
		test.Env = env
		if err := test.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	Test.Flags().StringP("opts", "o", "-cover", "Options to pass to go test")
}
