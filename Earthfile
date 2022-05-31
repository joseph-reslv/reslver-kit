VERSION 0.6
FROM golang:1.17
WORKDIR /reslver-kit

# build go tool

clone-reslver:
  GIT CLONE git@git.k8s.app:joseph/reslver.git reslver
  WORKDIR reslver
  SAVE ARTIFACT . /main
  SAVE ARTIFACT ./sources /sources

clone-reslver-tf-loader:
  GIT CLONE git@git.k8s.app:joseph/reslver-tf-loader.git reslver-tf-loader
  WORKDIR reslver-tf-loader
  SAVE ARTIFACT . /main
  SAVE ARTIFACT ./sources /sources

clone-reslver-graph-exporter:
  GIT CLONE git@git.k8s.app:joseph/reslver-graph-exporter.git reslver-graph-exporter
  WORKDIR reslver-graph-exporter
  SAVE ARTIFACT . /main
  SAVE ARTIFACT ./sources /sources

clone-reslver-excel-exporter:
  GIT CLONE git@git.k8s.app:joseph/reslver-excel-exporter.git reslver-excel-exporter
  WORKDIR reslver-excel-exporter
  SAVE ARTIFACT . /main

clone-reslver-static-graph-exporter:
  GIT CLONE git@git.k8s.app:joseph/reslver-static-graph-exporter.git reslver-static-graph-exporter
  WORKDIR reslver-static-graph-exporter
  RUN tar cvzf reslver.tar.gz ./reslver-graph
  SAVE ARTIFACT ./reslver.tar.gz /sources/reslver.tar.gz

clone-reslver-configs:
  GIT CLONE git@git.k8s.app:joseph/reslver-configs reslver-configs
  WORKDIR reslver-configs
  SAVE ARTIFACT . /sources

clone:
  COPY --dir +clone-reslver/ reslver-repo
  COPY --dir +clone-reslver-tf-loader/ reslver-tf-loader-repo
  COPY --dir +clone-reslver-graph-exporter/ reslver-graph-exporter-repo
  COPY --dir +clone-reslver-excel-exporter/ reslver-excel-exporter-repo
  COPY --dir +clone-reslver-configs/ reslver-configs-repo
  COPY --dir +clone-reslver-static-graph-exporter/ reslver-static-graph-exporter-repo

  SAVE ARTIFACT reslver-repo/main reslver
  SAVE ARTIFACT reslver-tf-loader-repo/main reslver-tf-loader
  SAVE ARTIFACT reslver-graph-exporter-repo/main reslver-graph-exporter
  SAVE ARTIFACT reslver-excel-exporter-repo/main reslver-excel-exporter

  SAVE ARTIFACT reslver-repo/sources sources/reslver
  SAVE ARTIFACT reslver-tf-loader-repo/sources sources/reslver-tf-loader
  SAVE ARTIFACT reslver-graph-exporter-repo/sources sources/reslver-graph-exporter
  SAVE ARTIFACT reslver-configs-repo sources/reslver-configs
  SAVE ARTIFACT reslver-static-graph-exporter-repo sources/reslver-static-graph-exporter

deps:
  COPY --dir +clone/ ./
  COPY go.mod go.sum ./
  RUN go mod download

build:
  FROM +deps
  COPY --dir cmd kit logger templates types ./
  COPY main.go ./
  RUN go mod tidy

  RUN go build -o build/reslver-kit
  SAVE ARTIFACT build/reslver-kit /reslver-kit AS LOCAL build/reslver-kit

