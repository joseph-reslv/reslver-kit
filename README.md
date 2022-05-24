# Reslver Kit

> Related Confluence Page: [Model Language](https://resolve.atlassian.net/wiki/spaces/P/pages/2165637181)

---

## Usages

### reslver-kit --help

```
NAME:
   reslver-kit - Reslver Toolkit

USAGE:
   reslver-kit [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   init     initialize reslver toolkit
   apply    generate diagrams from terraform states
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

### reslver-kit init --help

```
NAME:
   reslver-kit init - initialize reslver toolkit

USAGE:
   reslver-kit init [command options] [arguments...]

OPTIONS:
   --config DIR, -c DIR  Load configuration from DIR (default: "/**/Project/reslver/reslver-kit/examples/.reslver/configs/") [$RESLVER_PATH]
   --force, -f           Force initialize all reslver configurations (default: false)
   --help, -h            show help (default: false)
```

### reslver-kit apply --help

```
NAME:
   reslver-kit apply - generate diagrams from terraform states

USAGE:
   reslver-kit apply [command options] [arguments...]

OPTIONS:
   --debug, -d                  Enable debug mode (default: false)
   --yaml-config FILE, -y FILE  Load graph YAML configuration from FILE (default: "/**/Project/reslver/reslver-kit/examples/graph.yaml")
   --config DIR, -c DIR         Load configurations from DIR (default: "/**/reslver/reslver-kit/examples/.reslver/configs/") [$RESLVER_PATH]
   --input DIR, -i DIR          Load terraform states from DIR (default: "/**/Project/reslver/reslver-kit/examples/")
   --output DIR, -o DIR         Output results to DIR (default: "/**/Project/reslver/reslver-kit/examples/")
   --help, -h                   show help (default: false)
```
