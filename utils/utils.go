package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)


type EnvironmentConfig struct {
	environment          string
	project              string
	machine_type         string
	region               string
	credential_file_path string
	zone                 string
	app_name             string
	github_repo_name     string
}

func ConvertTfConfigToJson(fileName string) EnvironmentConfig {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jsonMap := map[string]interface{}{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "=")
		value := s[1]
		jsonMap[s[0]] = strings.TrimSuffix(value[1:len(value)-1], "\n") // this will remove quote & breakline
	}

	fmt.Println(jsonMap["app_name"])
	if scanner.Err() != nil {
		log.Fatal(err)
	}

	return EnvironmentConfig{
		environment:          jsonMap["environment"].(string),
		project:              jsonMap["project"].(string),
		machine_type:         jsonMap["machine_type"].(string),
		region:               jsonMap["region"].(string),
		credential_file_path: jsonMap["credential_file_path"].(string),
		zone:                 jsonMap["zone"].(string),
		app_name:             jsonMap["app_name"].(string),
		github_repo_name:     jsonMap["github_repo_name"].(string),
	}
}