Cloud Foundry Buildpack Management Plugin
==================

# Overview

The Cloud Foundry Buildpack Management Plugin provides a declarative way to configure system buildpacks 
in a Cloud Foundry installation. Define the desired state of buildpack configuration and the plugin
determines what buildpacks must be added, deleted, or updated.

# Installation

## Install from CLI

```
$ cf add-plugin-repo CF-Community http://plugins.cloudfoundry.org/
$ cf install-plugin 'buildpack-management' -r CF-Community
```
#### Install from binary

- Download the appropriate plugin binary from [releases](https://github.com/davidehringer/cf-buildpack-management-plugin/releases)
- Install the plugin: `$ cf install-plugin <binary>`

## Install From Source

[Go](http://golang.org/dl/) must be installed.

```
$ go get github.com/cloudfoundry/cli
$ go get github.com/davidehringer/cf-buildpack-management-plugin
$ cd $GOPATH/src/github.com/davidehringer/cf-buildpack-management-plugin
$ go build
$ cf install-plugin cf-buildpack-management-plugin
```

# Usage

```
$ cf configure-buildpacks PATH_TO_YAML_CONFIG_FILE [-dryRun] 
```

The provided path for the configuration file can be an absolute or relative path to a file.

## Configuration File Format

The file should have a map named "buildpacks" containing an array of buildpacks. The "filename" 
values for each buildpack can be an absolute or relative path to a file.

```
---
buildpacks:
- name: java
 position: 1
 enabled: true
 locked: false
 filename: java-buildpack-offline-v3.0.zip
- name: ruby_buildpack
 position: 2
 enabled: true
 locked: false
 filename: ruby_buildpack-cached-v1.3.0.zip
```

# Algorithm

Within the Cloud Controller APIs, there is not an explict way to uniquely identify a system buildpack that has been
added to the system. The Cloud Controller API does not expose a hash value of the buildpacks that could be used for comparison.
Buildpacks can be and are frequently renamed so the name is not a reliable identifier. Filename is the most reliable identifier
for a buildpack. The plugin uses filename as an identifier. While filename is not guaranteed to be unique, the likelihood that 
two different buildpacks configured in the Cloud Controller would use the same file is small (note: I believe the behavior of the
plugin in this case still correct although it will potentially do more work to get you to your desired state and certain commands such as renames will result in ignorable errors... create a test for this).

The plugin will determine what buildpacks need to be deleted, added, and renamed. Then the plugin will ensure correct positioning and other attributes as set for all buildpacks. The order that operations occur is:
* Delete no longer needed buildpacks
* Rename buildpacks whose name has changed
* Add buildpacks that don't exist
* Execute an update on all buildpacks to ensure all attributes are correct

# Example

Given the following configuration file.

```
---
buildpacks:
- name: php_buildpack
  position: 1
  enabled: true
  locked: false
  filename: php-buildpack-3.1.0.zip
- name: java_53f901b
  position: 2
  enabled: true
  locked: false
  filename: java-buildpack-offline-53f901b.zip
- name: nodejs_buildpack
  position: 3
  enabled: true
  locked: false
  filename: nodejs_buildpack-cached-v1.2.0.zip 
- name: go_buildpack
  position: 4
  enabled: true
  locked: false
  filename: go_buildpack-cached-v1.2.0.zip
- name: python_buildpack
  position: 5
  enabled: true
  locked: false
  filename: python_buildpack-cached-v1.2.0.zip 
- name: ruby_buildpack
  position: 6
  enabled: true
  locked: false
  filename: ruby_buildpack-cached-v1.3.0.zip
```
If the existing buildpacks configured look like:
```
$ cf buildpacks
Getting buildpacks...

buildpack                position   enabled   locked   filename   
java_53f901b             1          true      false    java-buildpack-offline-53f901b.zip   
java_buildpack_offline   2          true      false    java-buildpack-offline-v2.7.1.zip   
php_buildpack            3          true      false    php-buildpack-3.1.0.zip   
```

When you run 
```
$ ls -l
total 8146992
-rw-r--r--  1 dehringer  staff        814 Sep  3 21:09 config.yml
-rw-r-----@ 1 dehringer  staff  663443569 Sep  3 21:07 go_buildpack-cached-v1.2.0.zip
-rw-r-----@ 1 dehringer  staff  328395926 Sep  3 20:30 java-buildpack-offline-53f901b.zip
-rw-r-----@ 1 dehringer  staff  328395926 Sep  3 20:18 java-buildpack-offline-unlimited-crypto-3.0.zip
-rw-r-----@ 1 dehringer  staff  298273819 Sep  3 20:18 java-buildpack-offline-v2.7.1.zip
-rw-r-----@ 1 dehringer  staff  422081161 Sep  3 20:41 nodejs_buildpack-cached-v1.2.0.zip
-rw-r-----@ 1 dehringer  staff    6902368 Sep  3 20:14 php-buildpack-3.1.0.zip
-rw-r-----@ 1 dehringer  staff  471687722 Sep  3 21:07 php_buildpack-cached-v4.1.2.zip
-rw-r-----@ 1 dehringer  staff  685706721 Sep  3 21:07 python_buildpack-cached-v1.2.0.zip
-rw-r-----@ 1 dehringer  staff  966348817 Sep  3 21:07 ruby_buildpack-cached-v1.3.0.zip

$ cf configure-buildpacks config.yml --dryRun
The following actions are required to configure buildpacks:
	- delete buildpack named java_buildpack_offline
	- add buildpack named nodejs_buildpack
	- add buildpack named go_buildpack
	- add buildpack named python_buildpack
	- add buildpack named ruby_buildpack
	- update buildpack named ruby_buildpack
	- update buildpack named php_buildpack
	- update buildpack named java_53f901b
	- update buildpack named nodejs_buildpack
	- update buildpack named go_buildpack
	- update buildpack named python_buildpack
Dry run mode. No actions will be executed.
```
If `configure-buildpacks` was run without dryMode, then the stated actions would have been executed.

# Development

## Running Tests

```
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega
$ go test ./...
```