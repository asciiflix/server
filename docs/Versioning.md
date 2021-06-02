# Versioning

## Versioning
Our version numbers will look like the following:
``v.<release>.<milestone>.<issue>``

## Building with a version tag
To build a executable binary with versioning, use ``ldflags``.
```bash
go build -ldflags="-X 'github.com/asciiflix/server/config.Version=${VERSION}'"
```
The version of the running API will now be displayed on the main landing page.


## Merging feature branches/ pulling to master
For every completed issue, we will merge that respecting feature branch into develop.

After a certain milestone has been completed, the develop branch will be merged into master.
