# gcputil/project

## Usage

Many of the GPC client libraries still require `projectID` as an input parameter. For a long time the practice was to set that as an environment variable. GCP also not provides a metadata client to extract that at runtime. This utility exposes two simple functions to test for project ID being set in environment variable or in metadata.

#### Import

```shell
import "github.com/mchmarny/gcputil/pkg/project"
```

#### Usage

```shell
p, err := project.GetID()
```

Or alternatively, fail if not set

```shell
p := project.GetID()
```

Or gt the configured meta object

```shell
import "github.com/mchmarny/gcputil/pkg/meta"
```

...and get access to all the other methods

```shell
name := meta.GetClient().InstanceName()
```

