/*
Copyright 2020 The Scribe authors.

This file may be used, at your option, according to either the GNU AGPL 3.0 or
the Apache V2 license.

---
This program is free software: you can redistribute it and/or modify it under
the terms of the GNU Affero General Public License as published by the Free
Software Foundation, either version 3 of the License, or (at your option) any
later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.  See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along
with this program.  If not, see <https://www.gnu.org/licenses/>.

---
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

//+kubebuilder:validation:Required
package v1alpha1

import (
	"github.com/operator-framework/operator-lib/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ReplicationDestinationSpec defines the desired state of ReplicationDestination
type ReplicationDestinationSpec struct {
	// replicationMethod determines the method used to replicate the volume.
	// This may be either a method built into the Scribe controller or the name
	// of an external plugin. It must match the replicationMethod of the
	// corresponding source.
	ReplicationMethod ReplicationMethodType `json:"replicationMethod,omitempty"`
	// parameters are method-specific key/value configuration parameters. For
	// more information, please see the documentation of the specific
	// replicationMethod being used.
	Parameters map[string]string `json:"parameters,omitempty"`
}

type ReplicationDestinationRsyncStatus struct {
	// sshKeys is the name of a Secret that contains the SSH keys to be used for
	// authentication. If not provided in .spec.rsync.sshKeys, SSH keys will be
	// generated and the appropriate keys for the remote side will be placed
	// here.
	//+optional
	SSHKeys *string `json:"sshKeys,omitempty"`
	// address is the address to connect to for incoming SSH replication
	// connections.
	//+optional
	Address *string `json:"address,omitempty"`
	// port is the SSH port to connect to for incoming SSH replication
	// connections.
	//+optional
	Port *int32 `json:"port,omitempty"`
}

// ReplicationDestinationStatus defines the observed state of ReplicationDestination
type ReplicationDestinationStatus struct {
	// lastSyncTime is the time of the most recent successful synchronization.
	//+optional
	LastSyncTime *metav1.Time `json:"lastSyncTime,omitempty"`
	// lastSyncDuration is the amount of time required to send the most recent
	// update.
	//+optional
	LastSyncDuration *metav1.Duration `json:"lastSyncDuration,omitempty"`
	// rsync contains status information for Rsync-based replication.
	Rsync *ReplicationDestinationRsyncStatus `json:"rsync,omitempty"`
	// external contains provider-specific status information. For more details,
	// please see the documentation of the specific replication provider being
	// used.
	//+optional
	External map[string]string `json:"external,omitempty"`
	// conditions represent the latest available observations of the
	// destination's state.
	Conditions status.Conditions `json:"conditions,omitempty"`
}

// ReplicationDestination defines the destination for a replicated volume
//+kubebuilder:object:root=true
//+kubebuilder:resource:scope=Namespaced
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Last sync",type="string",format="date-time",JSONPath=`.status.lastSyncTime`
//+kubebuilder:printcolumn:name="Duration",type="string",JSONPath=`.status.lastSyncDuration`
type ReplicationDestination struct {
	metav1.TypeMeta `json:",inline"`
	//+optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// spec is the desired state of the ReplicationDestination, including the
	// replication method to use and its configuration.
	Spec ReplicationDestinationSpec `json:"spec,omitempty"`
	// status is the observed state of the ReplicationDestination as determined
	// by the controller.
	//+optional
	Status *ReplicationDestinationStatus `json:"status,omitempty"`
}

// ReplicationDestinationList contains a list of ReplicationDestination
//+kubebuilder:object:root=true
type ReplicationDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicationDestination `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ReplicationDestination{}, &ReplicationDestinationList{})
}
