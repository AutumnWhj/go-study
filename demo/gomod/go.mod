// 模块名
module gomod
// go sdk版本
go 1.21.4

// 依赖管理
require (
  github.com/gin-gonic/gin v1.6.3
  github.com/go-sql-driver/mysql v1.5.0
  github.com/jinzhu/gorm v1.9.16
  github.com/jinzhu/inflection v1.0.0
  github.com/jinzhu/now v1.1.1
)
// 排除第三方包
exclude (
  github.com/jinzhu/gorm v1.9.16
  github.com/jinzhu/inflection v1.0.0
  github.com/jinzhu/now v1.1.1
)
// 替换版本
// 修复依赖包的路径和版本
// 依赖包发送迁移，原始包无法访问，使用本地包替换原始包
replace (
  source latest => target latest
)

// 撤回
// d当前项目作为其他项目的依赖，如果某个版本出现问题则撤回该版本
retract (
  v1.0.0
)