/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package core

import (
	"net"
	//	"github.com/jasobrown/ebsim/node"
)

type VersionedValue struct {
	value   string
	version uint32
}

type Node struct {
	Addr       net.TCPAddr
	generation uint64
	heartbeat  uint32

	valueGenerator uint32

	// app state varibles
	datacenter *VersionedValue
	rack       *VersionedValue
	// add more states here, if you care ....

	// state to represent if this node is GC'ing or blocked or just unresponsive
	Health NodeHealth
}

type NodeHealth struct {
	state NodeState
}

type NodeState int

const (
	NodeHealthy NodeState = iota
	NodeBlocked
	NodeDown
)

func (node *Node) NextRound() {
	node.determineHealth()
}

// statistically determine if this node should be in a blocked state, like GC or process blocked,
// or update an existing blocked state
func (node *Node) determineHealth() {
	if node.Health.state == NodeHealthy {
		// determine if we should go into an unhealthy state
	} else {
		// determine if we should go back to healthy
	}
}
