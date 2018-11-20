package coordinator

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/caicloud/cyclone/pkg/workflow/common"
)

func getPodName() string {
	return os.Getenv(common.EnvStagePodName)
}

func getWorkflowrunName() string {
	return os.Getenv(common.EnvWorkflowrunName)
}

func getStageName() string {
	return os.Getenv(common.EnvStageName)
}

func getWorkloadContainer() string {
	return os.Getenv(common.EnvWorkloadContainerName)
}

func getNamespace() string {
	n := os.Getenv(common.EnvNamespace)
	if n == "" {
		return "default"
	}

	return n
}

func createDirectory(dirName string) bool {
	src, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirName, 0755)
		if errDir != nil {
			log.Errorf("mkdir %s failed: %v", dirName, errDir)
			panic(errDir)
		}
		return true
	}

	if src.Mode().IsRegular() {
		log.Error(dirName, "already exist as a file!")
		return false
	}

	return false
}

// refineContainerID strips the 'docker://' prefix from k8s ContainerID string
func refineContainerID(id string) string {
	schemeIndex := strings.Index(id, "://")
	if schemeIndex == -1 {
		return id
	}
	return id[schemeIndex+3:]
}
