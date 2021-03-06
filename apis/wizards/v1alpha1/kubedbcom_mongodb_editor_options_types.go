/*
Copyright AppsCode Inc. and Contributors

Licensed under the PolyForm Noncommercial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/PolyForm-Noncommercial-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	api "kubepack.dev/lib-app/api/v1alpha1"
)

// KubedbcomMongodbEditorOptions defines the schama for MongoDB Editor UI Options.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=KubedbcomMongodbEditorOptionss,singular=KubedbcomMongodbEditorOptions
type KubedbcomMongodbEditorOptions struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubedbcomMongodbEditorOptionsSpec `json:"spec,omitempty"`
}

// KubedbcomMongodbEditorOptionsSpec is the schema for MongoDB profile values file
type KubedbcomMongodbEditorOptionsSpec struct {
	api.Metadata `json:"metadata,omitempty"`
	Spec         KubedbcomMongodbEditorOptionsSpecSpec `json:"spec"`
}

type KubedbcomMongodbEditorOptionsSpecSpec struct {
	Version           string                    `json:"version"`
	Mode              MongodbMode               `json:"mode"`
	Replicas          int                       `json:"replicas"`
	ReplicaSet        ReplicaSet                `json:"replicaSet"`
	ShardTopology     ShardTopology             `json:"shardTopology"`
	ClusterAuthMode   ClusterAuthMode           `json:"clusterAuthMode"`
	SslMode           SSLMode                   `json:"sslMode"`
	TerminationPolicy TerminationPolicy         `json:"terminationPolicy"`
	StorageClass      StorageClass              `json:"storageClass"`
	Persistence       Persistence               `json:"persistence"`
	Machine           MachineType               `json:"machine"`
	Resources         core.ResourceRequirements `json:"resources"`
	AuthSecret        AuthSecret                `json:"authSecret"`
	ServiceAccount    ServiceAccount            `json:"serviceAccount"`
}

// +kubebuilder:validation:Enum=db.t.micro;db.t.small;db.t.medium;db.t.large;db.t.xlarge;db.t.2xlarge;db.m.small;db.m.large;db.m.xlarge;db.m.2xlarge;db.m.4xlarge;db.m.8xlarge;db.m.12xlarge;db.m.16xlarge;db.m.24xlarge;db.r.large;db.r.xlarge;db.r.2xlarge;db.r.4xlarge;db.r.8xlarge;db.r.12xlarge;db.r.16xlarge;db.r.24xlarge
type MachineType string

// +kubebuilder:validation:Enum=Halt;Delete;WipeOut;DoNotTerminate
type TerminationPolicy string

// +kubebuilder:validation:Enum=Standalone;Replicaset;Sharded
type MongodbMode string

// +kubebuilder:validation:Enum=keyFile;sendKeyFile;sendX509;x509
type ClusterAuthMode string

// +kubebuilder:validation:Enum=disabled;allowSSL;preferSSL;requireSSL
type SSLMode string

type ReplicaSet struct {
	Name string `json:"name"`
}

type Persistence struct {
	Size string `json:"size"`
}

type Shard struct {
	Replicas    int         `json:"replicas"`
	Shards      int         `json:"shards"`
	Persistence Persistence `json:"persistence"`
}

type ConfigServer struct {
	Replicas    int         `json:"replicas"`
	Persistence Persistence `json:"persistence"`
}

type Mongos struct {
	Replicas int `json:"replicas"`
}

type ShardTopology struct {
	Shard        Shard        `json:"shard"`
	ConfigServer ConfigServer `json:"configServer"`
	Mongos       Mongos       `json:"mongos"`
}

type StorageClass struct {
	Name string `json:"name"`
}

type AuthSecret struct {
	Create      bool              `json:"create"`
	Annotations map[string]string `json:"annotations"`
	Name        string            `json:"name"`
	Password    string            `json:"password"`
}

type ServiceAccount struct {
	Create      bool              `json:"create"`
	Annotations map[string]string `json:"annotations"`
	Name        string            `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubedbcomMongodbEditorOptionsList is a list of KubedbcomMongodbEditorOptionss
type KubedbcomMongodbEditorOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubedbcomMongodbEditorOptions CRD objects
	Items []KubedbcomMongodbEditorOptions `json:"items,omitempty"`
}
