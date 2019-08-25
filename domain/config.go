package domain

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// Config struct
type Config struct {
	config map[string]interface{}
}

//SetFromBytes func to read bytes
func (c *Config) SetFromBytes(data []byte) error {
	var rawConfig interface{}
	if err := yaml.Unmarshal(data, &rawConfig); err != nil {
		return err
	}
	untypedConfig, ok := rawConfig.(map[interface{}]interface{})
	if !ok {
		return fmt.Errorf("not interface type")
	}

	typedConfig, err := convertKeysToString(untypedConfig)
	if err != nil {
		return err
	}
	c.config = typedConfig
	return nil
}

// Get returns the service for a perticular service
func (c *Config) Get(serviceName string) (map[string]interface{}, error) {
	a, ok := c.config["base"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("base config not a string interface")
	}
	// if no config exists for base service
	if _, ok := c.config[serviceName]; !ok {
		//return base config
		return nil, nil
	}

	b, ok := c.config[serviceName].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("service %q config is not a map ", serviceName)
	}

	//merge maps with service taking precedence
	config := make(map[string]interface{})
	for k, v := range a {
		config[k] = v
	}
	for k, v := range b {
		config[k] = v
	}
	return config, nil
}

func convertKeysToString(m map[interface{}]interface{}) (map[string]interface{}, error) {
	n := make(map[string]interface{})
	for k, v := range m {
		str, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf(" key is not string")
		}
		if vMap, ok := v.(map[interface{}]interface{}); ok {
			var err error
			v, err = convertKeysToString(vMap)
			if err != nil {
				return nil, fmt.Errorf(" key is not string")
			}
		}
		n[str] = v

	}
	return n, nil
}
