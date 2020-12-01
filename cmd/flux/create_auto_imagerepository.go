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
	"fmt"
	"time"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"

	imagev1 "github.com/fluxcd/image-reflector-controller/api/v1alpha1"
)

var createAutoImageRepositoryCmd = &cobra.Command{
	Use:   "imagerepository <name>",
	Short: "Create or update an ImageRepository object",
	Long: `The create auto imagerepository command generates an ImageRepository resource.
An ImageRepository object specifies an image repository to scan.`,
	RunE: createAutoImageRepositoryRun,
}

type imageRepoFlags struct {
	image     string
	secretRef string
	timeout   time.Duration
}

var imageRepoArgs = imageRepoFlags{}

func init() {
	flags := createAutoImageRepositoryCmd.Flags()
	flags.StringVar(&imageRepoArgs.image, "image", "", "the image repository to scan; e.g., library/alpine")
	flags.StringVar(&imageRepoArgs.secretRef, "secret-ref", "", "the name of a docker-registry secret to use for credentials")
	flags.DurationVar(&imageRepoArgs.timeout, "timeout", 0, "a timeout for scanning; this defaults to the interval if not set")

	createAutoCmd.AddCommand(createAutoImageRepositoryCmd)
}

func createAutoImageRepositoryRun(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("ImageRepository name is required")
	}
	objectName := args[0]

	if imageRepoArgs.image == "" {
		return fmt.Errorf("an image repository (--image) is required")
	}

	if _, err := name.NewRepository(imageRepoArgs.image); err != nil {
		return fmt.Errorf("unable to parse image value: %w", err)
	}

	labels, err := parseLabels()
	if err != nil {
		return err
	}

	var repo = imagev1.ImageRepository{
		ObjectMeta: metav1.ObjectMeta{
			Name:      objectName,
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: imagev1.ImageRepositorySpec{
			Image:    imageRepoArgs.image,
			Interval: metav1.Duration{Duration: interval},
		},
	}
	if imageRepoArgs.timeout != 0 {
		repo.Spec.Timeout = &metav1.Duration{Duration: imageRepoArgs.timeout}
	}
	if imageRepoArgs.secretRef != "" {
		repo.Spec.SecretRef = &corev1.LocalObjectReference{
			Name: imageRepoArgs.secretRef,
		}
	}

	return exportImageRepo(repo)
}

func exportImageRepo(repo imagev1.ImageRepository) error {
	gvk := imagev1.GroupVersion.WithKind(imagev1.ImageRepositoryKind)
	export := imagev1.ImageRepository{
		TypeMeta: metav1.TypeMeta{
			Kind:       gvk.Kind,
			APIVersion: gvk.GroupVersion().String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        repo.Name,
			Namespace:   repo.Namespace,
			Labels:      repo.Labels,
			Annotations: repo.Annotations,
		},
		Spec: repo.Spec,
	}

	data, err := yaml.Marshal(export)
	if err != nil {
		return err
	}

	fmt.Println("---")
	fmt.Println(resourceToString(data))
	return nil
}
