//  Copyright 2020 Authors of Cilium
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:singular="ciliumexternalworkload",path="ciliumexternalworkloads",scope="Namespaced",shortName={cew}
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type=date

// CiliumExternalWorkload is a Kubernetes Custom Resource that contains a
// specification for an external workload that can join the cluster.  The
// name of the CRD is the FQDN of the external workload, and it needs to
// match the name in the workload registration. The labels on the CRD
// object are the (kubernetes) labels that will, together with the CRD
// namespace, be used to allocate a Cilium Identity for the external
// workload.
type CiliumExternalWorkload struct {
	// +k8s:openapi-gen=false
	// +deepequal-gen=false
	metav1.TypeMeta `json:",inline"`
	// +k8s:openapi-gen=false
	// +deepequal-gen=false
	metav1.ObjectMeta `json:"metadata"`

	// Spec is the desired configuration of the external Cilium workload.
	Spec CiliumExternalWorkloadSpec `json:"spec,omitempty"`

	// Status is the most recent status of the external Cilium workload.
	// It is a read-only field.
	//
	// +deepequal-gen=false
	// +kubebuilder:validation:Optional
	Status CiliumExternalWorkloadStatus `json:"status"`
}

// CiliumExternalWorkloadSpec specifies the configurations for redirecting traffic
// within a workload.
//
// +kubebuilder:validation:Type=object
type CiliumExternalWorkloadSpec struct {
	// TBD, stuff that reduces the burden on Cilium Agent configuration on the external workloads.
}

// CiliumExternalWorkloadStatus is the status of a the external Cilium workload.
type CiliumExternalWorkloadStatus struct {
	// IP is the IP address of the workload. The IP must be reachable between
	// nodes and workloads. Empty if the workload has not registered.
	IP string `json:"ip,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=false
// +deepequal-gen=false

// CiliumExternalWorkloadList is a list of CiliumExternalWorkload objects.
type CiliumExternalWorkloadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// Items is a list of CiliumExternalWorkload
	Items []CiliumExternalWorkload `json:"items"`
}
