/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package customclient

import (
	"fmt"

	alertingv1alpha1 "github.com/caicloud/clientset/customclient/typed/alerting/v1alpha1"
	alertingv1alpha2 "github.com/caicloud/clientset/customclient/typed/alerting/v1alpha2"
	alertingv1beta1 "github.com/caicloud/clientset/customclient/typed/alerting/v1beta1"
	apiregistrationv1 "github.com/caicloud/clientset/customclient/typed/apiregistration/v1"
	cleverv1alpha1 "github.com/caicloud/clientset/customclient/typed/clever/v1alpha1"
	cleverv1alpha2 "github.com/caicloud/clientset/customclient/typed/clever/v1alpha2"
	cleverv1alpha3 "github.com/caicloud/clientset/customclient/typed/clever/v1alpha3"
	cnetworkingv1alpha1 "github.com/caicloud/clientset/customclient/typed/cnetworking/v1alpha1"
	configv1alpha1 "github.com/caicloud/clientset/customclient/typed/config/v1alpha1"
	datasetv1alpha1 "github.com/caicloud/clientset/customclient/typed/dataset/v1alpha1"
	datasetv1alpha2 "github.com/caicloud/clientset/customclient/typed/dataset/v1alpha2"
	devopsv1 "github.com/caicloud/clientset/customclient/typed/devops/v1"
	evaluationv1alpha1 "github.com/caicloud/clientset/customclient/typed/evaluation/v1alpha1"
	loadbalancev1alpha2 "github.com/caicloud/clientset/customclient/typed/loadbalance/v1alpha2"
	loggingv1alpha1 "github.com/caicloud/clientset/customclient/typed/logging/v1alpha1"
	microservicev1alpha1 "github.com/caicloud/clientset/customclient/typed/microservice/v1alpha1"
	modelv1alpha1 "github.com/caicloud/clientset/customclient/typed/model/v1alpha1"
	orchestrationv1alpha1 "github.com/caicloud/clientset/customclient/typed/orchestration/v1alpha1"
	releasev1alpha1 "github.com/caicloud/clientset/customclient/typed/release/v1alpha1"
	resourcev1alpha1 "github.com/caicloud/clientset/customclient/typed/resource/v1alpha1"
	resourcev1beta1 "github.com/caicloud/clientset/customclient/typed/resource/v1beta1"
	servicemeshv1alpha1 "github.com/caicloud/clientset/customclient/typed/servicemesh/v1alpha1"
	servingv1alpha1 "github.com/caicloud/clientset/customclient/typed/serving/v1alpha1"
	tenantv1alpha1 "github.com/caicloud/clientset/customclient/typed/tenant/v1alpha1"
	workloadv1alpha1 "github.com/caicloud/clientset/customclient/typed/workload/v1alpha1"
	workloadv1beta1 "github.com/caicloud/clientset/customclient/typed/workload/v1beta1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AlertingV1beta1() alertingv1beta1.AlertingV1beta1Interface
	AlertingV1alpha2() alertingv1alpha2.AlertingV1alpha2Interface
	AlertingV1alpha1() alertingv1alpha1.AlertingV1alpha1Interface
	ApiregistrationV1() apiregistrationv1.ApiregistrationV1Interface
	CleverV1alpha3() cleverv1alpha3.CleverV1alpha3Interface
	CleverV1alpha2() cleverv1alpha2.CleverV1alpha2Interface
	CleverV1alpha1() cleverv1alpha1.CleverV1alpha1Interface
	CnetworkingV1alpha1() cnetworkingv1alpha1.CnetworkingV1alpha1Interface
	ConfigV1alpha1() configv1alpha1.ConfigV1alpha1Interface
	DatasetV1alpha2() datasetv1alpha2.DatasetV1alpha2Interface
	DatasetV1alpha1() datasetv1alpha1.DatasetV1alpha1Interface
	DevopsV1() devopsv1.DevopsV1Interface
	EvaluationV1alpha1() evaluationv1alpha1.EvaluationV1alpha1Interface
	LoadbalanceV1alpha2() loadbalancev1alpha2.LoadbalanceV1alpha2Interface
	LoggingV1alpha1() loggingv1alpha1.LoggingV1alpha1Interface
	MicroserviceV1alpha1() microservicev1alpha1.MicroserviceV1alpha1Interface
	ModelV1alpha1() modelv1alpha1.ModelV1alpha1Interface
	OrchestrationV1alpha1() orchestrationv1alpha1.OrchestrationV1alpha1Interface
	ReleaseV1alpha1() releasev1alpha1.ReleaseV1alpha1Interface
	ResourceV1beta1() resourcev1beta1.ResourceV1beta1Interface
	ResourceV1alpha1() resourcev1alpha1.ResourceV1alpha1Interface
	ServicemeshV1alpha1() servicemeshv1alpha1.ServicemeshV1alpha1Interface
	ServingV1alpha1() servingv1alpha1.ServingV1alpha1Interface
	TenantV1alpha1() tenantv1alpha1.TenantV1alpha1Interface
	WorkloadV1beta1() workloadv1beta1.WorkloadV1beta1Interface
	WorkloadV1alpha1() workloadv1alpha1.WorkloadV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	alertingV1beta1       *alertingv1beta1.AlertingV1beta1Client
	alertingV1alpha2      *alertingv1alpha2.AlertingV1alpha2Client
	alertingV1alpha1      *alertingv1alpha1.AlertingV1alpha1Client
	apiregistrationV1     *apiregistrationv1.ApiregistrationV1Client
	cleverV1alpha3        *cleverv1alpha3.CleverV1alpha3Client
	cleverV1alpha2        *cleverv1alpha2.CleverV1alpha2Client
	cleverV1alpha1        *cleverv1alpha1.CleverV1alpha1Client
	cnetworkingV1alpha1   *cnetworkingv1alpha1.CnetworkingV1alpha1Client
	configV1alpha1        *configv1alpha1.ConfigV1alpha1Client
	datasetV1alpha2       *datasetv1alpha2.DatasetV1alpha2Client
	datasetV1alpha1       *datasetv1alpha1.DatasetV1alpha1Client
	devopsV1              *devopsv1.DevopsV1Client
	evaluationV1alpha1    *evaluationv1alpha1.EvaluationV1alpha1Client
	loadbalanceV1alpha2   *loadbalancev1alpha2.LoadbalanceV1alpha2Client
	loggingV1alpha1       *loggingv1alpha1.LoggingV1alpha1Client
	microserviceV1alpha1  *microservicev1alpha1.MicroserviceV1alpha1Client
	modelV1alpha1         *modelv1alpha1.ModelV1alpha1Client
	orchestrationV1alpha1 *orchestrationv1alpha1.OrchestrationV1alpha1Client
	releaseV1alpha1       *releasev1alpha1.ReleaseV1alpha1Client
	resourceV1beta1       *resourcev1beta1.ResourceV1beta1Client
	resourceV1alpha1      *resourcev1alpha1.ResourceV1alpha1Client
	servicemeshV1alpha1   *servicemeshv1alpha1.ServicemeshV1alpha1Client
	servingV1alpha1       *servingv1alpha1.ServingV1alpha1Client
	tenantV1alpha1        *tenantv1alpha1.TenantV1alpha1Client
	workloadV1beta1       *workloadv1beta1.WorkloadV1beta1Client
	workloadV1alpha1      *workloadv1alpha1.WorkloadV1alpha1Client
}

// AlertingV1beta1 retrieves the AlertingV1beta1Client
func (c *Clientset) AlertingV1beta1() alertingv1beta1.AlertingV1beta1Interface {
	return c.alertingV1beta1
}

// AlertingV1alpha2 retrieves the AlertingV1alpha2Client
func (c *Clientset) AlertingV1alpha2() alertingv1alpha2.AlertingV1alpha2Interface {
	return c.alertingV1alpha2
}

// AlertingV1alpha1 retrieves the AlertingV1alpha1Client
func (c *Clientset) AlertingV1alpha1() alertingv1alpha1.AlertingV1alpha1Interface {
	return c.alertingV1alpha1
}

// ApiregistrationV1 retrieves the ApiregistrationV1Client
func (c *Clientset) ApiregistrationV1() apiregistrationv1.ApiregistrationV1Interface {
	return c.apiregistrationV1
}

// CleverV1alpha3 retrieves the CleverV1alpha3Client
func (c *Clientset) CleverV1alpha3() cleverv1alpha3.CleverV1alpha3Interface {
	return c.cleverV1alpha3
}

// CleverV1alpha2 retrieves the CleverV1alpha2Client
func (c *Clientset) CleverV1alpha2() cleverv1alpha2.CleverV1alpha2Interface {
	return c.cleverV1alpha2
}

// CleverV1alpha1 retrieves the CleverV1alpha1Client
func (c *Clientset) CleverV1alpha1() cleverv1alpha1.CleverV1alpha1Interface {
	return c.cleverV1alpha1
}

// CnetworkingV1alpha1 retrieves the CnetworkingV1alpha1Client
func (c *Clientset) CnetworkingV1alpha1() cnetworkingv1alpha1.CnetworkingV1alpha1Interface {
	return c.cnetworkingV1alpha1
}

// ConfigV1alpha1 retrieves the ConfigV1alpha1Client
func (c *Clientset) ConfigV1alpha1() configv1alpha1.ConfigV1alpha1Interface {
	return c.configV1alpha1
}

// DatasetV1alpha2 retrieves the DatasetV1alpha2Client
func (c *Clientset) DatasetV1alpha2() datasetv1alpha2.DatasetV1alpha2Interface {
	return c.datasetV1alpha2
}

// DatasetV1alpha1 retrieves the DatasetV1alpha1Client
func (c *Clientset) DatasetV1alpha1() datasetv1alpha1.DatasetV1alpha1Interface {
	return c.datasetV1alpha1
}

// DevopsV1 retrieves the DevopsV1Client
func (c *Clientset) DevopsV1() devopsv1.DevopsV1Interface {
	return c.devopsV1
}

// EvaluationV1alpha1 retrieves the EvaluationV1alpha1Client
func (c *Clientset) EvaluationV1alpha1() evaluationv1alpha1.EvaluationV1alpha1Interface {
	return c.evaluationV1alpha1
}

// LoadbalanceV1alpha2 retrieves the LoadbalanceV1alpha2Client
func (c *Clientset) LoadbalanceV1alpha2() loadbalancev1alpha2.LoadbalanceV1alpha2Interface {
	return c.loadbalanceV1alpha2
}

// LoggingV1alpha1 retrieves the LoggingV1alpha1Client
func (c *Clientset) LoggingV1alpha1() loggingv1alpha1.LoggingV1alpha1Interface {
	return c.loggingV1alpha1
}

// MicroserviceV1alpha1 retrieves the MicroserviceV1alpha1Client
func (c *Clientset) MicroserviceV1alpha1() microservicev1alpha1.MicroserviceV1alpha1Interface {
	return c.microserviceV1alpha1
}

// ModelV1alpha1 retrieves the ModelV1alpha1Client
func (c *Clientset) ModelV1alpha1() modelv1alpha1.ModelV1alpha1Interface {
	return c.modelV1alpha1
}

// OrchestrationV1alpha1 retrieves the OrchestrationV1alpha1Client
func (c *Clientset) OrchestrationV1alpha1() orchestrationv1alpha1.OrchestrationV1alpha1Interface {
	return c.orchestrationV1alpha1
}

// ReleaseV1alpha1 retrieves the ReleaseV1alpha1Client
func (c *Clientset) ReleaseV1alpha1() releasev1alpha1.ReleaseV1alpha1Interface {
	return c.releaseV1alpha1
}

// ResourceV1beta1 retrieves the ResourceV1beta1Client
func (c *Clientset) ResourceV1beta1() resourcev1beta1.ResourceV1beta1Interface {
	return c.resourceV1beta1
}

// ResourceV1alpha1 retrieves the ResourceV1alpha1Client
func (c *Clientset) ResourceV1alpha1() resourcev1alpha1.ResourceV1alpha1Interface {
	return c.resourceV1alpha1
}

// ServicemeshV1alpha1 retrieves the ServicemeshV1alpha1Client
func (c *Clientset) ServicemeshV1alpha1() servicemeshv1alpha1.ServicemeshV1alpha1Interface {
	return c.servicemeshV1alpha1
}

// ServingV1alpha1 retrieves the ServingV1alpha1Client
func (c *Clientset) ServingV1alpha1() servingv1alpha1.ServingV1alpha1Interface {
	return c.servingV1alpha1
}

// TenantV1alpha1 retrieves the TenantV1alpha1Client
func (c *Clientset) TenantV1alpha1() tenantv1alpha1.TenantV1alpha1Interface {
	return c.tenantV1alpha1
}

// WorkloadV1beta1 retrieves the WorkloadV1beta1Client
func (c *Clientset) WorkloadV1beta1() workloadv1beta1.WorkloadV1beta1Interface {
	return c.workloadV1beta1
}

// WorkloadV1alpha1 retrieves the WorkloadV1alpha1Client
func (c *Clientset) WorkloadV1alpha1() workloadv1alpha1.WorkloadV1alpha1Interface {
	return c.workloadV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("Burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.alertingV1beta1, err = alertingv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.alertingV1alpha2, err = alertingv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.alertingV1alpha1, err = alertingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.apiregistrationV1, err = apiregistrationv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.cleverV1alpha3, err = cleverv1alpha3.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.cleverV1alpha2, err = cleverv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.cleverV1alpha1, err = cleverv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.cnetworkingV1alpha1, err = cnetworkingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.configV1alpha1, err = configv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.datasetV1alpha2, err = datasetv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.datasetV1alpha1, err = datasetv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.devopsV1, err = devopsv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.evaluationV1alpha1, err = evaluationv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.loadbalanceV1alpha2, err = loadbalancev1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.loggingV1alpha1, err = loggingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.microserviceV1alpha1, err = microservicev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.modelV1alpha1, err = modelv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.orchestrationV1alpha1, err = orchestrationv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.releaseV1alpha1, err = releasev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.resourceV1beta1, err = resourcev1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.resourceV1alpha1, err = resourcev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.servicemeshV1alpha1, err = servicemeshv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.servingV1alpha1, err = servingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.tenantV1alpha1, err = tenantv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.workloadV1beta1, err = workloadv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.workloadV1alpha1, err = workloadv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.alertingV1beta1 = alertingv1beta1.NewForConfigOrDie(c)
	cs.alertingV1alpha2 = alertingv1alpha2.NewForConfigOrDie(c)
	cs.alertingV1alpha1 = alertingv1alpha1.NewForConfigOrDie(c)
	cs.apiregistrationV1 = apiregistrationv1.NewForConfigOrDie(c)
	cs.cleverV1alpha3 = cleverv1alpha3.NewForConfigOrDie(c)
	cs.cleverV1alpha2 = cleverv1alpha2.NewForConfigOrDie(c)
	cs.cleverV1alpha1 = cleverv1alpha1.NewForConfigOrDie(c)
	cs.cnetworkingV1alpha1 = cnetworkingv1alpha1.NewForConfigOrDie(c)
	cs.configV1alpha1 = configv1alpha1.NewForConfigOrDie(c)
	cs.datasetV1alpha2 = datasetv1alpha2.NewForConfigOrDie(c)
	cs.datasetV1alpha1 = datasetv1alpha1.NewForConfigOrDie(c)
	cs.devopsV1 = devopsv1.NewForConfigOrDie(c)
	cs.evaluationV1alpha1 = evaluationv1alpha1.NewForConfigOrDie(c)
	cs.loadbalanceV1alpha2 = loadbalancev1alpha2.NewForConfigOrDie(c)
	cs.loggingV1alpha1 = loggingv1alpha1.NewForConfigOrDie(c)
	cs.microserviceV1alpha1 = microservicev1alpha1.NewForConfigOrDie(c)
	cs.modelV1alpha1 = modelv1alpha1.NewForConfigOrDie(c)
	cs.orchestrationV1alpha1 = orchestrationv1alpha1.NewForConfigOrDie(c)
	cs.releaseV1alpha1 = releasev1alpha1.NewForConfigOrDie(c)
	cs.resourceV1beta1 = resourcev1beta1.NewForConfigOrDie(c)
	cs.resourceV1alpha1 = resourcev1alpha1.NewForConfigOrDie(c)
	cs.servicemeshV1alpha1 = servicemeshv1alpha1.NewForConfigOrDie(c)
	cs.servingV1alpha1 = servingv1alpha1.NewForConfigOrDie(c)
	cs.tenantV1alpha1 = tenantv1alpha1.NewForConfigOrDie(c)
	cs.workloadV1beta1 = workloadv1beta1.NewForConfigOrDie(c)
	cs.workloadV1alpha1 = workloadv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.alertingV1beta1 = alertingv1beta1.New(c)
	cs.alertingV1alpha2 = alertingv1alpha2.New(c)
	cs.alertingV1alpha1 = alertingv1alpha1.New(c)
	cs.apiregistrationV1 = apiregistrationv1.New(c)
	cs.cleverV1alpha3 = cleverv1alpha3.New(c)
	cs.cleverV1alpha2 = cleverv1alpha2.New(c)
	cs.cleverV1alpha1 = cleverv1alpha1.New(c)
	cs.cnetworkingV1alpha1 = cnetworkingv1alpha1.New(c)
	cs.configV1alpha1 = configv1alpha1.New(c)
	cs.datasetV1alpha2 = datasetv1alpha2.New(c)
	cs.datasetV1alpha1 = datasetv1alpha1.New(c)
	cs.devopsV1 = devopsv1.New(c)
	cs.evaluationV1alpha1 = evaluationv1alpha1.New(c)
	cs.loadbalanceV1alpha2 = loadbalancev1alpha2.New(c)
	cs.loggingV1alpha1 = loggingv1alpha1.New(c)
	cs.microserviceV1alpha1 = microservicev1alpha1.New(c)
	cs.modelV1alpha1 = modelv1alpha1.New(c)
	cs.orchestrationV1alpha1 = orchestrationv1alpha1.New(c)
	cs.releaseV1alpha1 = releasev1alpha1.New(c)
	cs.resourceV1beta1 = resourcev1beta1.New(c)
	cs.resourceV1alpha1 = resourcev1alpha1.New(c)
	cs.servicemeshV1alpha1 = servicemeshv1alpha1.New(c)
	cs.servingV1alpha1 = servingv1alpha1.New(c)
	cs.tenantV1alpha1 = tenantv1alpha1.New(c)
	cs.workloadV1beta1 = workloadv1beta1.New(c)
	cs.workloadV1alpha1 = workloadv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
