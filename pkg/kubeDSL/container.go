// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kubeDSL

// NewContainer creates a new container builder
func NewContainer() ContainerBuilder {
	return &Container{}
}

// ContainerBuilder container builder interface
type ContainerBuilder interface {
	SetName(string) ContainerBuilder
	SetImage(string) ContainerBuilder
	SetArgs(...string) ContainerBuilder
	SetCommand(...string) ContainerBuilder
	Build() Container
}

// Container container type defines an abstraction for containers
type Container struct {
	name    string
	image   string
	command []string
	args    []string
}

// Command returns a container command
func (c *Container) Command() []string {
	return c.command
}

// SetCommand sets a container command
func (c *Container) SetCommand(command ...string) ContainerBuilder {
	c.command = command
	return c
}

// Args returns a container args
func (c *Container) Args() []string {
	return c.args
}

// SetArgs sets a container arguments
func (c *Container) SetArgs(args ...string) ContainerBuilder {
	c.args = args
	return c
}

// SetName sets a container name
func (c *Container) SetName(name string) ContainerBuilder {
	c.name = name
	return c
}

// Name returns the name of a container
func (c *Container) Name() string {
	return c.name
}

// SetImage sets a container image
func (c *Container) SetImage(image string) ContainerBuilder {
	c.image = image
	return c
}

// Image returns the name of a container image
func (c *Container) Image() string {
	return c.image
}

// Build builds a container
func (c *Container) Build() Container {
	return Container{
		name:    c.name,
		image:   c.image,
		command: c.command,
		args:    c.args,
	}
}
