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

package provider

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ContextOption func(*Context)

// Context represents a context for the providers
type Context struct {
	context.Context

	command *cobra.Command
	logger  *logrus.Entry

	interactive bool

	//ClusterProvider ClusterProvider
}

// NewContext creates a new context
func NewContext(cmd *cobra.Command, opts ...ContextOption) *Context {
	defaultContext := context.Background()

	c := &Context{
		Context:     defaultContext,
		command:     cmd,
		logger:      logrus.StandardLogger().WithContext(defaultContext),
		interactive: true,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithContext(ctx context.Context) ContextOption {
	return func(c *Context) {
		c.Context = ctx
	}
}

// func WithClusterProvider(provider ClusterProvider) ContextOption {
// 	return func(c *Context) {
// 		c.ClusterProvider = provider
// 	}
// }

func WithLogger(logger *logrus.Entry) ContextOption {
	return func(c *Context) {
		c.logger = logger
	}
}

func WithInteractive(interactive bool) ContextOption {
	return func(c *Context) {
		c.interactive = interactive
	}
}

func (c *Context) Command() *cobra.Command {
	return c.command
}

func (c *Context) Logger() *logrus.Entry {
	return c.logger
}

func (c *Context) IsInteractive() bool {
	return c.interactive
}
