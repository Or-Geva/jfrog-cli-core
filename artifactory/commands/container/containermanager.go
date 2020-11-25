package container

import (
	"github.com/jfrog/jfrog-cli-core/artifactory/utils"
	"github.com/jfrog/jfrog-cli-core/utils/config"
)

type ContainerManagerCommand struct {
	imageTag           string
	repo               string
	buildConfiguration *utils.BuildConfiguration
	rtDetails          *config.ArtifactoryDetails
	skipLogin          bool
}

func (cmc *ContainerManagerCommand) ImageTag() string {
	return cmc.imageTag
}

func (cmc *ContainerManagerCommand) SetImageTag(imageTag string) *ContainerManagerCommand {
	cmc.imageTag = imageTag
	return cmc
}

func (cmc *ContainerManagerCommand) Repo() string {
	return cmc.repo
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

func (cmc *ContainerManagerCommand) RtDetails() *config.ArtifactoryDetails {
	return cmc.rtDetails
}

func (cmc *ContainerManagerCommand) SetRtDetails(rtDetails *config.ArtifactoryDetails) *ContainerManagerCommand {
	cmc.rtDetails = rtDetails
	return cmc
}
