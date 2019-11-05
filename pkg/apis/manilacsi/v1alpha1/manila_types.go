package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ManilaSpec defines the desired state of Manila
// +k8s:openapi-gen=true
type ManilaSpec struct {
	Config       ManilaCSIConfig   `json:"config"`
	Image        string            `json:"image",omitempty`
	NodeSelector map[string]string `json:"nodeSelector",omitempty`
	Tolerations  []v1.Toleration   `json:"tolerations",omitempty`
	Topologies   []Topologies      `json:"topologies",omitempty`
}

type Topologies struct {
	// Key-value pairs corresponding to the NodeName
	Topology map[string]string `json:"topology,omitempty"`
	// Node Hostname with its allowed topology
	Nodes []v1.NodeSelectorRequirement `json:"nodes,omitempty"`
}

type ManilaCSIConfig struct {
	EnvVars  EnvVars `json:"envVars"`
}

type EnvVars struct {

}


// ManilaStatus defines the observed state of Manila
// +k8s:openapi-gen=true
type ManilaStatus struct {
	// Node Hostname with its allowed topology
	Nodes []v1.NodeSelectorRequirement `json:"nodes,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Manila is the Schema for the manilas API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=manilas,scope=Namespaced
type Manila struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManilaSpec   `json:"spec,omitempty"`
	Status ManilaStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ManilaList contains a list of Manila
type ManilaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Manila `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Manila{}, &ManilaList{})
}
