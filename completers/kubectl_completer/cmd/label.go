package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var labelCmd = &cobra.Command{
	Use:   "label",
	Short: "Update the labels on a resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(labelCmd).Standalone()
	labelCmd.Flags().Bool("all", false, "Select all resources, in the namespace of the specified resource types")
	labelCmd.Flags().BoolP("all-namespaces", "A", false, "If true, check the specified action in all namespaces.")
	labelCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	labelCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	labelCmd.Flags().String("field-manager", "kubectl-label", "Name of the manager used to track field ownership.")
	labelCmd.Flags().String("field-selector", "", "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.")
	labelCmd.Flags().StringSliceP("filename", "f", []string{}, "Filename, directory, or URL to files identifying the resource to update the labels")
	labelCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	labelCmd.Flags().Bool("list", false, "If true, display the labels for a given resource.")
	labelCmd.Flags().Bool("local", false, "If true, label will NOT contact api-server but run locally.")
	labelCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	labelCmd.Flags().Bool("overwrite", false, "If true, allow labels to be overwritten, otherwise reject label updates that overwrite existing labels.")
	labelCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	labelCmd.Flags().String("resource-version", "", "If non-empty, the labels update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource.")
	labelCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.")
	labelCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	labelCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	labelCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	rootCmd.AddCommand(labelCmd)

	carapace.Gen(labelCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   action.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(labelCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if labelCmd.Flag("filename").Changed {
				return carapace.ActionValues() // TODO get labels for file
			} else {
				return action.ActionApiResources()
			}
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if labelCmd.Flag("all").Changed {
				return carapace.ActionValues()
			} else {
				return action.ActionResources("", c.Args[0])
			}
		}),
	)

	carapace.Gen(labelCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if labelCmd.Flag("filename").Changed || labelCmd.Flag("all").Changed {
				return carapace.ActionValues() // TODO support labels for file
			} else {
				return action.ActionLabels("", c.Args[0]+"/"+c.Args[1])
			}
		}),
	)
}
