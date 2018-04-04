package handlers

import (
	"os"
	"strings"
)

func getEnvironment() map[string]string {
	result := map[string]string{}

	for _, nvPair := range os.Environ() {
		elements := strings.SplitN(nvPair, "=", 2)

		k := elements[0]

		v := ""

		if 2 == len(elements) {
			v = elements[1]
		}

		result[k] = v
	}

	return result
}
