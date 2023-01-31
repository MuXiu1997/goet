package main

import (
	"fmt"
	"github.com/MuXiu1997/goet/pkg/file"
	"github.com/MuXiu1997/goet/pkg/templaterenderer"
	"os"

	"github.com/spf13/cobra"
)

const projectName = "goet"

//goland:noinspection GoUnusedGlobalVariable
var (
	version string
	commit  string
	date    string
	builtBy string
)

var (
	cmd = &cobra.Command{
		Use:   projectName,
		Short: projectName,
		Long:  "single-executable template renderer, powered by go template, sprig.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if printVersion {
				runPrintVersion()
				return nil
			}
			return run()
		},
	}

	template     string
	output       string
	printVersion bool
	options      Options
)

func main() {
	err := cmd.Execute()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initConfig() {
	cmd.Flags().StringVarP(&template, "template", "t", "-", "specify template file, \"-\" or unset means stdin")
	cmd.Flags().StringVarP(&output, "output", "o", "", "specify output file, unset means stdout")
	addValueOptionsFlags(cmd.Flags(), &options)
	cmd.Flags().BoolVarP(&printVersion, "version", "v", false, "print version")
	cmd.Flags().SortFlags = false
}

func runPrintVersion() {
	fmt.Printf("goet version %s, commit %s, built at %s", version, commit, date)
}

func run() error {
	templateContent, err := file.ReadFile(template)
	if err != nil {
		return err
	}
	values, err := options.MergeValues()
	if err != nil {
		return err
	}
	result, err := templaterenderer.Render(template, string(templateContent), values)
	if err != nil {
		return err
	}
	err = file.WriteFile(output, []byte(result))
	if err != nil {
		return err
	}
	return nil
}

func init() {
	initConfig()
}
