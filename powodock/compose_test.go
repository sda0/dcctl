package powodock

import (
	"fmt"
	"sort"
	"testing"
)

func TestFindYamlFiles(t *testing.T) {
	ymlList := FindYmlFiles("/home/dscheglov/powo/Dock/")
	sort.StringSlice(ymlList).Swap(sort.SearchStrings(ymlList, "powodock-compose.yml"), 0)

	fmt.Printf("%d: %s", len(ymlList), ymlList)
}
