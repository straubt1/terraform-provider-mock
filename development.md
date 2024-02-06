# Development

## Getting started

```
go mod edit -module github.com/straubt1/terraform-provider-mock
```

```
export GOBIN="/Users/tstraub/go/bin"
```

### Delve

https://github.com/go-delve/delve/tree/master/Documentation/installation
```
xcode-select --install
sudo /usr/sbin/DevToolsSecurity -enable
```

```

dlv exec --accept-multiclient --continue --headless ./terraform-provider-mock -- -debug

```

```
export TF_CLI_CONFIG_FILE=~/.terraformrc
```


