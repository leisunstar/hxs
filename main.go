package main

import (
	"flag"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/yzw/conf"
	"github.com/yzw/logs"
	"hxs/cache"
	"hxs/config"
	"hxs/controller"
	"hxs/routers"
	"net/http"
	"time"
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
	if confFile == "" {
		confFile = "conf/app_my.conf"
	}
	// 初始化配置
	if err = conf.LoadConfig(confFile, config.App); err != nil {
		panic(err)
	}

	// 初始化日志
	if err := logs.Init(config.App.LogsConf); err != nil {
		panic(err)
	}
	defer logs.Close()

	//初始化controller
	if err = controller.Init(); err != nil {
		panic(err)
	}

	cache.Init()

	//运行马提尼
	m := &martini.ClassicMartini{}
	m.Martini = martini.New()
	m.Router = martini.NewRouter()
	m.Use(martini.Recovery())
	m.Use(martini.Static("public"))
	m.MapTo(m.Router, (*martini.Routes)(nil))
	m.Action(m.Router.Handle)
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		Layout:     "layout",
	}))
	m.Use(func(c martini.Context, req *http.Request) {
		begin := time.Now().UnixNano()
		now := time.Now().Format("2006-01-02 15:04:05")
		c.Next()
		end := time.Now().UnixNano()
		logs.Info("%s| %s| %s|cost %d", now, req.Method, req.RequestURI, end-begin)
	})
	routers.Router(m)
	logs.Debug("listen on :%s", config.App.ListenHost)
	m.RunOnAddr(":" + config.App.ListenHost)
	m.Run()

}
