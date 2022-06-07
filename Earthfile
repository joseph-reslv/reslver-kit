VERSION 0.6
FROM golang:1.18
WORKDIR /reslver-kit

### Release Flow ###
# 1. clone submodules git repository
# 2. save them as source code (main) & template (sources) 
# 3. download go dependencies (go mod download)
# 4. copy reslver-kit source codes (folder & main.go)
# 5. build reslver-kit
# 6. install goreleaser to release go tool (it can build cross platforms binary)
# 7. copy reslver-kit .git for goreleaser (goreleaser requires git repo)
# 8. run goreleaser to release (build) cross platforms binary LOCALLY
# 9. release all binary to git repo

clone-reslver:
  GIT CLONE --branch v0.1.0 git@git.k8s.app:resolve/reslver.git reslver
  WORKDIR reslver
  SAVE ARTIFACT ./sources /sources

clone-reslver-tf-loader:
  GIT CLONE --branch  v0.1.0 git@git.k8s.app:resolve/reslver-tf-loader.git reslver-tf-loader
  WORKDIR reslver-tf-loader
  SAVE ARTIFACT ./sources /sources

clone-reslver-graph-exporter:
  GIT CLONE --branch  v0.1.0 git@git.k8s.app:resolve/reslver-graph-exporter.git reslver-graph-exporter
  WORKDIR reslver-graph-exporter
  SAVE ARTIFACT ./sources /sources

clone-reslver-static-graph-generator:
  GIT CLONE git@git.k8s.app:resolve/reslver-static-graph-generator.git reslver-static-graph-generator
  WORKDIR reslver-static-graph-generator
  RUN tar cvzf reslver.tar.gz ./reslver-graph
  SAVE ARTIFACT ./reslver.tar.gz /sources/reslver.tar.gz

clone-reslver-configs:
  GIT CLONE git@git.k8s.app:resolve/reslver-configs.git reslver-configs
  WORKDIR reslver-configs
  SAVE ARTIFACT . /sources

clone:
  COPY --dir +clone-reslver/ reslver-repo
  COPY --dir +clone-reslver-tf-loader/ reslver-tf-loader-repo
  COPY --dir +clone-reslver-graph-exporter/ reslver-graph-exporter-repo
  COPY --dir +clone-reslver-configs/ reslver-configs-repo
  COPY --dir +clone-reslver-static-graph-generator/ reslver-static-graph-generator-repo

  SAVE ARTIFACT reslver-repo/sources sources/reslver
  SAVE ARTIFACT reslver-tf-loader-repo/sources sources/reslver-tf-loader
  SAVE ARTIFACT reslver-graph-exporter-repo/sources sources/reslver-graph-exporter
  SAVE ARTIFACT reslver-configs-repo sources/reslver-configs
  SAVE ARTIFACT reslver-static-graph-generator-repo sources/reslver-static-graph-generator

deps:
  ### workaround to get private repository package for Golang ###
  # change git to use SSH instead of https
  # attach ssh (use ssh-agent from env) 
  # bypass golang security setting (must set GOINSECURE and GOPRIVATE)
  RUN git config --global url."ssh://git@git.k8s.app/".insteadOf "https://git.k8s.app/"
  RUN mkdir -p -m 0600 ~/.ssh && ssh-keyscan git.k8s.app >> ~/.ssh/known_hosts
  RUN go env -w GOINSECURE=git.k8s.app
  RUN go env -w GOPRIVATE=git.k8s.app
  RUN git config --global http.sslVerify false
  ###
  COPY --dir +clone/ ./
  COPY go.mod go.sum ./
  RUN --ssh go mod download
  SAVE ARTIFACT go.mod
  SAVE ARTIFACT go.sum

use-go-releaser:
  RUN go install github.com/goreleaser/goreleaser@latest
  SAVE ARTIFACT $GOPATH/bin

build:
  FROM +deps
  COPY . .
  RUN go build -o reslver-kit
  SAVE ARTIFACT reslver-kit AS LOCAL dist/reslver-kit

release-local:
  FROM +build
  COPY --dir +use-go-releaser/bin $GOPATH/
  RUN goreleaser release --snapshot --rm-dist
  SAVE ARTIFACT dist AS LOCAL dist

release:
  ARG GITHUB_TOKEN
  FROM +build
  COPY --dir +use-go-releaser/bin $GOPATH/
  RUN goreleaser release

test:
  # nothing to do
