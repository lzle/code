# Git

## 目录
* [基础](#基础)
    * [配置](#配置)
    * [commit 描述](#commit-描述)
* [命令](#命令)
    * [重命名](#重命名)
    * [历史](#历史)
    * [删除分支](#删除分支)
    * [修改 commit](#修改-commit)
    * [合并 commit](#合并-commit)
    * [查看差异](#查看差异)
    * [stash](#stash)


## 基础

### 配置

以用以下命令查看 Git 的配置。

1. 查看所有配置（包含系统级、全局和仓库级）

```bash
git config --list --show-origin
```

2. 查看某一级别的配置

```bash
# 系统级（对所有用户生效，一般在 /etc/gitconfig）：
git config --system --list

# 全局级（当前用户生效，一般在 ~/.gitconfig）：
git config --global --list

# 仓库级（当前仓库生效，一般在 .git/config）：
git config --local --list
```

3. 查看用户配置

```bash
git config user.name
git config user.email
```

4. 设置用户配置

```bash
# 设置全局用户名和邮箱
git config --global user.name "Your Name"
git config --global user.email "you@example.com"

# 设置仓库级用户名和邮箱
git config --local user.name "Your Name"
git config --local user.email "you@example.com"
```

### commit 描述

`commit` 基本格式。

```
<type>(<scope>): <subject>

<body>

<footer>
```

常见 `type` 类型：

- feat: 新功能（feature）
- fix: 修复 bug
- docs: 文档（documentation）
- style: 代码格式（不影响代码运行的变动）
- refactor: 重构（即不是新增功能，也不是修改 bug 的代码变动）
- perf: 性能优化
- test: 添加测试
- chore: 构建过程或辅助工具的变动
- revert: 回滚某个更早之前的提交
- build: 影响构建系统或外部依赖项的更改（例如：gulp、broccoli、npm）
- ci: 持续集成相关的更改（例如：Travis、Circle、BrowserStack、SauceLabs）
- wip: 开发中的功能
- release: 发布版本
- dep: 依赖更新
- hotfix: 紧急修复

## 命令

### 重命名

可以使用 `git mv` 命令来重命名文件或目录。这个命令会同时更新 `Git` 的索引和工作目录。

```bash
git mv old_filename new_filename
```

### 历史

`git log` 命令用于查看提交历史记录。它显示了项目的提交信息，包括提交的哈希值、作者、日期和提交消息。

```bash
git log -n8 --oneline --graph --all

* ec4f5c6c (origin/master_rocksdb, master_rocksdb) add 1 server to repair
* 76342e8e europe cluster expand 10 servers
* d9902424 completed the repair of 1 server
* cfe00f05 hkong cluster expand 7 servers
* 6c7acf9d hkong cluster replace zk/ectd/dashborad node
* b5435551 add 1 server to repair
* f106b2ec add 1 server to repair
* fcb53ea4 reduce group file consistent check sec
```

打开浏览器查看 `git log` 参数：

```bash
git config --global web.browser open

git help --web log
```

### 删除分支

删除本地分支：

```bash
git branch -d branch_name  # 删除已合并的分支
git branch -D branch_name  # 强制删除未合并的分支
```

删除远程分支：

```bash
git push origin --delete branch_name
# 或者
git push origin :branch_name
```


### 修改 commit

1. 修改最近一次提交的 commit 信息：

```bash
git commit --amend
```

2. 修改更早之前的 commit 信息：

```bash
git rebase -i HEAD~n  # n 是要修改的提交数量

# 在打开的编辑器中，将要修改的提交前的 "pick" 改为 "reword"
# 保存并关闭编辑器
# 按照提示修改提交信息
```

### 合并 commit

1. 使用交互式 rebase 合并多个 commit：

```bash
git rebase -i HEAD~n  # n 是要合并的提交数量

# 在打开的编辑器中，将要合并的提交前的 "pick" 改为 "squash" 或 "s"（表示将该提交合并到前一个提交中）。
# 保存并关闭编辑器。
# 按照提示修改合并后的提交信息，保存并关闭编辑器。
```
2. 合并不连续的多个 commit：

```bash
git rebase -i HEAD~n  # n 是要合并的提交数量

pick  aaaaa
pick  bbbb
pick  ccccc

修改为下面，合并 aaaaa 和 ccccc：

pick  aaaaa
squash  ccccc
pick  bbbb

# 保存并关闭编辑器。
# 按照提示修改合并后的提交信息，保存并关闭编辑器。
```

### 查看差异

比较暂存区和commit的差异：

```bash
git diff --cached
```

比较工作区和暂存区的差异：

```bash
git diff
```

查看不同 commit 之间的差异：

```bash
git diff commit1 commit2
```

查看某个文件在不同 commit 之间的差异：

```bash
git diff commit1 commit2 file_path
```

### stash

`git stash` 命令用于临时保存当前工作目录的修改，以便你可以切换分支或执行其他操作，而不需要提交这些修改。
1. 保存当前修改：

```bash
git stash
```

2. 查看保存的修改列表：

```bash
git stash list
```

3. 应用最近保存的修改：

```bash
git stash apply
```

4. 应用并删除最近保存的修改：
```bash
git stash pop
```

