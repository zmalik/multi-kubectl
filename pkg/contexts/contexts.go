package contexts

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// ContextName we are only interested in context name
type ContextName struct {
	Name string `yaml:"name"`
}

// KubeConfig representation to just get list of contexts
type KubeConfig struct {
	Contexts       []*ContextName `yaml:"contexts"`
	CurrentContext string         `yaml:"current-context,omitempty"`
}

// NewKubeConfigFromFile create a new struct KubeConfig
func NewKubeConfigFromFile(kubeconfig string) (*KubeConfig, error) {
	if _, err := os.Stat(kubeconfig); os.IsNotExist(err) {
		return nil, fmt.Errorf("File doesn't exist: %s", err)
	}

	content, err := ioutil.ReadFile(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("Error when reading kubeconfig file: %s", err)
	}

	k := &KubeConfig{}

	if err = k.Unmarshal(content); err != nil {
		return nil, fmt.Errorf("Can't unmarshal the kubeconfig file: %s", err)
	}

	return k, nil
}

// Unmarshal fill a kubeConfig struct with yaml.Unmarshal
func (k *KubeConfig) Unmarshal(config []byte) error {
	err := yaml.Unmarshal(config, k)
	if err != nil {
		return err
	}

	return nil
}

// ContextExists check if a context exists
func (k *KubeConfig) ContextExists(context string) bool {
	for _, c := range k.Contexts {
		if context == c.Name {
			return true
		}
	}
	return false
}

//GetMatchedContexts return all contexts that contain string match
func (k *KubeConfig) GetMatchedContexts(match string) (matched []string) {
	for _, c := range k.Contexts {
		if strings.Contains(c.Name, match) {
			matched = append(matched, c.Name)
		}
	}
	return
}
