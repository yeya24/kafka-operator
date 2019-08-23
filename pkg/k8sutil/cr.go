// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8sutil

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"emperror.dev/emperror"
	banzaicloudv1alpha1 "github.com/banzaicloud/kafka-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	runtimeClient "sigs.k8s.io/controller-runtime/pkg/client"
)

func updateCrWithNodeAffinity(current *corev1.Pod, cr *banzaicloudv1alpha1.KafkaCluster, client runtimeClient.Client) error {
	failureDomainSelectors, err := failureDomainSelectors(current.Spec.NodeName, client)
	if err != nil {
		return emperror.WrapWith(err, "determining Node selector failed")
	}

	// don't set node affinity when none of the selector labels are available for the node
	if len(failureDomainSelectors) < 1 {
		return nil
	}

	brokerConfigs := []banzaicloudv1alpha1.BrokerConfig{}

	for _, brokerConfig := range cr.Spec.BrokerConfigs {
		if strconv.Itoa(int(brokerConfig.Id)) == current.Labels["brokerId"] {
			nodeAffinity := &corev1.NodeAffinity{
				RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
					NodeSelectorTerms: failureDomainSelectors,
				},
			}
			brokerConfig.NodeAffinity = nodeAffinity
		}
		brokerConfigs = append(brokerConfigs, brokerConfig)
	}
	cr.Spec.BrokerConfigs = brokerConfigs
	return updateCr(cr, client)
}

func updateCrWithRackAwarenessConfig(pod *corev1.Pod, cr *banzaicloudv1alpha1.KafkaCluster, client runtimeClient.Client) error {

	rackConfigMap, err := getSpecificNodeLabels(pod.Spec.NodeName, client, cr.Spec.RackAwareness.Labels)
	if err != nil {
		return emperror.WrapWith(err, "fetching Node rack awareness labels failed")
	}
	rackConfigValues := make([]string, 0, len(rackConfigMap))
	for _, value := range rackConfigMap {
		rackConfigValues = append(rackConfigValues, value)
	}
	brokerConfigs := []banzaicloudv1alpha1.BrokerConfig{}

	for _, brokerConfig := range cr.Spec.BrokerConfigs {
		if strconv.Itoa(int(brokerConfig.Id)) == pod.Labels["brokerId"] {
			if !strings.Contains(brokerConfig.Config, "broker.rack=") {
				config := brokerConfig.Config + fmt.Sprintf("broker.rack=%s\n", strings.Join(rackConfigValues, ","))
				brokerConfig.Config = config
			}
		}
		brokerConfigs = append(brokerConfigs, brokerConfig)
	}
	cr.Spec.BrokerConfigs = brokerConfigs
	return updateCr(cr, client)
}

// AddNewBrokerToCr modifies the CR and adds a new broker
func AddNewBrokerToCr(brokerConfig banzaicloudv1alpha1.BrokerConfig, crName, namespace string, client runtimeClient.Client) error {
	cr, err := GetCr(crName, namespace, client)
	if err != nil {
		return err
	}
	cr.Spec.BrokerConfigs = append(cr.Spec.BrokerConfigs, brokerConfig)

	return updateCr(cr, client)
}

// RemoveBrokerFromCr modifies the CR and removes the given broker from the cluster
func RemoveBrokerFromCr(brokerId, crName, namespace string, client runtimeClient.Client) error {

	cr, err := GetCr(crName, namespace, client)
	if err != nil {
		return err
	}

	tmpBrokers := cr.Spec.BrokerConfigs[:0]
	for _, broker := range cr.Spec.BrokerConfigs {
		if strconv.Itoa(int(broker.Id)) != brokerId {
			tmpBrokers = append(tmpBrokers, broker)
		}
	}
	cr.Spec.BrokerConfigs = tmpBrokers
	return updateCr(cr, client)
}

// AddPvToSpecificBroker adds a new PV to a specific broker
func AddPvToSpecificBroker(brokerId, crName, namespace string, storageConfig *banzaicloudv1alpha1.StorageConfig, client runtimeClient.Client) error {
	cr, err := GetCr(crName, namespace, client)
	if err != nil {
		return err
	}
	tempConfigs := cr.Spec.BrokerConfigs[:0]
	for _, brokerConfig := range cr.Spec.BrokerConfigs {
		if strconv.Itoa(int(brokerConfig.Id)) == brokerId {
			brokerConfig.StorageConfigs = append(brokerConfig.StorageConfigs, *storageConfig)
		}
		tempConfigs = append(tempConfigs, brokerConfig)
	}

	cr.Spec.BrokerConfigs = tempConfigs
	return updateCr(cr, client)
}

// GetCr returns the given cr object
func GetCr(name, namespace string, client runtimeClient.Client) (*banzaicloudv1alpha1.KafkaCluster, error) {
	cr := &banzaicloudv1alpha1.KafkaCluster{}

	err := client.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: namespace}, cr)
	if err != nil {
		return nil, emperror.WrapWith(err, "could not get cr from k8s", "crName", name, "namespace", namespace)
	}
	return cr, nil
}

func updateCr(cr *banzaicloudv1alpha1.KafkaCluster, client runtimeClient.Client) error {
	err := client.Update(context.TODO(), cr)
	if err != nil {
		return err
	}
	return nil
}
