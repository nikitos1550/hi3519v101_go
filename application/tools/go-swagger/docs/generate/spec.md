# Generate a spec from source code

```
Usage:
  swagger [OPTIONS] generate spec [spec-OPTIONS]

generate a swagger spec document from a go application

Application Options:
  -q, --quiet                  silence logs
      --log-output=LOG-FILE    redirect logs to file

Help Options:
  -h, --help                   Show this help message

[spec command options]
      -b, --base-path=         the base path to use (default: .)
      -t, --tags=              build tags
      -m, --scan-models        includes models that were annotated with 'swagger:model'
          --compact            when present, doesn't prettify the json
      -o, --output=            the file to write to
      -i, --input=             the file to use as input
      -c, --include=           include packages matching pattern
      -x, --exclude=           exclude packages matching pattern
          --include-tag=       include routes having specified tags (can be specified many times)
          --exclude-tag=       exclude routes having specified tags (can be specified many times)
```

See code annotation rules [here](../use/spec.md)
