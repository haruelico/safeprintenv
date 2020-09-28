# safeprintenv
This is cli tool like `printenv` command but filter sensitive environment variable (e.g. AWS_SECRET_ACCESS_KEY) from output.

## Install
Download latest binary from Release page.

## Configure

Create `~/.config/safeprintenv/config.toml` and write some settings.

It is example config.

```toml
[SensitiveList]
keys = ['ADDITIONAL_KEY']
```

See also `config.example.toml` .

## Usage

```shell
$ SENSITIVE_VAR=aaa safeprintenv | grep SENSITIVE_VAR
$  # no match
$ SENSITIVE_VAR=aaa safeprintenv --unsafe-all | grep SENSITIVE_VAR
SENSITIVE_VAR=aaa
$
```
