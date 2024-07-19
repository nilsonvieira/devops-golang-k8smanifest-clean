package main

type Metadata struct {
	Name      string            `yaml:"name"`
	Namespace string            `yaml:"namespace"`
	Labels    map[string]string `yaml:"labels,omitempty"`
}

type DeploymentSpec struct {
	Replicas int `yaml:"replicas"`
	// Adicione outros campos conforme necessário
}

type Deployment struct {
	APIVersion string         `yaml:"apiVersion"`
	Kind       string         `yaml:"kind"`
	Metadata   Metadata       `yaml:"metadata"`
	Spec       DeploymentSpec `yaml:"spec"`
}

type KubernetesResource struct {
	APIVersion string      `yaml:"apiVersion"`
	Kind       string      `yaml:"kind"`
	Metadata   Metadata    `yaml:"metadata"`
	Spec       interface{} `yaml:"spec"`
}
