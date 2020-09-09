package test

import (
	"testing"
)

func Test_YamlSpecFNNamespace_NoOverrides(t *testing.T) {
	parts := []string{}
	want := buildFnNetworkPolicy("openfaas-fn")
	runYamlTest(parts, "./tmp/openfaas-cloud/templates/network-policy/ns-openfaas-fn-net-policy.yml", want, t)
}

func Test_YamlSpecFNNamespace_Overrides(t *testing.T) {
	parts := []string{
		"--set", "global.functionsNamespace=some-fn-namespace",
	}
	want := buildFnNetworkPolicy("some-fn-namespace")
	runYamlTest(parts, "./tmp/openfaas-cloud/templates/network-policy/ns-openfaas-fn-net-policy.yml", want, t)

}

func Test_CoreNetworkNamespace_NoOverrides(t *testing.T) {
	parts := []string{}
	want := buildCoreNetworkPolicy("openfaas", "openfaas-fn")
	runYamlTest(parts, "./tmp/openfaas-cloud/templates/network-policy/ns-openfaas-net-policy.yml", want, t)
}

func Test_CoreNetworkPolicy_Overrides(t *testing.T) {
	parts := []string{
		"--set", "global.functionsNamespace=some-fn-namespace",
		"--set", "global.coreNamespace=some-namespace",
	}
	want := buildCoreNetworkPolicy("some-namespace", "some-fn-namespace")
	runYamlTest(parts, "./tmp/openfaas-cloud/templates/network-policy/ns-openfaas-net-policy.yml", want, t)
}

func Test_EdgeRouterNetworkPolicy_NoOverrides(t *testing.T) {
	parts := []string{}
	want := buildEdgeRouterNetworkPolicy("openfaas", "openfaas-fn")
	runYamlTest(parts, "./tmp/openfaas-cloud/templates/network-policy/edge-router-net-policy.yml", want, t)
}

func Test_EdgeRouterNetworkPolicy_Overrides(t *testing.T) {
	parts := []string{
		"--set", "global.functionsNamespace=some-fn-namespace",
		"--set", "global.coreNamespace=some-namespace",
	}
	want := buildEdgeRouterNetworkPolicy("some-namespace", "some-fn-namespace")
	runYamlTest(parts, "./tmp/openfaas-cloud/templates/network-policy/edge-router-net-policy.yml", want, t)
}

func Test_GatewayNetworkPolicy_NoOverrides(t *testing.T) {
	parts := []string{}
	want := buildGatewayNetworkPolicy("openfaas", "openfaas-fn")
	runYamlTest(parts, "./tmp/openfaas-cloud/templates/network-policy/gateway-net-policy.yml", want, t)
}


func Test_GatewayNetworkPolicy_Overrides(t *testing.T) {
	parts := []string{
		"--set", "global.functionsNamespace=some-fn-namespace",
		"--set", "global.coreNamespace=some-namespace",
	}
	want := buildGatewayNetworkPolicy("some-namespace", "some-fn-namespace")
	runYamlTest(parts, "./tmp/openfaas-cloud/templates/network-policy/gateway-net-policy.yml", want, t)
}

func Test_PrometheusNetworkPolicy_NoOverrides(t *testing.T) {
	parts := []string{}
	want := buildPrometheusNetworkPolicy("openfaas", "openfaas-fn")
	runYamlTest(parts, "./tmp/openfaas-cloud/templates/network-policy/prometheus-net-policy.yml", want, t)
}


func Test_PrometheusNetworkPolicy_Overrides(t *testing.T) {
	parts := []string{
		"--set", "global.functionsNamespace=some-fn-namespace",
		"--set", "global.coreNamespace=some-namespace",
	}
	want := buildPrometheusNetworkPolicy("some-namespace", "some-fn-namespace")
	runYamlTest(parts, "./tmp/openfaas-cloud/templates/network-policy/prometheus-net-policy.yml", want, t)
}

func buildCoreNetworkPolicy(coreNamespace, functionNamespace string) YamlSpec {
	matchLabelsSystem := make(map[string]string)
	matchLabelsFunction := make(map[string]string)

	matchLabelsSystem["role"] = "openfaas-system"
	matchLabelsFunction["role"] = functionNamespace

	return YamlSpec{
		ApiVersion: "networking.k8s.io/v1",
		Kind:       "NetworkPolicy",
		Metadata: MetadataItems{
			Name:      coreNamespace,
			Namespace: coreNamespace,
		},
		Spec: Spec{
			PolicyTypes: []string{"Ingress"},
			PodSelector: MatchLabelSelector{},
			Ingress: []NetworkIngress{{
				From: []NetworkSelectors{
					{
						Namespace: NamespaceSelector{
							MatchLabels: matchLabelsSystem,
						},
					},
					{
						Namespace: NamespaceSelector{
							MatchLabels: matchLabelsFunction,
						},
						Pod: MatchLabelSelector{
							MatchLabels: matchLabelsSystem,
						},
					},
				},
			},
			},
		},
	}
}

func buildFnNetworkPolicy(functionNamespace string) YamlSpec {
	matchLabels := make(map[string]string)

	matchLabels["role"] = "openfaas-system"

	return YamlSpec{
		ApiVersion: "networking.k8s.io/v1",
		Kind:       "NetworkPolicy",
		Metadata: MetadataItems{
			Name:      functionNamespace,
			Namespace: functionNamespace,
		},
		Spec: Spec{
			PolicyTypes: []string{"Ingress"},
			PodSelector: MatchLabelSelector{},
			Ingress: []NetworkIngress{{
				From: []NetworkSelectors{
					{
						Namespace: NamespaceSelector{
							MatchLabels: matchLabels,
						},
					},
					{
						Pod: MatchLabelSelector{
							MatchLabels: matchLabels,
						},
					}},
			}},
		},
	}

}

func buildEdgeRouterNetworkPolicy(coreNamespace, functionNamespace string) YamlSpec {
	podSelector := make(map[string]string)
	nginxSelector := make(map[string]string)

	podSelector["app"] = "edge-router"
	nginxSelector["app.kubernetes.io/name"] = "ingress-nginx"

	return YamlSpec{
		ApiVersion: "networking.k8s.io/v1",
		Kind:       "NetworkPolicy",
		Metadata: MetadataItems{
			Name:      "edge-router",
			Namespace: coreNamespace,
		},
		Spec: Spec{
			PolicyTypes: []string{"Ingress"},
			PodSelector: MatchLabelSelector{
				MatchLabels: podSelector,
			},
			Ingress: []NetworkIngress{{
				From: []NetworkSelectors{
					{
						Namespace: NamespaceSelector{},
						Pod: MatchLabelSelector{
							MatchLabels: nginxSelector,
						},
					},
				},
			}},
		},
	}
}

func buildGatewayNetworkPolicy(coreNamespace, functionNamespace string) YamlSpec {
	podSelector := make(map[string]string)
	nginxSelector := make(map[string]string)
	auditEventSelector := make(map[string]string)
	matchLabelsFunction := make(map[string]string)

	podSelector["app"] = "gateway"
	nginxSelector["app.kubernetes.io/name"] = "ingress-nginx"
	auditEventSelector["faas_function"] = "audit-event"
	matchLabelsFunction["role"] = functionNamespace

	return YamlSpec{
		ApiVersion: "networking.k8s.io/v1",
		Kind:       "NetworkPolicy",
		Metadata: MetadataItems{
			Name:      "gateway",
			Namespace: coreNamespace,
		},
		Spec: Spec{
			PolicyTypes: []string{"Ingress"},
			PodSelector: MatchLabelSelector{
				MatchLabels: podSelector,
			},
			Ingress: []NetworkIngress{{
				From: []NetworkSelectors{
					{
						Namespace: NamespaceSelector{},
						Pod: MatchLabelSelector{
							MatchLabels: nginxSelector,
						},
					},
					{
						Namespace: NamespaceSelector{
							MatchLabels: matchLabelsFunction,
						},
						Pod: MatchLabelSelector{
							MatchLabels: auditEventSelector,
						},
					},
				},
			}},
		},
	}
}

func buildPrometheusNetworkPolicy(coreNamespace, functionNamespace string) YamlSpec {
	podSelector := make(map[string]string)
	metricsSelector := make(map[string]string)
	matchLabelsFunction := make(map[string]string)

	podSelector["app"] = "prometheus"
	metricsSelector["faas_function"] = "metrics"
	matchLabelsFunction["role"] = functionNamespace

	return YamlSpec{
		ApiVersion: "networking.k8s.io/v1",
		Kind:       "NetworkPolicy",
		Metadata: MetadataItems{
			Name:      "prometheus",
			Namespace: coreNamespace,
		},
		Spec: Spec{
			PolicyTypes: []string{"Ingress"},
			PodSelector: MatchLabelSelector{
				MatchLabels: podSelector,
			},
			Ingress: []NetworkIngress{{
				From: []NetworkSelectors{
					{
						Namespace: NamespaceSelector{
							MatchLabels: matchLabelsFunction,
						},
						Pod: MatchLabelSelector{
							MatchLabels: metricsSelector,
						},
					},
				},
			}},
		},
	}
}
