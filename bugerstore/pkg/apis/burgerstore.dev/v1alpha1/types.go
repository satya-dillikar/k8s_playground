/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// These const variables are used in our custom controller.
const (
	GroupName string = "burgerstore.dev"
	Kind      string = "BurgerStore"
	Version   string = "v1alpha1"
	Plural    string = "burgerstores"
	Singluar  string = "burgerstore"
	ShortName string = "bgrs"
	Name      string = Plural + "." + GroupName
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BurgerStore is a specification for a BurgerStore resource
type BurgerStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec BurgerStoreSpec `json:"spec,omitempty"`
}

// BurgerStoreSpec is the spec for a BurgerStore resource
type BurgerStoreSpec struct {
	Owner      string `json:"owner,omitempty"`
	Address    string `json:"address,omitempty"`
	Currency   string `json:"currency,omitempty"`
	Investment int    `json:"investment,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BurgerStoreList is a list of BurgerStore resources
type BurgerStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []BurgerStore `json:"items,omitempty"`
}
