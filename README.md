# Jsonix

[English Document](https://github.com/JsonLee12138/jsonix/blob/main/README.en.md)

Jsonix 是基于 Go + Fiber 开发的一套全新的框架，旨在提高项目开发效率。

该框架提供了一系列命令行工具（CLI），用于服务管理、项目初始化、代码生成以及自动化数据库迁移等任务。采用模块化设计，支持 controller、service、repository 等代码结构快速生成，并集成了丰富的功能模块。

## 安装

```bash
go install github.com/JsonLee12138/jsonix@latest
```

## 概览

- CLI 工具支持。
- 基础 `web` 功能支持。
- `fiber` 框架支持。
- `apifox` 自动上传支持。
- `swagger` 生成支持。
- `gorm` 数据库 ORM 支持。
- `redis` 缓存支持。
- `i18n` 国际化支持。
- `logger` 日志支持。
- `cors` 跨域支持。
- `jsonix` 热重载支持。
- `dig` 依赖注入支持。
- 配置文件读取支持。
- 提供多个基础配置:
  - [ApifoxConfig](https://github.com/JsonLee12138/jsonix/blob/main/pkg/configs/apifox.go)
  - [CorsConfig](https://github.com/JsonLee12138/jsonix/blob/main/pkg/configs/cors.go)
  - [I18nConfig](https://github.com/JsonLee12138/jsonix/blob/main/pkg/configs/i18n.go)
  - [LogConfig](https://github.com/JsonLee12138/jsonix/blob/main/pkg/configs/log.go)
  - [MysqlConfig](https://github.com/JsonLee12138/jsonix/blob/main/pkg/configs/mysql.go)
  - [RedisConfig](https://github.com/JsonLee12138/jsonix/blob/main/pkg/configs/redis.go)
  - [SwaggerConfig](https://github.com/JsonLee12138/jsonix/blob/main/pkg/configs/swagger.go)
  - [SystemConfig](https://github.com/JsonLee12138/jsonix/blob/main/pkg/configs/system.go)

## CLI

```bash
$ jsonix --help

  Usage: jsonix server

  Options:
    -e, --env <string>                  设置当前服务运行环境 (dev, prod, test)
    -w, --watch                         启用配置文件热重载
    -s, --show <string>                 查看指定端口是否被占用
    -k, --kill <string>                 杀死指定端口的进程
    -h, --help                          查询文档

  Usage: jsonix gen

  Options:
    -m, --module <string>               指定模块名称
    -s, --service                       生成服务代码
    -c, --controller                    生成控制器代码
    -r, --repository                    生成仓库代码
    -e, --entity                        生成实体代码
    -o, --override                      覆盖已存在的文件
    -h, --help                          查询文档

  Usage: jsonix migrate                 生成数据库自动迁移文件

  Options:
    -r, --root <string>                 指定项目根目录
    -d, --dest <string>                 指定自动迁移文件输出目录
    -h, --help                          查询文档

  Usage: jsonix init                    初始化项目

  Options:
    -h, --help                          查询文档
```

## 项目文件结构

```
├─ .vscode                     # VSCode 推荐配置
├─ .run                        # GoLand 运行命令
├─ apps                        # 应用模块
│  ├─ example                  # 示例模块
│  │  ├─ controller           # 控制器
│  │  ├─ repository          # 仓库
│  │  ├─ service             # 服务
│  │  └─ entry               # 入口文件
│  └─ ...                     # 其他模块
├─ auto_migrate               # 自动迁移目录(可通过 jsonix migrate 命令生成, 无需改动)
│  └─ ...                     # 自动迁移文件
├─ config                     # 配置文件目录
│  ├─ config.yaml            # 配置文件
│  ├─ config.dev.yaml        # 开发环境配置
│  ├─ config.test.yaml       # 测试环境配置
│  ├─ config.prod.yaml       # 生产环境配置
│  ├─ config.local.yaml      # 本地环境配置
│  ├─ config.dev.local.yaml  # 开发本地环境配置
│  ├─ config.test.local.yaml # 测试本地环境配置
│  ├─ config.prod.local.yaml # 生产本地环境配置
│  ├─ regexes.yaml           # uaparser 配置文件
│  └─ ...                     # 其他配置文件
├─ configs                    # 配置文件实例目录
│  ├─ configs.yaml           # 全部配置实例
│  └─ ...                     # 其他配置实例文件
├─ core                       # 核心模块目录
├─ docs                       # swagger 文档目录
├─ locales                    # 语言包目录
├─ logs                       # 日志目录
├─ middleware                 # 中间件目录
├─ tmp                        # air 启动临时文件(不要改动, 不要提交)
├─ utils                      # 工具函数目录
├─ .air.toml                  # air 配置文件
├─ main.go                    # 主函数
├─ go.mod                     # Go 模块文件
└─ go.sum                     # Go 模块文件
```

## 反馈建议

如果你在使用过程中遇到问题或有新功能建议,欢迎通过以下方式反馈:

- 在 GitHub 提交 Issue
- 发送邮件至: lijunsong2@gmail.com
- 个人微信: Json_Lee12138
- 个人 QQ: 2622336659
