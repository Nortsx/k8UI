package kubernetes

import (
	"flag"
	"github.com/goccy/go-yaml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var kubeconfig *string

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func GetContexts() []string {
	if kubeconfig == nil {
		if home := homeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()
	}

	configYaml, _ := os.Open(*kubeconfig)
	byteValue, _ := ioutil.ReadAll(configYaml)

	path, err := yaml.PathString("$.clusters[*].name")
	if err != nil {
		return nil
	}
	s := string(byteValue)

	var names []string
	if err := path.Read(strings.NewReader(s), &names); err != nil {
		return nil
	}
	defer configYaml.Close()

	return names

}
