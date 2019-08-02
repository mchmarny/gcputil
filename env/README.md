# gcputil/conf

## Import


```shell
import "github.com/mchmarny/gcputil/conf"
```

## Usage

Parse string environment variable

```shell
name := conf.MustGetEnvVar("ENV_VAR_NAME", "default-value")
```

Parse int environment variable

```shell
name := conf.MustGetIntEnvVar("HTTP_PORT", 8080)
```