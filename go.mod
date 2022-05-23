module git.k8s.app/joseph/reslver-kit

go 1.17

require (
	github.com/go-git/go-git/v5 v5.4.2
	github.com/urfave/cli/v2 v2.7.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/Microsoft/go-winio v0.4.16 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20210428141323-04723f9f07d7 // indirect
	github.com/acomagu/bufpipe v1.0.3 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/go-git/gcfg v1.5.0 // indirect
	github.com/go-git/go-billy/v5 v5.3.1 // indirect
	github.com/google/go-jsonnet v0.18.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	golang.org/x/net v0.0.0-20210326060303-6b1517762897 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20210502180810-71e4cd670f79 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	sigs.k8s.io/yaml v1.1.0 // indirect
)

require (
	git.k8s.app/joseph/reslver v0.0.0
	git.k8s.app/joseph/reslver-graph-exporter v0.0.0
	git.k8s.app/joseph/reslver-tf-loader v0.0.0
	github.com/antzucaro/matchr v0.0.0-20210222213004-b04723ef80f0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
)

replace (
	git.k8s.app/joseph/reslver v0.0.0 => ./reslver
	git.k8s.app/joseph/reslver-graph-exporter v0.0.0 => ./reslver-graph-exporter
	git.k8s.app/joseph/reslver-tf-loader v0.0.0 => ./reslver-tf-loader
)
