# Tomcat JMX attributes logging tool

A small utility which dumps Tomcat JMX attributes in CSV format.

## Prerequisite

It is required to enable JMX proxy servlet of the target Tomcat instance.
This can be done by the following steps:

(A) Add the following Java System variables to enable JMX features

  ```bash
    -Dcom.sun.management.jmxremote
    -Dcom.sun.management.jmxremote.local.only=false
    -Dcom.sun.management.jmxremote.authenticate=false
    -Dcom.sun.management.jmxremote.port=$JMX_PORT
    -Dcom.sun.management.jmxremote.rmi.port=$JMX_PORT
    -Djava.rmi.server.hostname=127.0.0.1
    -Dcom.sun.management.jmxremote.ssl=false
  ```

(B) Add JMX roles in the `tomcat-users.xml`

```xml
<?xml version="1.0" encoding="UTF-8"?>
<tomcat-users>
  <role rolename="manager-gui"/>
  <role rolename="manager-script"/>
  <role rolename="manager-jmx"/>
  <role rolename="manager-status"/>
  <role rolename="admin-gui"/>
  <role rolename="admin-script"/>
  <user
    username="tomcat"
    password="tomcat"
    roles="
      manager-gui,
      manager-script,
      manager-jmx,
      manager-status,
      admin-gui,
      admin-script"
  />
</tomcat-users>
```

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

JMX attributes are emitted to the standard output by default.

```console
$ jmx-logger -f jmxclient/config.yml

maxThreads,activeSessions
200,0
200,0
200,1
200,1
200,1
```
