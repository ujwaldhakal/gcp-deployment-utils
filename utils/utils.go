package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type EnvironmentConfig struct {
	Environment          string
	Project              string
	Machine_type         string
	Region               string
	Credential_file_path string
	Zone                 string
	App_name             string
	Github_repo_name     string
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
		Environment:          jsonMap["environment"].(string),
		Project:              jsonMap["project"].(string),
		Machine_type:         jsonMap["machine_type"].(string),
		Region:               jsonMap["region"].(string),
		Credential_file_path: jsonMap["credential_file_path"].(string),
		Zone:                 jsonMap["zone"].(string),
		App_name:             jsonMap["app_name"].(string),
		Github_repo_name:     jsonMap["github_repo_name"].(string),
	}
}
