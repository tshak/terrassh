gox -osarch="darwin/amd64 linux/amd64 windows/amd64" -output="release/{{.OS}}_{{.Arch}}"
pushd release
gzip *
popd