package flag_test

import (
	"flag"
	"testing"
	"time"
)

// Arg 返回第 i 个命令行参数
func TestFlagArg(t *testing.T) {
	flag.Parse()
	t.Logf("The first arg: %v", flag.Arg(0))
}

func TestFlagArgs(t *testing.T) {
	flag.Parse()
	t.Logf("The args: %v", flag.Args())
}

func TestMain(m *testing.M) {
	// 带默认值和提示语句的bool类型flag， 例如：-arg1=true
	boolArgPtr := flag.Bool("arg1", false, "This is a bool argument")

	// 将指定flag值绑定到一个bool变量上，返回该变量的指针
	var boolVarArg bool
	flag.BoolVar(&boolVarArg, "arg2", false, "This is a bool argument")

	// 带默认值和提示语句的duration类型flag， 例如：-arg3=5m
	durationArgPtr := flag.Duration("arg3", 0, "This is a duration argument")

	// 将一个time.Duration类型的flag值绑定到一个time.Duration的变量
	var durationVarArg time.Duration
	flag.DurationVar(&durationVarArg, "arg4", 0, "This is a duration argument")

	flag.Parse()

	println("The value of 'arg1': ", *boolArgPtr)
	println("The value of 'arg2': ", boolVarArg)
	println("The value of 'arg3': ", *durationArgPtr)
	println("The value of 'arg4': ", durationVarArg)
}
