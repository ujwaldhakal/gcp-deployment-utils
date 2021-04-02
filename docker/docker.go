package docker

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func PullImageByUrl(imageUrl string) error {
	p := exec.Command(
		"docker",
		"pull",
		imageUrl,
	)
	p.Stdout = os.Stdout
	p.Stderr = os.Stderr
	err := p.Run()

	return err
}

func Build(project string, githubRepoName string, commitHash string) {

	fullGcrUrl := fmt.Sprintf("gcr.io/%s/%s:%s", project, githubRepoName, commitHash)
	localBuiltImagePath := fmt.Sprintf("%s:%s", githubRepoName, commitHash)
	pullImage := PullImageByUrl(fullGcrUrl)

	if pullImage != nil {
		buildImage := exec.Command(
			"docker",
			"build",
			"-t",
			localBuiltImagePath,
			"../",
			"--target",
			"production",
		)
		buildImage.Stdout = os.Stdout
		buildImage.Stderr = os.Stderr
		buildImage.Run()
		CreateImageTag(localBuiltImagePath, fullGcrUrl)
		PushImage(fullGcrUrl)
	}
}

func PushImage(fullGcrUrl string) {
	push := exec.Command(
		"docker",
		"push",
		fullGcrUrl,
	)
	push.Stdout = os.Stdout
	push.Stderr = os.Stderr
	push.Run()
}

func CreateImageTag(localBuiltImagePath string, fullDockerUrl string) {
	tag := exec.Command(
		"docker",
		"tag",
		localBuiltImagePath,
		fullDockerUrl,
	)
	tag.Stdout = os.Stdout
	tag.Stderr = os.Stderr
	err := tag.Run()
	if err != nil {
		panic("Could not create tag")
	}
}

func Login(filePath string) {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", filePath)
	p := exec.Command(
		"docker",
		"login",
		"-u",
		"_json_key",
		"--password-stdin",
		"https://gcr.io",
	)

	file, _ := ioutil.ReadFile(filePath)
	p.Stdin = strings.NewReader(string(file))
	p.Stdout = os.Stdout
	p.Stderr = os.Stderr
	err := p.Run()
	if err != nil {
		panic("Could not login into docker")
	}
}
