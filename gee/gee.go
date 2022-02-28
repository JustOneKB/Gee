package gee

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

/*
	golang 中允许定义函数类变量
	这地方是用 HandlerFunc 这个别名
	取代 框架自定义的 Context
	Context 中封装了 (http.ResponseWriter, *http.Request)
	方便变量定义时候的书写

	同时，被这么定义的函数，都可以实现 ServeHTTP
	也就是说，可以用来实现 http 的相关服务
*/
type HandlerFunc func(*Context)

type Engine struct {
	*RouterGroup                //嵌套，表示Engine 继承 RouterGroup，此处是*号声明，还有不带*号的，区别可能是能否改动父的内容
	router       *router        //路由解析
	groups       []*RouterGroup //记录所有分组

	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

type RouterGroup struct {
	prefix      string        //这个分组的前缀标志
	middlewares []HandlerFunc //分组中使用的中间件
	engine      *Engine       //最终http的服务还是由engine来实现
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

/*
	创建一个分组
	并且使用prefix表示分组的识别前缀
*/
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine //浅拷贝
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix, //新分组的前缀 = 当前所在组前缀 + 自身特有的前缀
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

//将分组调用的中间件（函数）添加进middlewares中，后面按照顺序调用
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

//用 engine 这个对象承接请求并完成回应
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

/*
	解析请求路径，查找路由映射表
	如果查到，就执行注册的处理方法
	如果查不到，就返回 404 NOT FOUND
*/
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	c.engine = engine
	engine.router.handle(c)
}

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	group.GET(urlPattern, handler)
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}
