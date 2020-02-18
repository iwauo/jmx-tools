package jmxclient

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

func GetConfig(path string) (*Config, error) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	input, err := ioutil.ReadFile(absolutePath)
	if err != nil {
		return nil, err
	}
	config := Config{}
	err = yaml.Unmarshal(input, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

type Config struct {
	Endpoint string
	Output   struct {
		Interval int
		Rows     int
		UseCRLF  bool
	}
	Credential struct {
		User string
		Pass string
	}
	Columns []struct {
		Name      string
		Type      string
		Attribute string
	}
}

func Start(config Config) error {
	writer := csv.NewWriter(os.Stdout)
	writer.UseCRLF = config.Output.UseCRLF

	header := []string{}
	for _, column := range config.Columns {
		header = append(header, column.Name)
	}
	writer.Write(header)
	writer.Flush()

	duration := time.Duration(config.Output.Interval) * time.Second
	for row := 1; row <= config.Output.Rows; row++ {
		err := EmitRecord(writer, config)
		if err != nil {
			return err
		}
		time.Sleep(duration)
	}
	return nil
}

func EmitRecord(writer *csv.Writer, config Config) error {
	request, err := http.NewRequest("GET", config.Endpoint, nil)
	if err != nil {
		return err
	}
	request.SetBasicAuth(config.Credential.User, config.Credential.Pass)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	snapshot, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	record := []string{}

	for _, column := range config.Columns {
		typePattern := regexp.MustCompile(fmt.Sprintf(`Name: (?s)%s.*?\r\n\r\n`, column.Type))
		typeDesc := typePattern.Find(snapshot)
		attrPattern := regexp.MustCompile(fmt.Sprintf(`%s: (.*)`, column.Attribute))
		attrDesc := attrPattern.FindSubmatch(typeDesc)[1]
		record = append(record, strings.TrimSpace(string(attrDesc)))
	}

	writer.Write(record)
	writer.Flush()
	return nil
}
