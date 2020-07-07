/*
Copyright 2020 The kconnect Authors.

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

package resolver

import "github.com/spf13/pflag"

// FlagsResolver is used to resolve the values for flags interactively.
// There will be a flags resolver for Azure, AWS and Rancher initially.
type FlagsResolver interface {

	// Resolve will resolve the values for the supplied flags. It would interactively
	// resolve the values by asking the user for selections. It will basically set the
	// Value on each pflag.Flag
	Resolve(identity Identity, flags []*pflag.Flag) error
}

type AzureFlagsResolver struct{}
type AWSFlagsResolver struct{}
type RancherFlagsResolver struct{}