package container

import (
	"net/http"

	"github.com/jfrog/jfrog-cli-core/v2/artifactory/utils"
	"github.com/jfrog/jfrog-cli-core/v2/artifactory/utils/container"
	"github.com/jfrog/jfrog-cli-core/v2/utils/config"
	"github.com/jfrog/jfrog-client-go/utils/errorutils"
)

type ContainerManagerCommand struct {
	image              *container.Image
	repo               string
	buildConfiguration *utils.BuildConfiguration
	serverDetails      *config.ServerDetails
	skipLogin          bool
}

func (cmc *ContainerManagerCommand) ImageTag() string {
	return cmc.image.Name()
}

func (cmc *ContainerManagerCommand) SetImageTag(imageTag string) *ContainerManagerCommand {
	cmc.image = container.NewImage(imageTag)
	return cmc
}

func (cmc *ContainerManagerCommand) Repo() (string, error) {
	if cmc.repo != "" {
		return cmc.repo, nil
	}
	// Add version check it this is supported in artifactory
	containerRegistryUrl, err := cmc.image.GetRegistry()
	if err != nil {
		return "", err
	}
	imageName, err := cmc.image.GetImageBaseName()
	if err != nil {
		return "", err
	}
	imageTag := cmc.image.GetImageTag()
	endpoint := "/v2/" + imageName + "/manifests/" + imageTag
	serviceManager, err := utils.CreateServiceManager(cmc.serverDetails, -1, 0, false)
	if err != nil {
		return "", err
	}
	cd := serviceManager.GetConfig().GetServiceDetails().CreateHttpClientDetails()
	if serviceManager.GetConfig().GetServiceDetails().GetUrl()
	cd.u
	resp, body, err := serviceManager.Client().SendHead("https://"+containerRegistryUrl+endpoint, &cd)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errorutils.CheckErrorf("Artifactory response: " + resp.Status + "for" + string(body))
	}
	if dockerRepo := resp.Header["X-Artifactory-Docker-Registry"]; len(dockerRepo) != 0 {
		cmc.repo = dockerRepo[0]
		return cmc.repo, nil
	}
	return "", errorutils.CheckErrorf("couldn't find docker repository from artifactory")
}

func (cmc *ContainerManagerCommand) SetRepo(repo string) *ContainerManagerCommand {
	cmc.repo = repo
	return cmc
}

func (cmc *ContainerManagerCommand) BuildConfiguration() *utils.BuildConfiguration {
	return cmc.buildConfiguration
}

func (cmc *ContainerManagerCommand) SetBuildConfiguration(buildConfiguration *utils.BuildConfiguration) *ContainerManagerCommand {
	cmc.buildConfiguration = buildConfiguration
	return cmc
}

func (cmc *ContainerManagerCommand) SetSkipLogin(skipLogin bool) *ContainerManagerCommand {
	cmc.skipLogin = skipLogin
	return cmc
}

func (cmc *ContainerManagerCommand) ServerDetails() *config.ServerDetails {
	return cmc.serverDetails
}

func (cmc *ContainerManagerCommand) SetServerDetails(serverDetails *config.ServerDetails) *ContainerManagerCommand {
	cmc.serverDetails = serverDetails
	return cmc
}

func (cmc *ContainerManagerCommand) PerformLogin(serverDetails *config.ServerDetails, containerManagerType container.ContainerManagerType) error {
	if !cmc.skipLogin {
		// Exclude refreshable tokens when working with external tools (build tools, curl, etc)
		// Otherwise refresh Token may be expireted and docker login will fail.
		serverDetailsWithoutRefreshToken, err := config.GetSpecificConfig(serverDetails.ServerId, true, true)
		if err != nil {
			return err
		}
		loginConfig := &container.ContainerManagerLoginConfig{ServerDetails: serverDetailsWithoutRefreshToken}
		return container.ContainerManagerLogin(cmc.image, loginConfig, containerManagerType)
	}
	return nil
}
