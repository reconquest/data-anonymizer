package main

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/reconquest/pkg/log"

	"github.com/docopt/docopt-go"
	"github.com/zazab/zhash"
)

var (
	version = "[manual build]"
	usage   = "data-anonymizer " + version + `


Usage:
  data-anonymizer [options]
  data-anonymizer -h | --help
  data-anonymizer --version

Options:
  -h --help           Show this screen.
  -c --config <path>  Path to config. [default: data-anonymizer.yaml]
  --version           Show version.
`
)

func main() {
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		panic(err)
	}

	generators := map[string]func(*zhash.Hash, ...string){
		"name":        GenerateName,
		"company":     GenerateCompany,
		"email":       GenerateEmail,
		"int_string":  GenerateIntString,
		"city":        GenerateCity,
		"zip":         GenerateZip,
		"sen_license": GenerateSenLicense,
		"at_license":  GenerateAtLicense,
		"state":       GenerateState,
	}

	config, err := getConfig(args["--config"].(string))
	if err != nil {
		log.Fatalf(err, "unable to read config")
	}

	var units []map[string]interface{}
	err = json.NewDecoder(os.Stdin).Decode(&units)
	if err != nil {
		log.Fatal(err)
	}

	var result []map[string]interface{}
	for _, unit := range units {
		hash := zhash.NewHashPtr()
		hash.SetRoot(unit)
		for key, generator := range config {
			path := strings.Split(key, ".")

			value := hash.Get(path...)
			if value == nil {
				continue
			}

			fn, ok := generators[generator]
			if !ok {
				log.Fatalf(nil, "no such generator: %s", generator)
			}

			fn(hash, path...)
		}

		result = append(result, hash.GetRoot())
	}

	err = json.NewEncoder(os.Stdout).Encode(result)
	if err != nil {
		log.Fatal(err)
	}
}
