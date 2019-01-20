###安装bee工具
    go get github.com/beego/bee

###安装依赖工具
    go get github.com/kardianos/govendor

###初始化依赖配置(可跳过)
    govendor init
 
###添加依赖包  
    govendor add +external
   
###直接获取依赖包
    govendor fetch golang.org/x/net/context@v1 

###依赖类型
 | 状态	|  缩写状态	|    含义 |
 | ----	| :---:|   :--- |
 |+local|	    l	|    本地包，即项目自身的包组织|
|+external	|e	|    外部包，即被 $GOPATH 管理，但不在 vendor 目录下|
|+vendor|	    v	|    已被 govendor 管理，即在 vendor 目录下|
|+std	 |   s	|    标准库中的包|
|+unused	|    u	|    未使用的包，即包在 vendor 目录下，但项目并没有用到|
|+missing	|m	|    代码引用了依赖包，但该包并没有找到|
|+program	|p	|    主程序包，意味着可以编译为执行文件|
|+outside	|	|    外部包和缺失的包|
|+all		|    |    所有的包|

###常用命令
 |命令| 说明 |
 | ----:	| :---|
|init   |  创建 vendor 文件夹和 vendor.json 文件|
|list   |  列出已经存在的依赖包|
|add    |  从 $GOPATH 中添加依赖包，会加到 vendor.json|
|update  | 从 $GOPATH 升级依赖包|
|remove |  从 vendor 文件夹删除依赖|
|status |  列出本地丢失的、过期的和修改的package|
|fetch |  从远端库增加新的，或者更新 vendor 文件中的依赖包|
|sync  |   Pull packages into vendor folder from remote repository with revisions|
|migrate | Move packages from a legacy tool to the vendor folder with metadata.|
|get |    类似 go get，但是会把依赖包拷贝到 vendor 目录|
|license | List discovered licenses for the given status or import paths.|
|shell   | Run a "shell" to make multiple sub-commands more efficient for large projects.|

go tool commands that are wrapped:
      `+<status>` package selection may be used with them
    fmt, build, install, clean, test, vet, generate, tool

