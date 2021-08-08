package config

import (
	"flag"
	"strings"
)

// 直接写配置成配置文件，不多使用额外的配置文件，例如不使用config.yaml

var (
	Plugins        []string
	TreeBaToken    = flag.String("t", "", "baidu tongji token")
	GaToken        = flag.String("g", "", "google tongji token")
	Title          = flag.String("n", "Title", "Book name")
	ExtraPlugins   = flag.String("e", "", `use "," split extra plugins.`)
	ProjectPath    = flag.String("p", ".", "project path")
	Keywords       = flag.String("k", "keywords", "keywords")
	Description    = flag.String("d", "description", "description")
	Author         = flag.String("a", "author", "author")
	SidebarTitle   = flag.String("st", "", "sidebarTitle")
	SidebarLink    = flag.String("sl", "", "sidebarLink")
	AbsProjectPath string
)

func DefaultConfig() {
	flag.Parse()
	// "ga" , "3-ba"
	Plugins = []string{"-highlight", "toggle-chapters", "codeblock-filename", "sectionx", "splitter", "-search",
		"-lunr", "search-pro", "theme-default", "-highlight", "prism", "prism-themes", "theme-comscore", "local-video",
		"include", "favicon", "anchors", "donate", "tbfed-pagefooter", "hide-element", "sitemap-general"}
	ExtraPluginsSlice := strings.Split(*ExtraPlugins, ",")
	if *ExtraPlugins != "" {
		Plugins = append(Plugins, ExtraPluginsSlice...)
	}
}
