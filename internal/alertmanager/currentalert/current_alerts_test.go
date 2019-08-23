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

package currentalert

//
//import (
//	stdlog "log"
//	"os"
//	"path/filepath"
//	"sync"
//	"testing"
//
//	"github.com/banzaicloud/kafka-operator/api"
//	"github.com/onsi/gomega"
//	"github.com/prometheus/common/model"
//	"k8s.io/client-go/kubernetes/scheme"
//	"k8s.io/client-go/rest"
//	"sigs.k8s.io/controller-runtime/pkg/client"
//	"sigs.k8s.io/controller-runtime/pkg/envtest"
//	"sigs.k8s.io/controller-runtime/pkg/manager"
//)
//
//var c client.Client
//var cfg *rest.Config
//
//func TestMain(m *testing.M) {
//	t := &envtest.Environment{
//		CRDDirectoryPaths: []string{filepath.Join("..", "..", "..", "config", "crds")},
//	}
//	api.AddToScheme(scheme.Scheme)
//
//	var err error
//	if cfg, err = t.Start(); err != nil {
//		stdlog.Fatal(err)
//	}
//
//	code := m.Run()
//	t.Stop()
//	os.Exit(code)
//}
//
//// StartTestManager adds recFn
//func StartTestManager(mgr manager.Manager, g *gomega.GomegaWithT) (chan struct{}, *sync.WaitGroup) {
//	stop := make(chan struct{})
//	wg := &sync.WaitGroup{}
//	go func() {
//		wg.Add(1)
//		g.Expect(mgr.Start(stop)).NotTo(gomega.HaveOccurred())
//		wg.Done()
//	}()
//	return stop, wg
//}
//
//func TestGetCurrentAlerts(t *testing.T) {
//	g := gomega.NewGomegaWithT(t)
//
//	mgr, err := manager.New(cfg, manager.Options{})
//	g.Expect(err).NotTo(gomega.HaveOccurred())
//	c = mgr.GetClient()
//
//	stopMgr, mgrStopped := StartTestManager(mgr, g)
//
//	defer func() {
//		close(stopMgr)
//		mgrStopped.Wait()
//	}()
//
//	alerts1 := GetCurrentAlerts()
//	if alerts1 == nil {
//		t.Error("expected pointer to Singleton after calling GetCurrentAlerts(), not nil")
//	}
//
//	singleAlerts := alerts1
//
//	testAlert1 := AlertState{
//		FingerPrint: model.Fingerprint(1111),
//		Status:      model.AlertStatus("firing"),
//		Labels: model.LabelSet{
//			"alertname": "PodAllert",
//			"test":      "test",
//		},
//	}
//
//	testAlert2 := AlertState{
//		FingerPrint: model.Fingerprint(1111),
//		Status:      model.AlertStatus("resolved"),
//		Labels: model.LabelSet{
//			"alertname": "PodAllert",
//			"test":      "test",
//		},
//	}
//
//	a1 := alerts1.AddAlert(testAlert1)
//	if a1.Status != "firing" {
//		t.Error("AdAlert failed a1")
//	}
//
//	list1 := alerts1.ListAlerts()
//	if list1 == nil || list1[testAlert1.FingerPrint].Status != "firing" || list1[testAlert1.FingerPrint].Labels["alertname"] != "PodAllert" {
//		t.Error("Listing alerts failed a1")
//	}
//
//	currAlert, err := alerts1.HandleAlert(testAlert1.FingerPrint, c)
//	if err != nil {
//		t.Error("Hanlde alert failed a1 with error")
//	}
//	t.Log(currAlert)
//
//	if list1 == nil || list1[testAlert1.FingerPrint].Status != "firing" || list1[testAlert1.FingerPrint].Processed != true {
//		t.Error("handle alert failed a1")
//	}
//
//	alerts2 := GetCurrentAlerts()
//	if alerts2 != singleAlerts {
//		t.Error("Expected same instance in alerts2 but it got a different instance")
//	}
//
//	a2 := alerts2.AddAlert(testAlert2)
//	if a2.Status != "resolved" {
//		t.Error("AdAlert failed a2")
//	}
//
//	list2 := alerts2.ListAlerts()
//	if list2 == nil || list2[testAlert2.FingerPrint].Status != "resolved" || list2[testAlert2.FingerPrint].Labels["alertname"] != "PodAllert" {
//		t.Error("Listing alerts failed a2")
//	}
//
//	alerts3 := GetCurrentAlerts()
//	if alerts3.AlertGC(testAlert2) != nil {
//		t.Error("Unable to delete alert a2")
//	}
//
//	list3 := alerts3.ListAlerts()
//	if list3 == nil || list3[testAlert2.FingerPrint] != nil {
//		t.Error("1111 alert wasn't deleted")
//	}
//
//	_, err = alerts3.HandleAlert(model.Fingerprint(1111), c)
//	expected := "alert doesn't exist"
//	if err == nil || err.Error() != expected {
//		t.Errorf("alert with 2222 should be %s", err)
//	}
//}
