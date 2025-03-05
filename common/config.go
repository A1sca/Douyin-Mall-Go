package common

import (
	"io/ioutil"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"gopkg.in/yaml.v2"
)

// LoadConfig loads configuration from yaml file
func LoadConfig(path string, config interface{}) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		klog.Errorf("read config file error: %v", err)
		return err
	}

	if err = yaml.Unmarshal(content, config); err != nil {
		klog.Errorf("unmarshal config error: %v", err)
		return err
	}

	return nil
}