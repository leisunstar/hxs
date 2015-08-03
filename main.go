package main

import (
	"flag"
	"github.com/go-martini/martini"
	"github.com/yzw/common/http"
	"github.com/yzw/conf"
	"github.com/yzw/logs"
	"hxs/config"
	"hxs/controller"
	"hxs/routers"
    "github.com/martini-contrib/render"
)

var (
	// 配置文件路径
	confFile string
)

// 初始化启动参数
func init() {
	flag.StringVar(&confFile, "c", "", "conf filename")
	flag.Parse()
}

// 主函数
func main() {
	var err error

	// 初始化配置
	if err = conf.LoadConfig(confFile, config.App); err != nil {
		panic(err)
	}

	// 初始化日志
	if err := logs.Init(config.App.LogsConf); err != nil {
		panic(err)
	}
	defer logs.Close()

	// 初始化pprof
	go http.RunPprof(config.App.HttpPprofAddr)

	//初始化controller
	if err = controller.Init(); err != nil {
		panic(err)
	}
	//运行马提尼
	m := martini.Classic()
    m.Use(render.Renderer(render.Options{
        Directory:  "templates",                // Specify what path to load the templates from.
        Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
        Layout:     "layout",
    }))
	routers.Router(m)
	logs.Debug("listen on :%s", config.App.ListenHost)
	m.RunOnAddr(":" + config.App.ListenHost)
	m.Run()

}
