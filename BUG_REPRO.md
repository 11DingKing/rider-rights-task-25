# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

女骑手事项详情页只显示最近 100 条审计记录，早期的权益帮扶和转派操作在详情里消失了，后台总数却更大。请修复详情审计的分页读取，把所有页面汇总起来并让空页安全结束；本次只动生产代码，不要修改测试文件。

## 含 Bug 版本

- 仓库：11DingKing/rider-rights-task-25
- 仓库地址：https://github.com/11DingKing/rider-rights-task-25.git
- parent SHA：359e8aa563c92cd6fd23937760826832c54085aa

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/rider-rights-task-25.git bug-repro
cd bug-repro
git checkout --detach 359e8aa563c92cd6fd23937760826832c54085aa
go test ./internal/domain -run "^TestDetailAuditContinuesAfterFirstPage$" -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestDetailAuditContinuesAfterFirstPage$" -count=1
--- FAIL: TestDetailAuditContinuesAfterFirstPage (0.00s)
    task25_test.go:8: detail audit stopped after first page
FAIL
FAIL	riderguard/internal/domain	0.045s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestDetailAuditContinuesAfterFirstPage$" -count=1
--- FAIL: TestDetailAuditContinuesAfterFirstPage (0.00s)
    task25_test.go:8: detail audit stopped after first page
FAIL
FAIL	riderguard/internal/domain	0.002s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

修复后，完整页之后必须继续读取下一页，最后一页或空页才结束，详情返回的审计顺序和总量正确；上下文取消和不前进游标应安全失败。定向测试、相关包测试及全量回归必须通过，不得删除、跳过或削弱测试。
