package main

import (
	"bytes"
	"gopkg.in/yaml.v2"
)

func parseAndTransformYAML(input []byte, newNamespace string, newName string) ([]byte, error) {
	var resources []KubernetesResource
	decoder := yaml.NewDecoder(bytes.NewReader(input))

	for {
		var resource KubernetesResource
		err := decoder.Decode(&resource)
		if err != nil {
			break
		}

		// Alterar o namespace e o nome
		resource.Metadata.Namespace = newNamespace
		resource.Metadata.Name = newName

		// Remover campos desnecessários
		switch resource.Kind {
		case "Deployment":
			var deployment Deployment
			err := yaml.Unmarshal(input, &deployment)
			if err != nil {
				return nil, err
			}
			deployment.Metadata.Namespace = newNamespace
			deployment.Metadata.Name = newName
			// Remova campos desnecessários aqui
			// Exemplo: remover status
			// deployment.Status = nil
			resources = append(resources, KubernetesResource{
				APIVersion: deployment.APIVersion,
				Kind:       deployment.Kind,
				Metadata:   deployment.Metadata,
				Spec:       deployment.Spec,
			})
		default:
			resources = append(resources, resource)
		}
	}

	var output bytes.Buffer
	encoder := yaml.NewEncoder(&output)
	for _, resource := range resources {
		err := encoder.Encode(&resource)
		if err != nil {
			return nil, err
		}
	}
	return output.Bytes(), nil
}
