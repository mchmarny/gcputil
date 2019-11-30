# gcputil

Collection of common golang GCP development patterns

## Usage

* [Project](project/) (projectID as well as other meta client helpers)
* [Metric](metric/) (custom Stackdriver metric helper)
* [Cloud Run](cr/) (Inter-Cloud Run service auth helper)
* [Metadata](meta/) (metadata service helper)
* [VM](vm/) (GCE VM helper)
* [EnvVars](env/) (environment variables helper)

> The general consensus of the Go community nowadays is that the pkg/ directory is a useless abstraction/indirection. So, this projects keeps its packages in top-level of the repository, based on their primary functionality.