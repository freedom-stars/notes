# git

## 本地新建git推送到已存在分支

```shell
git init
git add .
gitcommit -m 'init commit'
git remote add origin git@github.com:freedom-stars/notes.git
```

因为远端有本地没有的提交 所以推送会失败

```shell
To github.com:freedom-stars/notes.git
 ! [rejected]        master -> master (fetch first)
error: 推送一些引用到 'git@github.com:freedom-stars/notes.git' 失败
提示：更新被拒绝，因为远程仓库包含您本地尚不存在的提交。这通常是因为另外
提示：一个仓库已向该引用进行了推送。再次推送前，您可能需要先整合远程变更
提示：（如 'git pull ...'）。
提示：详见 'git push --help' 中的 'Note about fast-forwards' 小节。
```

直接git pull也会失败， 需要添加上 `–allow-unrelated-histories` 参数

```shell
 * branch            master     -> FETCH_HEAD
 fatal: 拒绝合并无关的历史
```

```shell
git pull origin master --allow-unrelated-histories
```

然后在本地合并，提交 `git push -u origin master`

```shell
# 若push失败 可能需要关联远程分支
git branch –set-upstream-to=origin/master master
```

最终提交到远程
