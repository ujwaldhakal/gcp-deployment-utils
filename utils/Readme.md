# utils
--
    import "."


## Usage

#### type EnvironmentConfig

```go
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
```


#### func  ConvertTfConfigToJson
It will convert the main.tf content into json data
```go
func ConvertTfConfigToJson(fileName string) EnvironmentConfig
```
