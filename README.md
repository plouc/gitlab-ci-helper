# gitlab-ci-helper

[![Build Status](https://travis-ci.org/rande/gitlab-ci-helper.png?branch=master)](https://travis-ci.org/rande/gitlab-ci-helper)
[![Coverage Status](https://coveralls.io/repos/github/rande/gitlab-ci-helper/badge.svg?branch=master)](https://coveralls.io/github/rande/gitlab-ci-helper?branch=master)
[![GoDoc](https://godoc.org/github.com/rande/gitlab-ci-helper?status.svg)](https://godoc.org/github.com/rande/gitlab-ci-helper)
[![GitHub license](https://img.shields.io/github/license/rande/gitlab-ci-helper.svg)](https://github.com/rande/gitlab-ci-helper/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/rande/gitlab-ci-helper)](https://goreportcard.com/report/github.com/rande/gitlab-ci-helper)
[![GitHub issues](https://img.shields.io/github/issues/rande/gitlab-ci-helper.svg)](https://github.com/rande/gitlab-ci-helper/issues)

This tool provides a binary cli to execute common commands inside a gitlab's job.

[installation](#installation) | [commands](#commands)

## Installation

**gitlab-ci-helper** is a single binary with no external dependencies, released for several platforms.
Go to the [releases page](https://github.com/rande/gitlab-ci-helper/releases),
download the package for your OS, and copy the binary to somewhere on your PATH.
Please make sure to rename the binary to `gitlab-ci-helper` and make it executable.

As it primarily targets GitLab CI environment, a simple curl will also do :).

A release also includes checksums for various build flavors,
as a build has access to sensible information (GitLab tokens, AWS credentials…),
please use them.

## Commands

- [gitlab-ci-helper aws](#gitlab-ci-helper-aws)	*aws commands*
- [gitlab-ci-helper aws s3-archive](#gitlab-ci-helper-aws-s3-archive)	*Send archive to a S3 bucket*
- [gitlab-ci-helper aws s3-extract](#gitlab-ci-helper-aws-s3-extract)	*Extract archive from a S3 bucket*
- [gitlab-ci-helper doc](#gitlab-ci-helper-doc)	*Generate CLI documentation in markdown format*
- [gitlab-ci-helper download-artifacts](#gitlab-ci-helper-download-artifacts)	*Download job artifacts and extract them to specified path if provided*
- [gitlab-ci-helper dump-meta](#gitlab-ci-helper-dump-meta)	*Dump meta information about CI in a file*
- [gitlab-ci-helper dump-revision](#gitlab-ci-helper-dump-revision)	*Dump a REVISION file with the current sha1*
- [gitlab-ci-helper environments](#gitlab-ci-helper-environments)	*Project environments commands*
- [gitlab-ci-helper environments add](#gitlab-ci-helper-environments-add)	*Add a new project environment*
- [gitlab-ci-helper environments list](#gitlab-ci-helper-environments-list)	*List project's environments*
- [gitlab-ci-helper environments remove](#gitlab-ci-helper-environments-remove)	*Remove a project environment*
- [gitlab-ci-helper hipchat](#gitlab-ci-helper-hipchat)	*hipchat commands*
- [gitlab-ci-helper hipchat status](#gitlab-ci-helper-hipchat-status)	*Sends job status to hipchat*
- [gitlab-ci-helper list-projects](#gitlab-ci-helper-list-projects)	*List available projects*
- [gitlab-ci-helper slack](#gitlab-ci-helper-slack)	*slack commands*
- [gitlab-ci-helper slack status](#gitlab-ci-helper-slack-status)	*Sends job status to slack*
- [gitlab-ci-helper version](#gitlab-ci-helper-version)	*Print the version number of gitlab-ci-helper*



### gitlab-ci-helper aws

aws commands

#### Synopsis

aws commands

#### Options

```
      --aws-endpoint string   aws endpoint
                              default to environment variable: AWS_ENDPOINT
                              or 'x.aws.endpoint' path in config file if exists
      --aws-profile string    aws profile
                              default to environment variable: AWS_PROFILE
                              or 'x.aws.profile' path in config file if exists
      --aws-region string     aws region
                              default to environment variable: AWS_REGION
                              or 'x.aws.region' path in config file if exists
  -h, --help                  help for aws
```

#### Options inherited from parent commands

```
      --api-path string   gitlab api path
                          default to env var: GITLAB_API_PATH
                          or 'api_path' key in config file if exists,
                          otherwise: '/api/v4'
  -c, --config string     config file (default is .ci.yaml)
      --host string       gitlab host
                          default to env var: GITLAB_HOST
                          or 'host' key in config file if exists
      --no-color          disable color output
      --token string      gitlab token
                          default to env var: GITLAB_TOKEN
                          or 'token' key in config file if exists
  -v, --verbose           enable verbose mode
```

#### See also

- [gitlab-ci-helper aws s3-archive](#gitlab-ci-helper-aws-s3-archive)	*Send archive to a S3 bucket*
- [gitlab-ci-helper aws s3-extract](#gitlab-ci-helper-aws-s3-extract)	*Extract archive from a S3 bucket*



### gitlab-ci-helper aws s3-archive

Send archive to a S3 bucket

#### Synopsis

Send archive to a S3 bucket

```
gitlab-ci-helper aws s3-archive [flags]
```

#### Options

```
  -h, --help   help for s3-archive
```

#### Options inherited from parent commands

```
      --api-path string       gitlab api path
                              default to env var: GITLAB_API_PATH
                              or 'api_path' key in config file if exists,
                              otherwise: '/api/v4'
      --aws-endpoint string   aws endpoint
                              default to environment variable: AWS_ENDPOINT
                              or 'x.aws.endpoint' path in config file if exists
      --aws-profile string    aws profile
                              default to environment variable: AWS_PROFILE
                              or 'x.aws.profile' path in config file if exists
      --aws-region string     aws region
                              default to environment variable: AWS_REGION
                              or 'x.aws.region' path in config file if exists
  -c, --config string         config file (default is .ci.yaml)
      --host string           gitlab host
                              default to env var: GITLAB_HOST
                              or 'host' key in config file if exists
      --no-color              disable color output
      --token string          gitlab token
                              default to env var: GITLAB_TOKEN
                              or 'token' key in config file if exists
  -v, --verbose               enable verbose mode
```

#### See also

- [gitlab-ci-helper aws](#gitlab-ci-helper-aws)	*aws commands*



### gitlab-ci-helper aws s3-extract

Extract archive from a S3 bucket

#### Synopsis

Extract archive from a S3 bucket

```
gitlab-ci-helper aws s3-extract [flags]
```

#### Options

```
  -h, --help   help for s3-extract
```

#### Options inherited from parent commands

```
      --api-path string       gitlab api path
                              default to env var: GITLAB_API_PATH
                              or 'api_path' key in config file if exists,
                              otherwise: '/api/v4'
      --aws-endpoint string   aws endpoint
                              default to environment variable: AWS_ENDPOINT
                              or 'x.aws.endpoint' path in config file if exists
      --aws-profile string    aws profile
                              default to environment variable: AWS_PROFILE
                              or 'x.aws.profile' path in config file if exists
      --aws-region string     aws region
                              default to environment variable: AWS_REGION
                              or 'x.aws.region' path in config file if exists
  -c, --config string         config file (default is .ci.yaml)
      --host string           gitlab host
                              default to env var: GITLAB_HOST
                              or 'host' key in config file if exists
      --no-color              disable color output
      --token string          gitlab token
                              default to env var: GITLAB_TOKEN
                              or 'token' key in config file if exists
  -v, --verbose               enable verbose mode
```

#### See also

- [gitlab-ci-helper aws](#gitlab-ci-helper-aws)	*aws commands*



### gitlab-ci-helper doc

Generate CLI documentation in markdown format

#### Synopsis

Generate CLI documentation in markdown format

```
gitlab-ci-helper doc [flags]
```

#### Options

```
  -h, --help   help for doc
```

#### Options inherited from parent commands

```
      --api-path string   gitlab api path
                          default to env var: GITLAB_API_PATH
                          or 'api_path' key in config file if exists,
                          otherwise: '/api/v4'
  -c, --config string     config file (default is .ci.yaml)
      --host string       gitlab host
                          default to env var: GITLAB_HOST
                          or 'host' key in config file if exists
      --no-color          disable color output
      --token string      gitlab token
                          default to env var: GITLAB_TOKEN
                          or 'token' key in config file if exists
  -v, --verbose           enable verbose mode
```



### gitlab-ci-helper download-artifacts

Download job artifacts and extract them to specified path if provided

#### Synopsis

Download job artifacts and extract them to specified path if provided

```
gitlab-ci-helper download-artifacts [flags]
```

#### Options

```
  -h, --help   help for download-artifacts
```

#### Options inherited from parent commands

```
      --api-path string   gitlab api path
                          default to env var: GITLAB_API_PATH
                          or 'api_path' key in config file if exists,
                          otherwise: '/api/v4'
  -c, --config string     config file (default is .ci.yaml)
      --host string       gitlab host
                          default to env var: GITLAB_HOST
                          or 'host' key in config file if exists
      --no-color          disable color output
      --token string      gitlab token
                          default to env var: GITLAB_TOKEN
                          or 'token' key in config file if exists
  -v, --verbose           enable verbose mode
```



### gitlab-ci-helper dump-meta

Dump meta information about CI in a file

#### Synopsis

Dump meta information about CI in a file

```
gitlab-ci-helper dump-meta [flags]
```

#### Options

```
  -f, --file string   The meta file
                      can also be defined with env var: META_FILE (default "ci.json")
  -h, --help          help for dump-meta
```

#### Options inherited from parent commands

```
      --api-path string   gitlab api path
                          default to env var: GITLAB_API_PATH
                          or 'api_path' key in config file if exists,
                          otherwise: '/api/v4'
  -c, --config string     config file (default is .ci.yaml)
      --host string       gitlab host
                          default to env var: GITLAB_HOST
                          or 'host' key in config file if exists
      --no-color          disable color output
      --token string      gitlab token
                          default to env var: GITLAB_TOKEN
                          or 'token' key in config file if exists
  -v, --verbose           enable verbose mode
```



### gitlab-ci-helper dump-revision

Dump a REVISION file with the current sha1

#### Synopsis

Dump a REVISION file with the current sha1

```
gitlab-ci-helper dump-revision [flags]
```

#### Options

```
  -f, --file string   The revision file
                      can also be defined with env var: REVISION_FILE (default "REVISION")
  -h, --help          help for dump-revision
  -r, --ref string    The sha1, default to env var: CI_COMMIT_SHA
```

#### Options inherited from parent commands

```
      --api-path string   gitlab api path
                          default to env var: GITLAB_API_PATH
                          or 'api_path' key in config file if exists,
                          otherwise: '/api/v4'
  -c, --config string     config file (default is .ci.yaml)
      --host string       gitlab host
                          default to env var: GITLAB_HOST
                          or 'host' key in config file if exists
      --no-color          disable color output
      --token string      gitlab token
                          default to env var: GITLAB_TOKEN
                          or 'token' key in config file if exists
  -v, --verbose           enable verbose mode
```



### gitlab-ci-helper environments

Project environments commands

#### Synopsis

Project environments commands

#### Options

```
  -h, --help                help for environments
  -p, --project-id string   project id
                            default to environment variable: CI_PROJECT_ID
```

#### Options inherited from parent commands

```
      --api-path string   gitlab api path
                          default to env var: GITLAB_API_PATH
                          or 'api_path' key in config file if exists,
                          otherwise: '/api/v4'
  -c, --config string     config file (default is .ci.yaml)
      --host string       gitlab host
                          default to env var: GITLAB_HOST
                          or 'host' key in config file if exists
      --no-color          disable color output
      --token string      gitlab token
                          default to env var: GITLAB_TOKEN
                          or 'token' key in config file if exists
  -v, --verbose           enable verbose mode
```

#### See also

- [gitlab-ci-helper environments add](#gitlab-ci-helper-environments-add)	*Add a new project environment*
- [gitlab-ci-helper environments list](#gitlab-ci-helper-environments-list)	*List project's environments*
- [gitlab-ci-helper environments remove](#gitlab-ci-helper-environments-remove)	*Remove a project environment*



### gitlab-ci-helper environments add

Add a new project environment

#### Synopsis

Add a new project environment

```
gitlab-ci-helper environments add NAME URL [flags]
```

#### Options

```
  -h, --help   help for add
```

#### Options inherited from parent commands

```
      --api-path string     gitlab api path
                            default to env var: GITLAB_API_PATH
                            or 'api_path' key in config file if exists,
                            otherwise: '/api/v4'
  -c, --config string       config file (default is .ci.yaml)
      --host string         gitlab host
                            default to env var: GITLAB_HOST
                            or 'host' key in config file if exists
      --no-color            disable color output
  -p, --project-id string   project id
                            default to environment variable: CI_PROJECT_ID
      --token string        gitlab token
                            default to env var: GITLAB_TOKEN
                            or 'token' key in config file if exists
  -v, --verbose             enable verbose mode
```

#### See also

- [gitlab-ci-helper environments](#gitlab-ci-helper-environments)	*Project environments commands*



### gitlab-ci-helper environments list

List project's environments

#### Synopsis

List project's environments

```
gitlab-ci-helper environments list [flags]
```

#### Options

```
  -h, --help   help for list
```

#### Options inherited from parent commands

```
      --api-path string     gitlab api path
                            default to env var: GITLAB_API_PATH
                            or 'api_path' key in config file if exists,
                            otherwise: '/api/v4'
  -c, --config string       config file (default is .ci.yaml)
      --host string         gitlab host
                            default to env var: GITLAB_HOST
                            or 'host' key in config file if exists
      --no-color            disable color output
  -p, --project-id string   project id
                            default to environment variable: CI_PROJECT_ID
      --token string        gitlab token
                            default to env var: GITLAB_TOKEN
                            or 'token' key in config file if exists
  -v, --verbose             enable verbose mode
```

#### See also

- [gitlab-ci-helper environments](#gitlab-ci-helper-environments)	*Project environments commands*



### gitlab-ci-helper environments remove

Remove a project environment

#### Synopsis

Remove a project environment

```
gitlab-ci-helper environments remove ENV_ID [flags]
```

#### Options

```
  -h, --help   help for remove
```

#### Options inherited from parent commands

```
      --api-path string     gitlab api path
                            default to env var: GITLAB_API_PATH
                            or 'api_path' key in config file if exists,
                            otherwise: '/api/v4'
  -c, --config string       config file (default is .ci.yaml)
      --host string         gitlab host
                            default to env var: GITLAB_HOST
                            or 'host' key in config file if exists
      --no-color            disable color output
  -p, --project-id string   project id
                            default to environment variable: CI_PROJECT_ID
      --token string        gitlab token
                            default to env var: GITLAB_TOKEN
                            or 'token' key in config file if exists
  -v, --verbose             enable verbose mode
```

#### See also

- [gitlab-ci-helper environments](#gitlab-ci-helper-environments)	*Project environments commands*



### gitlab-ci-helper hipchat

hipchat commands

#### Synopsis

hipchat commands

#### Options

```
  -h, --help                   help for hipchat
      --hipchat-host string    hipchat host
                               default to environment variable: HIPCHAT_HOST
                               or 'x.hipchat.host' path in config file if exists
      --hipchat-token string   hipchat token
                               default to environment variable: HIPCHAT_TOKEN
                               or 'x.hipchat.token' path in config file if exists
```

#### Options inherited from parent commands

```
      --api-path string   gitlab api path
                          default to env var: GITLAB_API_PATH
                          or 'api_path' key in config file if exists,
                          otherwise: '/api/v4'
  -c, --config string     config file (default is .ci.yaml)
      --host string       gitlab host
                          default to env var: GITLAB_HOST
                          or 'host' key in config file if exists
      --no-color          disable color output
      --token string      gitlab token
                          default to env var: GITLAB_TOKEN
                          or 'token' key in config file if exists
  -v, --verbose           enable verbose mode
```

#### See also

- [gitlab-ci-helper hipchat status](#gitlab-ci-helper-hipchat-status)	*Sends job status to hipchat*



### gitlab-ci-helper hipchat status

Sends job status to hipchat

#### Synopsis

Sends job status to hipchat

```
gitlab-ci-helper hipchat status ROOM MESSAGE [flags]
```

#### Options

```
  -h, --help   help for status
```

#### Options inherited from parent commands

```
      --api-path string        gitlab api path
                               default to env var: GITLAB_API_PATH
                               or 'api_path' key in config file if exists,
                               otherwise: '/api/v4'
  -c, --config string          config file (default is .ci.yaml)
      --hipchat-host string    hipchat host
                               default to environment variable: HIPCHAT_HOST
                               or 'x.hipchat.host' path in config file if exists
      --hipchat-token string   hipchat token
                               default to environment variable: HIPCHAT_TOKEN
                               or 'x.hipchat.token' path in config file if exists
      --host string            gitlab host
                               default to env var: GITLAB_HOST
                               or 'host' key in config file if exists
      --no-color               disable color output
      --token string           gitlab token
                               default to env var: GITLAB_TOKEN
                               or 'token' key in config file if exists
  -v, --verbose                enable verbose mode
```

#### See also

- [gitlab-ci-helper hipchat](#gitlab-ci-helper-hipchat)	*hipchat commands*



### gitlab-ci-helper list-projects

List available projects

#### Synopsis

List available projects

```
gitlab-ci-helper list-projects [flags]
```

#### Options

```
  -h, --help   help for list-projects
```

#### Options inherited from parent commands

```
      --api-path string   gitlab api path
                          default to env var: GITLAB_API_PATH
                          or 'api_path' key in config file if exists,
                          otherwise: '/api/v4'
  -c, --config string     config file (default is .ci.yaml)
      --host string       gitlab host
                          default to env var: GITLAB_HOST
                          or 'host' key in config file if exists
      --no-color          disable color output
      --token string      gitlab token
                          default to env var: GITLAB_TOKEN
                          or 'token' key in config file if exists
  -v, --verbose           enable verbose mode
```



### gitlab-ci-helper slack

slack commands

#### Synopsis

slack commands

#### Options

```
  -h, --help                 help for slack
      --slack-token string   slack token
                             default to environment variable: SLACK_TOKEN
                             or 'x.slack.token' path in config file if exists
```

#### Options inherited from parent commands

```
      --api-path string   gitlab api path
                          default to env var: GITLAB_API_PATH
                          or 'api_path' key in config file if exists,
                          otherwise: '/api/v4'
  -c, --config string     config file (default is .ci.yaml)
      --host string       gitlab host
                          default to env var: GITLAB_HOST
                          or 'host' key in config file if exists
      --no-color          disable color output
      --token string      gitlab token
                          default to env var: GITLAB_TOKEN
                          or 'token' key in config file if exists
  -v, --verbose           enable verbose mode
```

#### See also

- [gitlab-ci-helper slack status](#gitlab-ci-helper-slack-status)	*Sends job status to slack*



### gitlab-ci-helper slack status

Sends job status to slack

#### Synopsis

Sends job status to slack

```
gitlab-ci-helper slack status [flags]
```

#### Options

```
  -h, --help   help for status
```

#### Options inherited from parent commands

```
      --api-path string      gitlab api path
                             default to env var: GITLAB_API_PATH
                             or 'api_path' key in config file if exists,
                             otherwise: '/api/v4'
  -c, --config string        config file (default is .ci.yaml)
      --host string          gitlab host
                             default to env var: GITLAB_HOST
                             or 'host' key in config file if exists
      --no-color             disable color output
      --slack-token string   slack token
                             default to environment variable: SLACK_TOKEN
                             or 'x.slack.token' path in config file if exists
      --token string         gitlab token
                             default to env var: GITLAB_TOKEN
                             or 'token' key in config file if exists
  -v, --verbose              enable verbose mode
```

#### See also

- [gitlab-ci-helper slack](#gitlab-ci-helper-slack)	*slack commands*



### gitlab-ci-helper version

Print the version number of gitlab-ci-helper

#### Synopsis

Print the version number of gitlab-ci-helper

```
gitlab-ci-helper version [flags]
```

#### Options

```
  -h, --help   help for version
```

#### Options inherited from parent commands

```
      --api-path string   gitlab api path
                          default to env var: GITLAB_API_PATH
                          or 'api_path' key in config file if exists,
                          otherwise: '/api/v4'
  -c, --config string     config file (default is .ci.yaml)
      --host string       gitlab host
                          default to env var: GITLAB_HOST
                          or 'host' key in config file if exists
      --no-color          disable color output
      --token string      gitlab token
                          default to env var: GITLAB_TOKEN
                          or 'token' key in config file if exists
  -v, --verbose           enable verbose mode
```



