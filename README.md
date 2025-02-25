# gefs

gin embed file system

# go:embed static

gin框架搭配embed文件系统生成静态文件服务器

静态目录`./static`（编译时）映射到url为`/fs/static`

如文件`./static/hello.html`映射到服务url为`/fs/static/hello.html`

# static参数

执行命令通过参数`-static=dir`，可以加载用户自定义的静态目录，映射到url为`/static`

如文件`dir/x.png`映射到url为`/static/x.png`

# index参数

执行命令通过参数`-index=hello.html`，可以使url`/`直接访问到`./static/hello.html`文件，响应类型为文件

但如果有设置static参数`-static=dir`，则会使`/`直接访问到`dir/hello.html`，响应类型为文件

# chart参数

`chart.html`实现了一个图表显示页面，读取`/static/public/data.js`获取数据显示在页面上

使用命令行参数`-index=chart.html`可以直接在服务url`/`访问图表

# port参数

使用参数`-port=:8888`可以修改服务启动监听的端口
