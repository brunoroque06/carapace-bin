package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_pushCmd = &cobra.Command{
	Use:   "push [OPTIONS] NAME[:TAG]",
	Short: "Upload an image to a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_pushCmd).Standalone()
	image_pushCmd.Flags().BoolP("all-tags", "a", false, "Push all tags of an image to the repository")
	image_pushCmd.Flags().Bool("disable-content-trust", true, "Skip image signing")
	image_pushCmd.Flags().BoolP("quiet", "q", false, "Suppress verbose output")
	imageCmd.AddCommand(image_pushCmd)

	rootAlias(image_pushCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalCompletion(
			docker.ActionRepositoryTags(),
		)
	})
}
