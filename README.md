# ToDo Server

This is an example service created with the aim of testing out different
frameworks and build processes.

## Tooling 

Todo Server follows the [tools.go pattern](https://www.jvt.me/posts/2022/06/15/go-tools-dependency-management/)
so you can run tools imported there using `go run`.

## Building 

Todo Server uses make as the main build tool. For a list of targets run `make help`:

```text
run                 Run the server using go run
build               build the server, default output to ${BIN_DIR}/${BIN_NAME}
rm-bin              remove the bin dir
help                Generate list of targets with descriptions
```