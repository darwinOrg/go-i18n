package dgi18n

import (
	"fmt"
	"os"
)

func WriteCommonYamlFile(rootPath string) {
	err := os.Mkdir(rootPath+"/common", 0750)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("create common dir error: %v", err)
		return
	}

	os.WriteFile(rootPath+"/common/common.en.yaml", []byte(EnYamlString), 0660)
	os.WriteFile(rootPath+"/common/common.zh.yaml", []byte(ZhYamlString), 0660)
}
