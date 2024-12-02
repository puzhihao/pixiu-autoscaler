/*
Copyright 2021 The Pixiu Authors.

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

package controller

const (
	PixiuManager    string = "kubez-autoscaler-controller"
	PixiuMain       string = "main" // Just for local test
	PixiuRootPrefix string = "hpa.caoyingjunz.io"
	PixiuSeparator  string = "/"
	PixiuDot        string = "."

	MinReplicas              string = "hpa.caoyingjunz.io/minReplicas"
	MaxReplicas              string = "hpa.caoyingjunz.io/maxReplicas"
	targetAverageUtilization string = "targetAverageUtilization"
	targetAverageValue       string = "targetAverageValue"

	cpu        string = "cpu"
	memory     string = "memory"
	prometheus string = "prometheus"

	cpuAverageUtilization        = "cpu." + PixiuRootPrefix + PixiuSeparator + targetAverageUtilization
	memoryAverageUtilization     = "memory." + PixiuRootPrefix + PixiuSeparator + targetAverageUtilization
	prometheusAverageUtilization = "prometheus." + PixiuRootPrefix + PixiuSeparator + targetAverageUtilization

	// PrometheusCustomMetric 指标来自 prometheus 时，需要指定指标名称
	PrometheusCustomMetric = PixiuRootPrefix + PixiuSeparator + "targetCustomMetric"

	// CPU, in cores. (500m = .5 cores)
	// Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)
	cpuAverageValue        = "cpu." + PixiuRootPrefix + PixiuSeparator + targetAverageValue
	memoryAverageValue     = "memory." + PixiuRootPrefix + PixiuSeparator + targetAverageValue
	prometheusAverageValue = "prometheus." + PixiuRootPrefix + PixiuSeparator + targetAverageValue
)

const (
	AppsAPIVersion        string = "apps/v1"
	AutoscalingAPIVersion string = "autoscaling/v2"

	Deployment              string = "Deployment"
	HorizontalPodAutoscaler string = "HorizontalPodAutoscaler"

	DesireConfigMapName string = "prometheus-adapter"
	NotifyAt            string = PixiuRootPrefix + PixiuSeparator + "notifyAt"
)

type PrometheusAdapterConfig struct {
	ExternalRules []ExternalRule `yaml:"externalRules"`
}

type ExternalRule struct {
	MetricsQuery string   `yaml:"metricsQuery"`
	Name         RuleName `yaml:"name"`
	Resources    Resource `yaml:"resources"`
	SeriesQuery  string   `yaml:"seriesQuery"`
}

type RuleName struct {
	As      string `yaml:"as"`
	Matches string `yaml:"matches"`
}

type Resource struct {
	Namespaced bool `yaml:"namespace"`
}
