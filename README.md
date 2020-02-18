# Tomcat JMX attributes logging tool

## Build

Built binary will be installed under `GOPATH`

```console
$ ./script/build.sh

$ which jmx-logger
/Users/iwauo/go/1.13.6/bin/jmx-logger
```

## Usage

### Sample configuration

```yaml
endpoint: http://localhost:8080/manager/jmxproxy

credential:
  user: tomcat
  pass: tomcat

output:
  interval: 5
  rows: 5
  useCRLF: true

columns:
- name: maxThreads
  type: Catalina:type=ThreadPool,name="http-apr-8080"
  attribute: maxThreads

- name: activeSessions
  type: Catalina:type=Manager,context=/,host=localhost
  attribute: activeSessions
```

### Run

```console
$ jmx-logger -f jmxclient/config.yml

maxThreads,activeSessions
200,0
200,0
200,1
200,1
200,1
```
