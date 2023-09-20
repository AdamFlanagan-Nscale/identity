/*
Copyright 2022-2023 EscherCloud.

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

package openstackplugincindercsi

import (
	"strings"

	unikornv1 "github.com/eschercloudai/unikorn/pkg/apis/unikorn/v1alpha1"
	"github.com/eschercloudai/unikorn/pkg/cd"
	"github.com/eschercloudai/unikorn/pkg/constants"
	"github.com/eschercloudai/unikorn/pkg/provisioners/application"
	"github.com/eschercloudai/unikorn/pkg/provisioners/helmapplications/openstackcloudprovider"
	"github.com/eschercloudai/unikorn/pkg/provisioners/util"

	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"sigs.k8s.io/yaml"
)

const (
	// applicationName is the unique name of the application.
	applicationName = "openstack-plugin-cinder-csi"
)

// Provisioner provides helm configuration interfaces.
type Provisioner struct {
	// driver is the CD driver that implments applications.
	driver cd.Driver

	// cluster is the Kubernetes cluster we're provisioning.
	cluster *unikornv1.KubernetesCluster
}

// New returns a new initialized provisioner object.
func New(driver cd.Driver, cluster *unikornv1.KubernetesCluster) *application.Provisioner {
	provisioner := &Provisioner{
		driver:  driver,
		cluster: cluster,
	}

	return application.New(driver, applicationName, cluster).WithGenerator(provisioner).InNamespace("ocp-system")
}

// Ensure the Provisioner interface is implemented.
var _ application.ValuesGenerator = &Provisioner{}

func (p *Provisioner) generateStorageClass(name string, reclaimPolicy corev1.PersistentVolumeReclaimPolicy, volumeBindingMode storagev1.VolumeBindingMode, isDefault, volumeExpansion bool) *storagev1.StorageClass {
	class := &storagev1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Provisioner: "cinder.csi.openstack.org",
		Parameters: map[string]string{
			"availability": *p.cluster.Spec.Openstack.VolumeFailureDomain,
		},
		ReclaimPolicy:        &reclaimPolicy,
		AllowVolumeExpansion: &volumeExpansion,
		VolumeBindingMode:    &volumeBindingMode,
	}

	if isDefault {
		class.Annotations = map[string]string{
			"storageclass.kubernetes.io/is-default-class": "true",
		}
	}

	return class
}

func (p *Provisioner) generateStorageClasses() []*storagev1.StorageClass {
	return []*storagev1.StorageClass{
		p.generateStorageClass("cinder", corev1.PersistentVolumeReclaimDelete, storagev1.VolumeBindingWaitForFirstConsumer, true, true),
	}
}

// Generate implements the application.ValuesGenerator interface.
func (p *Provisioner) Values(version *string) (interface{}, error) {
	storageClasses := p.generateStorageClasses()

	yamls := make([]string, len(storageClasses))

	for i, class := range storageClasses {
		var u unstructured.Unstructured

		// While we could hard code the TypeMeta, it's better practice and safer
		// to use the provided mechanisms.
		if err := p.driver.Client().Scheme().Convert(class, &u, nil); err != nil {
			return nil, err
		}

		y, err := yaml.Marshal(u.Object)
		if err != nil {
			return nil, err
		}

		yamls[i] = string(y)
	}

	cloudConfig, err := openstackcloudprovider.GenerateCloudConfig(p.cluster)
	if err != nil {
		return nil, err
	}

	cloudConfigHash, err := util.GetConfigurationHash(cloudConfig)
	if err != nil {
		return nil, err
	}

	values := map[string]interface{}{
		"commonAnnotations": map[string]interface{}{
			constants.ConfigurationHashAnnotation: cloudConfigHash,
		},
		// Allow scale to zero.
		"csi": map[string]interface{}{
			"plugin": map[string]interface{}{
				"controllerPlugin": map[string]interface{}{
					"nodeSelector": util.ControlPlaneNodeSelector(),
					"tolerations":  util.ControlPlaneTolerations(),
				},
			},
		},
		"storageClass": map[string]interface{}{
			"enabled": false,
			"custom":  strings.Join(yamls, "---\n"),
		},
	}

	return values, nil
}
