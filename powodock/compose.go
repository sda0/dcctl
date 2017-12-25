package powodock

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"github.com/spf13/viper"
	"sync"
)

const YmlPattern = "docker-.+ya?ml"
const DefaultComposeFile = "docker-compose.yml"


type composeFiles struct {
	files []string
	pattern string
	defaultFile string
}
var foundComposeFiles *composeFiles
var once sync.Once

func GetInstance() *composeFiles {
	once.Do(func() {
		foundComposeFiles = &composeFiles{}
		foundComposeFiles.pattern = viper.GetString("composer_pattern")
		foundComposeFiles.defaultFile = viper.GetString("composer_default")
		foundComposeFiles.files = FindYmlFiles(viper.GetString("PW_DOCK"))
		sort.StringSlice(foundComposeFiles.files).Swap(sort.SearchStrings(foundComposeFiles.files, foundComposeFiles.defaultFile ), 0)
	})
	return foundComposeFiles
}

func FindYmlFiles(path string) (files []string) {
	filepath.Walk(path, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(viper.GetString("composer_pattern"), f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	return
}

func GetAllComposeFiles(sep string) string {
	return strings.Join(GetInstance().files, sep)
}

func GetComposeFilesAndServicesByArg(args []string) (composeFiles string, services string) {
	if len(args) == 0 {
		services = "nginx"
		composeFiles = " -f " + GetInstance().defaultFile
		return
	}

	composeFiles = " -f " + GetInstance().defaultFile
	for _, service := range args {
		if service == "all" {
			composeFiles = " -f " + GetAllComposeFiles(" -f ")
			services = ""
			return
		}

		filename := "docker-" + service + ".yml"

		if contains(GetInstance().files, filename) {
			composeFiles += " -f " + filename
		}
	}

	services = strings.Join(args, " ")

	return
}

func contains(strings []string, search string) bool {
	for _, value := range strings {
		if value == search {
			return true
		}
	}
	return false
}
