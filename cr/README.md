# gcputil/cr

This utility enables service to service calls between Cloud Run services where the target service requires requires authentication. This utility will append default service accounts identity authorization header (`Authorization: Bearer ***`) to the referenced HTTP request.

## Import

```shell
import "github.com/mchmarny/gcputil/cr"
```

## Usage

```shell
req, err := http.NewRequest("GET", url, nil)
if err != nil {
    log.Fatal(err)
}
err = cr.AuthorizeRequest(req, url)
if err != nil {
    log.Fatal(err)
}
resp, err := http.DefaultClient.Do(req)
```
