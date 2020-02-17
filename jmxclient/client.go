package jmxclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"gopkg.in/yaml.v2"
)

var configData = `
Credential:
  User: tomcat
  Pass: tomcat
Columns:
- Name: maxThreads
  Type: Catalina:type=ThreadPool,name="http-apr-8080"
  Attribute: maxThreads
`

func GetConfig() (*Config, error) {
	config := Config{}

	err := yaml.Unmarshal([]byte(configData), &config)
	if err != nil {
		return nil, err
	}
}

type Config struct {
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

func Start() (string, error) {
	jmxEndpoint := "http://localhost:8080/manager/jmxproxy"
	objectType := regexp.QuoteMeta(`Catalina:type=ThreadPool,name="http-apr-8080"`)
	attribute := regexp.QuoteMeta(`maxThreads`)
	//endpoint := fmt.Sprintf("%s/?get=%s&att=%s&key=value", jmxEndpoint, objectType, attribute)
	request, err := http.NewRequest("GET", jmxEndpoint, nil)
	if err != nil {
		return "", err
	}
	request.SetBasicAuth("tomcat", "tomcat")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	snapshot, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	typePattern := regexp.MustCompile(fmt.Sprintf(`Name: (?s)%s.*?\r\n\r\n`, objectType))
	typeDesc := typePattern.Find(snapshot)
	attrPattern := regexp.MustCompile(fmt.Sprintf(`%s: (.*)`, attribute))
	attrDesc := attrPattern.FindSubmatch(typeDesc)
	return string(attrDesc[1]), nil
}
