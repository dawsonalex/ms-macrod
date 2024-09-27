# Macrod Server

ms-macrod provides a service for storing, tracking and calculating meal information.

## Tooling 

ms-macrod follows the [tools.go pattern](https://www.jvt.me/posts/2022/06/15/go-tools-dependency-management/)
so you can run tools imported there using `go run`.

## Building 

ms-macrod uses make as the main build tool. For a list of targets run `make help`:

```text
run                 Run the server using go run
build               build the server, default output to ${BIN_DIR}/${BIN_NAME}
rm-bin              remove the bin dir
help                Generate list of targets with descriptions
```