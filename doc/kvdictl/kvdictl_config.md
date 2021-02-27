## kvdictl config

Configuration commands

### Options

```
  -h, --help   help for config
```

### Options inherited from parent commands

```
  -C, --ca-file string         the CA certificate to use to verify the API certificate
  -c, --config string          configuration file (default "$HOME/.kvdi.yaml")
  -f, --filter string          a jmespath expression for filtering results (where applicable)
  -k, --insecure-skip-verify   skip verification of the API server certificate
  -o, --output string          the format to dump results in (default "json")
  -s, --server string          the address to the kvdi API server (default "https://127.0.0.1")
  -u, --user string            the username to use when authenticating against the API (default "admin")
```

### SEE ALSO

* [kvdictl](kvdictl.md)	 - 
* [kvdictl config client](kvdictl_config_client.md)	 - Client configuration commands
* [kvdictl config server](kvdictl_config_server.md)	 - Retrieve server configurations

###### Auto generated by spf13/cobra on 27-Feb-2021