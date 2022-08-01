package config

import (
	"flag"
)

// 直接写配置成配置文件，不多使用额外的配置文件，例如不使用config.yaml

var (
	Title          = flag.String("n", "demo", "Ansible scripts dir name.")
	ProjectPath    = flag.String("p", ".", "project path")
	Item           = flag.String("i", "ansible", "project")
	AbsProjectPath string
)

func DefaultConfig() {
	flag.Parse()
}
