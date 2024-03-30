package storage

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/samber/lo"
	"gopkg.in/yaml.v2"

	"github.com/bentohset/gnm/internal/logger"
	"github.com/bentohset/gnm/internal/model"
)

type YAMLHostWrapper struct {
	Host model.Host `yaml:"host"`
}

type YAMLStorage struct {
	innerStorage map[string]YAMLHostWrapper
	fsDataPath   string
	logger       *logger.AppLogger
}

const (
	hostsFile = "hosts.yaml"
	idEmpty   = 0
)

func NewYAML(appFolder string, logger *logger.AppLogger) (*YAMLStorage, error) {
	fsDataPath := path.Join(appFolder, hostsFile)

	yamlFile := YAMLStorage{
		innerStorage: make(map[string]YAMLHostWrapper),
		fsDataPath:   fsDataPath,
		logger:       logger,
	}

	// check if yaml exists
	fileData, err := os.ReadFile(fsDataPath)
	if err != nil {
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			// if yaml doesnt exist, return new state
			return &yamlFile, nil
		}
		return nil, err
	}

	// yaml exists, transfer yaml file to map
	var yamlHosts []YAMLHostWrapper
	err = yaml.Unmarshal(fileData, &yamlHosts)
	if err != nil {
		return nil, err
	}

	for _, wrapped := range yamlHosts {
		yamlFile.innerStorage[wrapped.Host.Alias] = wrapped
	}

	return &yamlFile, nil
}

func (s *YAMLStorage) saveToDisk() error {
	// get map
	mapValues := lo.Values(s.innerStorage)

	result, err := yaml.Marshal(mapValues)
	if err != nil {
		s.logger.Error("[LOGGER]: error converting map to yaml")
		return err
	}

	// write file to filepath
	err = os.WriteFile(s.fsDataPath, result, 0o600)
	if err != nil {
		s.logger.Error("[LOGGER]: error writing yaml to file")
		panic(err)
	}

	return nil
}

func (s *YAMLStorage) Save(host model.Host) (model.Host, error) {
	_, exists := s.innerStorage[host.Alias]
	if exists {
		s.logger.Error("[LOGGER]: Save host alias exists already")
		return host, fmt.Errorf("host alias exists already")
	}
	s.innerStorage[host.Alias] = YAMLHostWrapper{host}

	err := s.saveToDisk()
	if err != nil {
		s.logger.Error("[LOGGER]: Save failed to save to disk")
		return host, fmt.Errorf("failed to save to disk")
	}

	return host, err
}

func (s *YAMLStorage) Delete(hostAlias string) error {
	delete(s.innerStorage, hostAlias)

	err := s.saveToDisk()
	if err != nil {
		s.logger.Error("[LOGGER]: Delete failed to save to disk")
		return fmt.Errorf("failed to save to disk")
	}
	return err
}

func (s *YAMLStorage) Get(hostAlias string) (model.Host, error) {
	found, ok := s.innerStorage[hostAlias]

	if !ok {
		s.logger.Error("[LOGGER]: Get failed to retrieve alias")
		return model.Host{}, fmt.Errorf("failed to retrieve alias")
	}

	return found.Host, nil
}

// debug print
func (s *YAMLStorage) Print() {
	mp := s.innerStorage

	if len(mp) == 0 {
		fmt.Println("Storage empty")
	}

	for alias, host := range mp {
		fmt.Println(alias)
		fmt.Println("  description: \t" + host.Host.Description)
		fmt.Println("  hostname: \t" + host.Host.HostName)
		fmt.Println("  user: \t" + host.Host.User)
		fmt.Println("  key path: \t" + host.Host.PrivateKeyPath)
	}
}
