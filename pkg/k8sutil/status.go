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

	"emperror.dev/emperror"
	banzaicloudv1alpha1 "github.com/banzaicloud/kafka-operator/api/v1alpha1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func updateRackAwarenessStatus(c client.Client, brokerId string, cluster *banzaicloudv1alpha1.KafkaCluster, rackstatus banzaicloudv1alpha1.RackAwarenessState, logger logr.Logger) error {
	typeMeta := cluster.TypeMeta

	if cluster.Status.BrokersState == nil {
		cluster.Status.BrokersState = map[string]banzaicloudv1alpha1.BrokerState{brokerId: {RackAwarenessState: rackstatus}}
	} else if val, ok := cluster.Status.BrokersState[brokerId]; ok {
		val.RackAwarenessState = rackstatus
		cluster.Status.BrokersState[brokerId] = val
	} else {
		cluster.Status.BrokersState[brokerId] = banzaicloudv1alpha1.BrokerState{RackAwarenessState: rackstatus}
	}

	err := c.Status().Update(context.Background(), cluster)
	if errors.IsNotFound(err) {
		err = c.Update(context.Background(), cluster)
	}
	if err != nil {
		if !errors.IsConflict(err) {
			return emperror.Wrapf(err, "could not update Kafka broker %s rack state to '%s'", brokerId, rackstatus)
		}
		err := c.Get(context.TODO(), types.NamespacedName{
			Namespace: cluster.Namespace,
			Name:      cluster.Name,
		}, cluster)
		if err != nil {
			return emperror.Wrap(err, "could not get config for updating status")
		}

		if cluster.Status.BrokersState == nil {
			cluster.Status.BrokersState = map[string]banzaicloudv1alpha1.BrokerState{brokerId: {RackAwarenessState: rackstatus}}
		} else if val, ok := cluster.Status.BrokersState[brokerId]; ok {
			val.RackAwarenessState = rackstatus
			cluster.Status.BrokersState[brokerId] = val
		} else {
			cluster.Status.BrokersState[brokerId] = banzaicloudv1alpha1.BrokerState{RackAwarenessState: rackstatus}
		}

		err = c.Status().Update(context.Background(), cluster)
		if errors.IsNotFound(err) {
			err = c.Update(context.Background(), cluster)
		}
		if err != nil {
			return emperror.Wrapf(err, "could not update Kafka clusters broker %s rack state to '%s'", brokerId, rackstatus)
		}
	}
	// update loses the typeMeta of the config that's used later when setting ownerrefs
	cluster.TypeMeta = typeMeta
	logger.Info("Kafka cluster rack state updated", "status", rackstatus)
	return nil
}

func updateGracefulScaleStatus(c client.Client, brokerId string, cluster *banzaicloudv1alpha1.KafkaCluster, scaleStatus banzaicloudv1alpha1.GracefulActionState, logger logr.Logger) error {
	typeMeta := cluster.TypeMeta

	if cluster.Status.BrokersState == nil {
		cluster.Status.BrokersState = map[string]banzaicloudv1alpha1.BrokerState{brokerId: {GracefulActionState: scaleStatus}}
	} else if val, ok := cluster.Status.BrokersState[brokerId]; ok {
		val.GracefulActionState = scaleStatus
		cluster.Status.BrokersState[brokerId] = val
	} else {
		cluster.Status.BrokersState[brokerId] = banzaicloudv1alpha1.BrokerState{GracefulActionState: scaleStatus}
	}

	err := c.Status().Update(context.Background(), cluster)
	if errors.IsNotFound(err) {
		err = c.Update(context.Background(), cluster)
	}
	if err != nil {
		if !errors.IsConflict(err) {
			return emperror.Wrapf(err, "could not update Kafka broker %s graceful scale state to '%s'", brokerId, scaleStatus)
		}
		err := c.Get(context.TODO(), types.NamespacedName{
			Namespace: cluster.Namespace,
			Name:      cluster.Name,
		}, cluster)
		if err != nil {
			return emperror.Wrap(err, "could not get config for updating status")
		}

		if cluster.Status.BrokersState == nil {
			cluster.Status.BrokersState = map[string]banzaicloudv1alpha1.BrokerState{brokerId: {GracefulActionState: scaleStatus}}
		} else if val, ok := cluster.Status.BrokersState[brokerId]; ok {
			val.GracefulActionState = scaleStatus
			cluster.Status.BrokersState[brokerId] = val
		} else {
			cluster.Status.BrokersState[brokerId] = banzaicloudv1alpha1.BrokerState{GracefulActionState: scaleStatus}
		}

		err = c.Status().Update(context.Background(), cluster)
		if errors.IsNotFound(err) {
			err = c.Update(context.Background(), cluster)
		}
		if err != nil {
			return emperror.Wrapf(err, "could not update Kafka clusters broker %s graceful scale to '%s'", brokerId, scaleStatus)
		}
	}
	// update loses the typeMeta of the config that's used later when setting ownerrefs
	cluster.TypeMeta = typeMeta
	logger.Info("Kafka cluster graceful scale updated", "status", scaleStatus)
	return nil
}

// CCTopicStatus updates the given CC state in the CR
func UpdateCCTopicStatus(c client.Client, cluster *banzaicloudv1alpha1.KafkaCluster, ccTopicStatus banzaicloudv1alpha1.CruiseControlTopicStatus, logger logr.Logger) error {
	typeMeta := cluster.TypeMeta

	cluster.Status.CruiseControlTopicStatus = ccTopicStatus

	err := c.Status().Update(context.Background(), cluster)
	if errors.IsNotFound(err) {
		err = c.Update(context.Background(), cluster)
	}
	if err != nil {
		if !errors.IsConflict(err) {
			return emperror.Wrapf(err, "could not update CC topic state to '%s'", ccTopicStatus)
		}
		err := c.Get(context.TODO(), types.NamespacedName{
			Namespace: cluster.Namespace,
			Name:      cluster.Name,
		}, cluster)
		if err != nil {
			return emperror.Wrap(err, "could not get config for updating status")
		}

		cluster.Status.CruiseControlTopicStatus = ccTopicStatus

		err = c.Status().Update(context.Background(), cluster)
		if errors.IsNotFound(err) {
			err = c.Update(context.Background(), cluster)
		}
		if err != nil {
			return emperror.Wrapf(err, "could not update CC topic state to '%s'", ccTopicStatus)
		}
	}
	// update loses the typeMeta of the config that's used later when setting ownerrefs
	cluster.TypeMeta = typeMeta
	logger.Info("CC topic status updated", "status", ccTopicStatus)
	return nil
}

// DeleteStatus deletes the given broker state from the CR
func DeleteStatus(c client.Client, brokerId string, cluster *banzaicloudv1alpha1.KafkaCluster, logger logr.Logger) error {
	typeMeta := cluster.TypeMeta

	brokerStatus := cluster.Status.BrokersState

	delete(brokerStatus, brokerId)

	cluster.Status.BrokersState = brokerStatus

	err := c.Status().Update(context.Background(), cluster)
	if errors.IsNotFound(err) {
		err = c.Update(context.Background(), cluster)
	}
	if err != nil {
		if !errors.IsConflict(err) {
			return emperror.Wrapf(err, "could not delete Kafka cluster broker %s state ", brokerId)
		}
		err := c.Get(context.TODO(), types.NamespacedName{
			Namespace: cluster.Namespace,
			Name:      cluster.Name,
		}, cluster)
		if err != nil {
			return emperror.Wrap(err, "could not get config for updating status")
		}
		brokerStatus = cluster.Status.BrokersState

		delete(brokerStatus, brokerId)

		cluster.Status.BrokersState = brokerStatus
		err = c.Status().Update(context.Background(), cluster)
		if errors.IsNotFound(err) {
			err = c.Update(context.Background(), cluster)
		}
		if err != nil {
			return emperror.Wrapf(err, "could not delete Kafka clusters broker %s state ", brokerId)
		}
	}

	// update loses the typeMeta of the config that's used later when setting ownerrefs
	cluster.TypeMeta = typeMeta
	logger.Info(fmt.Sprintf("Kafka broker %s state deleted", brokerId))
	return nil
}
