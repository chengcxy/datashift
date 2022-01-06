
## 1. 基本介绍

### 1.1 项目介绍

> dadashift是一个中间件读写工具,日常开发中经常有类似


### 1.2 贡献指南
Hi! 首先感谢你使用 gin-vue-admin。

Gin-vue-admin 是一套为快速研发准备的一整套前后端分离架构式的开源框架，旨在快速搭建中小型项目。

Gin-vue-admin 的成长离不开大家的支持，如果你愿意为 gin-vue-admin 贡献代码或提供建议，请阅读以下内容。

#### 1.2.1 Issue 规范
- issue 仅用于提交 Bug 或 Feature 以及设计相关的内容，其它内容可能会被直接关闭。如果你在使用时产生了疑问，请到 Slack 或 [Gitter](https://gitter.im/ElemeFE/element) 里咨询。

- 在提交 issue 之前，请搜索相关内容是否已被提出。

#### 1.2.2 Pull Request 规范
- 请先 fork 一份到自己的项目下，不要直接在仓库下建分支。

- commit 信息要以`[文件名]: 描述信息` 的形式填写，例如 `README.md: fix xxx bug`。

- <font color=red>确保 PR 是提交到 `develop` 分支，而不是 `master` 分支。</font>

- 如果是修复 bug，请在 PR 中给出描述信息。

- 合并代码需要两名维护人员参与：一人进行 review 后 approve，另一人再次 review，通过后即可合并。

## 2. 使用说明

```
- node版本 > v12.18.3
- golang版本 >= v1.16
- IDE推荐：Goland
- 初始化项目： 不同版本数据库初始化不通 参见 https://www.gin-vue-admin.com/docs/first_master
- 替换掉项目中的七牛云公钥，私钥，仓名和默认url地址，以免发生测试文件数据错乱
```

### 2.1 server项目

```bash

# 克隆项目
git clone https://github.com/flipped-aurora/gin-vue-admin.git
# 进入bin文件夹
cd bin

# 编译 自动将datashift 移动到$GOBIN下
sh build.sh

# 运行任务
datashift 
```
# datashift
