//go:build linux
// +build linux

/*
Copyright 2016 The Kubernetes Authors.

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

package conntrack

import (
	"github.com/vishvananda/netlink"
)

// FakeInterface implements Interface by just recording entries that have been cleared.
type FakeInterface struct {
	entries []*netlink.ConntrackFlow
}

var _ Interface = &FakeInterface{}

// NewFake creates a new FakeInterface
func NewFake() *FakeInterface {
	return &FakeInterface{entries: make([]*netlink.ConntrackFlow, 0)}
}

// ListEntries is part of Interface
func (fake *FakeInterface) ListEntries(_ uint8) ([]*netlink.ConntrackFlow, error) {
	return fake.entries, nil
}

// ClearEntries is part of Interface
func (fake *FakeInterface) ClearEntries(_ uint8, filters ...netlink.CustomConntrackFilter) (int, error) {
	var flows []*netlink.ConntrackFlow
	before := len(fake.entries)
	for _, flow := range fake.entries {
		var matched bool
		for _, filter := range filters {
			matched = filter.MatchConntrackFlow(flow)
			if matched {
				break
			}
		}
		if !matched {
			flows = append(flows, flow)
		}
	}
	fake.entries = flows
	return before - len(fake.entries), nil
}