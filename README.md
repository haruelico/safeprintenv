# safeprintenv
This is cli tool like `printenv` command but filter sensitive environment variable (e.g. AWS_SECRET_ACCESS_KEY) from output.

## Install
TBD

## Configure

## Usage

```shell
$ SENSITIVE_VAR=aaa safeprintenv | grep SENSITIVE_VAR
$  # no match
```