/*
Copyright 2020 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/spf13/cobra"

	imagev1 "github.com/fluxcd/image-reflector-controller/api/v1alpha1"
)

var exportImageRepositoryCmd = &cobra.Command{
	Use:   "image-repository [name]",
	Short: "Export ImageRepository resources in YAML format",
	Long:  "The export image-repository command exports one or all ImageRepository resources in YAML format.",
	Example: `  # Export all ImageRepository resources
  flux export auto image-repository --all > image-repositories.yaml

  # Export a Provider
  flux export auto image-repository alpine > alpine.yaml
`,
	RunE: exportCommand{
		object: &imagev1.ImageRepository{},
		list:   &imagev1.ImageRepositoryList{},
	}.run,
}

func init() {
	exportAutoCmd.AddCommand(exportImageRepositoryCmd)
}
