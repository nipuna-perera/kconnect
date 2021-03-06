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

package aws

import (
	"fmt"

	"github.com/spf13/pflag"

	"github.com/fidelity/kconnect/pkg/aws"
	"github.com/fidelity/kconnect/pkg/flags"
	"github.com/fidelity/kconnect/pkg/provider"
)

// Resolve will resolve the values for the AWS specific flags that have no value. It will
// query AWS and interactively ask the user for selections.
func (p *ServiceProvider) ResolveFlags(ctx *provider.Context, flagset *pflag.FlagSet) error {
	p.logger.Debug("resolving AWS identity flags")

	// TODO: handle in declarative mannor
	// NOTE: resolution is only needed for required fields
	if err := p.resolveIdpProvider("idp-provider", flagset); err != nil {
		return fmt.Errorf("resolving idp-provider: %w", err)
	}
	if err := p.resolveIdpEndpoint("idp-endpoint", flagset); err != nil {
		return fmt.Errorf("resolving idp-endpoint: %w", err)
	}
	if err := p.resolveProfile("profile", flagset); err != nil {
		return fmt.Errorf("resolving profile: %w", err)
	}
	if err := p.resolveRegion("region", flagset); err != nil {
		return fmt.Errorf("resolving region: %w", err)
	}
	if err := p.resolveUsername("username", flagset); err != nil {
		return fmt.Errorf("resolving username: %w", err)
	}
	if err := p.resolvePassword("password", flagset); err != nil {
		return fmt.Errorf("resolving password: %w", err)
	}

	// if err := r.resolveRoleARN("role-arn", flagset); err != nil {
	// 	return fmt.Errorf("resolving role-arn: %w", err)
	// }

	return nil
}

func (p *ServiceProvider) ensureSessionFromFlags(flagset *pflag.FlagSet) error {
	if p.session != nil {
		return nil
	}

	if !flags.ExistsWithValue("region", flagset) {
		return ErrNoRegion

	}
	if !flags.ExistsWithValue("profile", flagset) {
		return ErrNoProfile
	}

	region, err := flagset.GetString("region")
	if err != nil {
		return fmt.Errorf("getting region flag: %w", err)
	}
	profile, err := flagset.GetString("profile")
	if err != nil {
		return fmt.Errorf("getting profile flag: %w", err)
	}

	session, err := aws.NewSession(region, profile)
	if err != nil {
		return fmt.Errorf("creating new AWS session: %w", err)
	}

	p.session = session

	return nil
}
