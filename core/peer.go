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
)

type VersionedValue struct {
	value   string
	version uint32
}

type Peer struct {
	Addr       net.TCPAddr
	Generation uint64
	Heartbeat  uint32

	valueGenerator uint32

	// app state varibles
	Datacenter *VersionedValue
	Rack       *VersionedValue
	// add more states here, if you care ....

	// state to represent if this peer is GC'ing or blocked or just unresponsive
	Health PeerHealth
	chan 
}

type PeerHealth struct {
	state PeerState
}

type PeerState int

const (
	PeerHealthy PeerState = iota
	PeerBlocked
	PeerDown
)

func (peer *Peer) NextRound() {
	peer.determineHealth()
}

// statistically determine if this peer should be in a blocked state, like GC or process blocked,
// or update an existing blocked state
func (peer *Peer) determineHealth() {
	if peer.Health.state == PeerHealthy {
		// determine if we should go into an unhealthy state
	} else {
		// determine if we should go back to healthy
	}
}
