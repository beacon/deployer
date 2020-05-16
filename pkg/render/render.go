package render

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"sigs.k8s.io/yaml" // Avoid map[interface{}]interface{} to break json
)

// ConfigMap is typically a map of config
type ConfigMap map[string]interface{}

// Execute render given files using specified data
func Execute(config ConfigMap, file, outputDir string, inputs ...string) error {
	rawConfig, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Failed to parse config from file %s:%v", file, err)
	}
	if config == nil {
		config = make(ConfigMap)
	}
	if err := yaml.Unmarshal(rawConfig, &config); err != nil {
		return fmt.Errorf("Failed to unmarshal config from file %s:%v", file, err)
	}

	if err := os.MkdirAll(outputDir, 0744); err != nil {
		return fmt.Errorf("Failed to make dir %s:%v", outputDir, err)
	}

	for _, input := range inputs {
		fileInfo, err := os.Stat(input)
		if err != nil {
			return fmt.Errorf("failed to read input %s:%v", input, err)
		}
		if fileInfo.IsDir() {
			err = filepath.Walk(input, func(filePath string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				dstPath := path.Join(outputDir, input, strings.TrimPrefix(filePath, input))
				if err := renderFile(filePath, dstPath, config); err != nil {
					return fmt.Errorf("Failed to render file %s -> %s: %v", filePath, dstPath, err)
				}
				return nil
			})
			if err != nil {
				return fmt.Errorf("Failed to walk through input dir %s:%v", input, err)
			}
		} else {
			dstPath := path.Join(outputDir, input)
			if err := renderFile(input, dstPath, config); err != nil {
				return fmt.Errorf("Failed to render file %s -> %s: %v", input, dstPath, err)
			}
		}
	}
	return nil
}

func renderFile(srcFile, dstFile string, config ConfigMap) error {
	dstDir := path.Dir(dstFile)
	if err := os.MkdirAll(dstDir, 0777); err != nil {
		return fmt.Errorf("Failed to create dir %s:%v", dstDir, err)
	}
	content, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return fmt.Errorf("failed to load config %s:%v", srcFile, err)
	}
	tmpl, err := template.New(srcFile).Funcs(funcMap()).Parse(string(content))
	if err != nil {
		return fmt.Errorf("Failed to parse file %s as template:%v", srcFile, err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, config); err != nil {
		return fmt.Errorf("Failed to execute template with file %s:%v", srcFile, err)
	}
	if err := ioutil.WriteFile(dstFile, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("Failed to write file %s:%v", dstFile, err)
	}
	log.Println("Parsed file", srcFile, "->", dstFile)
	return nil
}
