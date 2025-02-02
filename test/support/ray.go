/*
Copyright 2023.

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

package support

import (
	"github.com/onsi/gomega"
	rayv1alpha1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const RayJobDefaultClusterSelectorKey = "ray.io/cluster"

func RayJob(t Test, namespace, name string) func(g gomega.Gomega) *rayv1alpha1.RayJob {
	return func(g gomega.Gomega) *rayv1alpha1.RayJob {
		job, err := t.Client().Ray().RayV1alpha1().RayJobs(namespace).Get(t.Ctx(), name, metav1.GetOptions{})
		g.Expect(err).NotTo(gomega.HaveOccurred())
		return job
	}
}

func GetRayJob(t Test, namespace, name string) *rayv1alpha1.RayJob {
	t.T().Helper()
	return RayJob(t, namespace, name)(t)
}

func RayJobStatus(job *rayv1alpha1.RayJob) rayv1alpha1.JobStatus {
	return job.Status.JobStatus
}

func GetRayJobId(t Test, namespace, name string) string {
	t.T().Helper()
	job := RayJob(t, namespace, name)(t)
	return job.Status.JobId
}

func RayCluster(t Test, namespace, name string) func(g gomega.Gomega) *rayv1alpha1.RayCluster {
	return func(g gomega.Gomega) *rayv1alpha1.RayCluster {
		cluster, err := t.Client().Ray().RayV1alpha1().RayClusters(namespace).Get(t.Ctx(), name, metav1.GetOptions{})
		g.Expect(err).NotTo(gomega.HaveOccurred())
		return cluster
	}
}

func GetRayCluster(t Test, namespace, name string) *rayv1alpha1.RayCluster {
	t.T().Helper()
	return RayCluster(t, namespace, name)(t)
}

func RayClusterState(cluster *rayv1alpha1.RayCluster) rayv1alpha1.ClusterState {
	return cluster.Status.State
}

func WriteRayJobLogs(t Test, rayClient RayClusterClient, namespace, name string) {
	WriteRayJobAPILogs(t, rayClient, GetRayJobId(t, namespace, name))
}
