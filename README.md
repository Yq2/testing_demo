# testing_demo
golang测试语法提点
t *testing.T
t.SkipNow() 会跳过当前test case并做PASS处理，并继续执行下一个test case
t.Run()来执行subtest可以做到控制test输出以及test执行顺序，因为默认情况下每个test case会被go test并发执行
func TestMain(m *testing.M) {
	fmt.Println("test main first")
	m.Run()
}
TestMain在一个test case中只能有一个

使用TestMain作为初始化test，并使用m.Run()来调用其他test可以完成一些需要初始化的testing，
比如数据库连接，文件打开，REST服务登录
如果没有在TestMain中调用m.Run()则除了TestMain以外的其他tests都不会被执行
如果没有TestMain函数的话，每个合法test case都会正常执行

Benchmark测试case 同样受到TestMain规则限制
Benchmark测试 应该保证测试函数能达到稳态，否则会一直测试到卡死
