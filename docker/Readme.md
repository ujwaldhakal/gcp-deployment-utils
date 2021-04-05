# docker
--
    import "github.com/ujwaldhakal/go-gcp-docker-utils/docker"


## Usage

#### func  Build
 
By default the build would does following steps
* Tries to pull image if image exists it wont pull `PullImageByUrl`
* If image does not exist it will build docker
* After the docker image has been build with `docker build`. It will creates image tag using `CreateImageTag`
* Finally `PushImage` will push image into container registry

```go
func Build(project string, githubRepoName string, commitHash string)
```

#### func  CreateImageTag
* `localBuildImage` is string path which will be the path of image in os i.e where the command runs
* `fullDockerUrl` is string which will be the final path of image in the Google Container Registry
```go
func CreateImageTag(localBuiltImagePath string, fullDockerUrl string)
```

#### func  Login
`filePath` is the path of google json credentials which will be used to authenticate into Google Container Registry
```go
func Login(filePath string)
```

#### func  PullImageByUrl
`imageUrl` is the path of image in the Google Container Registry. 
```go
func PullImageByUrl(fullGcrUrl string) error
```

#### func  PushImage
`fullGcrUrl` is the final image url which will be uploaded into Google Container Registry
```go
func PushImage(fullGcrUrl string)
```
