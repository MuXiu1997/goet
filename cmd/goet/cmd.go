package main

import (
	"fmt"
	"github.com/MuXiu1997/goet/pkg/file"
	"github.com/MuXiu1997/goet/pkg/templaterenderer"
	"github.com/spf13/cobra"
)

var (
	template string
	output   string
	options  Options
)

var (
	cmd *cobra.Command
)

func initCmd() {
	cmd = &cobra.Command{
		Use:   projectName,
		Short: projectName,
		Long:  "single-executable template renderer, powered by go template, sprig.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		Version:      version,
		SilenceUsage: true,
	}
	cmd.SetVersionTemplate(fmt.Sprintf("goet version %s, commit %s, built at %s", version, commit, date))
}

func initConfig() {
	cmd.Flags().StringVarP(&template, "template", "t", "-", "specify template file, \"-\" or unset means stdin")
	cmd.Flags().StringVarP(&output, "output", "o", "", "specify output file, unset means stdout")
	addValueOptionsFlags(cmd.Flags(), &options)
	cmd.Flags().SortFlags = false
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
