package yaml

import (
	"github.com/jsipling/toolbox/fileutils"
	"gopkg.in/yaml.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ToString(filepath string) string {
	config := fileutils.Readfile(filepath)

	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(config), &m)
	check(err)

	d, err := yaml.Marshal(&m)
	check(err)

	return string(d)
}

func ToMap(filepath string) map[interface{}]interface{} {
	config := fileutils.Readfile(filepath)

	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(config), &m)
	check(err)

	return m
}

func ToYaml(m map[interface{}]interface{}) string {
	d, err := yaml.Marshal(&m)
	check(err)

	return string(d)
}
