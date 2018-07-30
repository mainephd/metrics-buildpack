# Cloud Foundry Metrics Buildpack

A Cloud Foundry buildpack for running a sidercar process to collect metrics from an application. This buildpack is meant as a supply only buildpack. Which means you will need to use another Cloud Foundry [buildpack](http://docs.cloudfoundry.org/buildpacks/) in addition to this buildpack to run your application. In order for this solution to work properly the Cloud Foundry foundation must support [pushing an application multiple buildpacks](https://docs.cloudfoundry.org/buildpacks/use-multiple-buildpacks.html). Also the buildpack that will run the application needs support multi-buildpack pushes and at the time of this writing [the java-buildpack](https://github.com/cloudfoundry/java-buildpack) does not support this model (See conversation [here](https://cloudfoundry.slack.com/archives/C03F5ELTK/p1532549907000513)).

## Telegraf Usage

In order for Telegraf to be installed you **MUST** have a `telegraf.conf` in the root of the directory of the application you are pushing to Cloud Foundry. You **MAY** provide a `telegraf.env` file in the root of the directory of the application to set environment variables to be used in the `telegraf.conf` before the Telegraf process is started. The use case for this approach is set the `os.hostname` telegraf property to some derived value from the application `name` and `instance_index`.

* Example Telegraf Environment File

```bash
#!/bin/bash

app_name=$(echo $VCAP_APPLICATION | jq -r '.name')
app_instance=$(echo $VCAP_APPLICATION | jq -r '.instance_index')

if [[ $app_instance == "null" ]]; then 
    export APP_INSTANCE=$app_name
else
    export APP_INSTANCE=$app_name"-"$app_instance
fi
```

Then `$APP_INSTANCE` can be referred to in the `telegraf.conf` as below.

```conf
...
[agent]
  ...
  ## Override default hostname, if empty use os.Hostname()
  hostname = "$APP_INSTANCE"
  ## If set to true, do no set the "host" tag in the telegraf agent.
  omit_hostname = false
  ...
```