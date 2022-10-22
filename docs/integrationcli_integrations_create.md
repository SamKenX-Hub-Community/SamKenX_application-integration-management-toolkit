## integrationcli integrations create

Create an integration flow with a draft version

### Synopsis

Create an integration flow with a draft version

```
integrationcli integrations create [flags]
```

### Options

```
  -f, --file string   Integration flow instance
  -h, --help          help for create
  -n, --name string   Integration flow name
      --new           Set this flag to true, if draft version is to be created for a brand new integration
```

### Options inherited from parent commands

```
  -a, --account string   Path Service Account private key in JSON
      --disable-check    Disable check for newer versions
  -p, --proj string      Integration GCP Project name
  -r, --reg string       Integration region name
  -t, --token string     Google OAuth Token
```

### SEE ALSO

* [integrationcli integrations](integrationcli_integrations.md)	 - Manage integrations in a GCP project

###### Auto generated by spf13/cobra on 7-Jun-2022