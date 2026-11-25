# dockertidy

Small CLI that lists Docker containers and images with their creation time.

```text
go run .
go run . --match web --dry-run
```

Use `--match` to filter containers and images by name or repository. The
`--dry-run` flag makes the no-change behavior explicit. Docker must be
installed and available in `PATH`; this tool never removes resources.
