package jsonconfig

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

var (
	confFilename string
	runtimeConf  Conf
)

type Conf struct {
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
	Version   string                 `json:"version"`
	Config    map[string]interface{} `json:"config"`
}

// initialize the configuration file based on file name.
// - new file will be created if not exists;
func InitConf(confFile string) (err error) {
	// record the filename for later config saving
	confFilename = confFile

	// read the whole file at once
	confStr, err := ioutil.ReadFile(confFile)
	if err != nil {
		if os.IsNotExist(err) {
			return createEmptyConfig()
		}

		return
	}

	// convert to runtime config
	json.Unmarshal([]byte(confStr), &runtimeConf)

	return nil
}

func createEmptyConfig() error {
	// create empty file with default
	runtimeConf = Conf{}
	runtimeConf.CreatedAt = time.Now()
	runtimeConf.UpdatedAt = time.Now()
	runtimeConf.Version = "0.1"
	runtimeConf.Config = make(map[string]interface{})
	return saveToFile()
}

// save all runtime config data to file
func saveToFile() (err error) {
	confStr, err := json.Marshal(runtimeConf)
	if err != nil {
		return
	}

	// write the whole body at once
	runtimeConf.UpdatedAt = time.Now()
	err = ioutil.WriteFile(confFilename, confStr, 0644)
	if err != nil {
		return
	}

	return nil
}

func Set(param string, value interface{}) error {
	runtimeConf.Config[param] = value

	return saveToFile()
}

func Get(param string) interface{} {
	return runtimeConf.Config[param]
}

func GetRuntimeData() Conf {
	return runtimeConf
}
