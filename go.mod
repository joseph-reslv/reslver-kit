module git.k8s.app/joseph/reslver-kit

go 1.17

require github.com/urfave/cli/v2 v2.7.1

require (
	github.com/google/go-jsonnet v0.18.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	sigs.k8s.io/yaml v1.1.0 // indirect
)

require (
	git.k8s.app/joseph/reslver v0.0.0
	git.k8s.app/joseph/reslver-tf-loader v0.0.0
	git.k8s.app/joseph/reslver-graph-exporter v0.0.0
	github.com/antzucaro/matchr v0.0.0-20210222213004-b04723ef80f0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
)

replace (
	git.k8s.app/joseph/reslver v0.0.0 => ./reslver
	git.k8s.app/joseph/reslver-tf-loader v0.0.0 => ./reslver-tf-loader
	git.k8s.app/joseph/reslver-graph-exporter v0.0.0 => ./reslver-graph-exporter
)
