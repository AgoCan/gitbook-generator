package main

import (
	_ "embed"
	"path"

	"github.com/agocan/gitbook-generator/config"
	"github.com/agocan/gitbook-generator/generator"
)

var Option generator.Option

var (
	//go:embed template/book.json.tmpl
	apiContent string
)

var files = map[string]string{
	"book.json": apiContent,
}

func main() {
	config.DefaultConfig()
	generator.Init()

	Option = generator.Option{
		AbsProjectPath: config.AbsProjectPath,
		Plugins:        config.Plugins,
		Title:          *config.Title,
		SidebarLink:    *config.SidebarLink,
		SidebarTitle:   *config.SidebarTitle,
		Author:         *config.Author,
		Description:    *config.Description,
		Keywords:       *config.Keywords,
		GaToken:        *config.GaToken,
		TreeBaToken:    *config.TreeBaToken,
	}
	// Option.AbsProjectPath = path.Join(*config.ProjectPath, *config.Title)
	Option.AbsProjectPath = path.Join(*config.ProjectPath, "")
	var fileGen generator.FileGenerator
	fileGen.Files = files
	// 注册
	generator.Register("files", &fileGen)
	generator.RunGenerator(&Option)
}
