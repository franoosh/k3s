//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by defaulter-gen. DO NOT EDIT.

package v1alpha1

import (
	rbacv1alpha1 "k8s.io/api/rbac/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&rbacv1alpha1.ClusterRoleBinding{}, func(obj interface{}) { SetObjectDefaults_ClusterRoleBinding(obj.(*rbacv1alpha1.ClusterRoleBinding)) })
	scheme.AddTypeDefaultingFunc(&rbacv1alpha1.ClusterRoleBindingList{}, func(obj interface{}) {
		SetObjectDefaults_ClusterRoleBindingList(obj.(*rbacv1alpha1.ClusterRoleBindingList))
	})
	scheme.AddTypeDefaultingFunc(&rbacv1alpha1.RoleBinding{}, func(obj interface{}) { SetObjectDefaults_RoleBinding(obj.(*rbacv1alpha1.RoleBinding)) })
	scheme.AddTypeDefaultingFunc(&rbacv1alpha1.RoleBindingList{}, func(obj interface{}) { SetObjectDefaults_RoleBindingList(obj.(*rbacv1alpha1.RoleBindingList)) })
	return nil
}

func SetObjectDefaults_ClusterRoleBinding(in *rbacv1alpha1.ClusterRoleBinding) {
	SetDefaults_ClusterRoleBinding(in)
	for i := range in.Subjects {
		a := &in.Subjects[i]
		SetDefaults_Subject(a)
	}
}

func SetObjectDefaults_ClusterRoleBindingList(in *rbacv1alpha1.ClusterRoleBindingList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ClusterRoleBinding(a)
	}
}

func SetObjectDefaults_RoleBinding(in *rbacv1alpha1.RoleBinding) {
	SetDefaults_RoleBinding(in)
	for i := range in.Subjects {
		a := &in.Subjects[i]
		SetDefaults_Subject(a)
	}
}

func SetObjectDefaults_RoleBindingList(in *rbacv1alpha1.RoleBindingList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_RoleBinding(a)
	}
}