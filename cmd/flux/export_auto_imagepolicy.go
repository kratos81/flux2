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

var exportImagePolicyCmd = &cobra.Command{
	Use:   "image-policy [name]",
	Short: "Export ImagePolicy resources in YAML format",
	Long:  "The export image-policy command exports one or all ImagePolicy resources in YAML format.",
	Example: `  # Export all ImagePolicy resources
  flux export auto image-policy --all > image-policies.yaml

  # Export a Provider
  flux export auto image-policy alpine1x > alpine1x.yaml
`,
	RunE: exportCommand{
		object: &imagev1.ImagePolicy{},
		list:   &imagev1.ImagePolicyList{},
	}.run,
}

func init() {
	exportAutoCmd.AddCommand(exportImagePolicyCmd)
}
