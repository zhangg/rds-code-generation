package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Mysql is a specification for a Mysql resource
type Mysql struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MysqlSpec   `json:"spec"`
	Status MysqlStatus `json:"status"`
}

// MysqlSpec is the spec for a Mysql resource
type MysqlSpec struct {
	Port *int32 `json:"port"`
	Replicas       *int32 `json:"replicas"`
}

// MysqlStatus is the status for a Mysql resource
type MysqlStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MysqlList is a list of Mysql resources
type MysqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Mysql `json:"items"`
}
