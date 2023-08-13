package flag_test

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
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

	// 带默认值和提示语句的float64类型flag，例如 -arg5=5.5
	float64ArgPtr := flag.Float64("arg5", 4.4, "This is a float64 argument")

	var float64VarArg float64
	flag.Float64Var(&float64VarArg, "arg6", 4.4, "This is a float64 argument")

	// 带默认值和提示语句的int类型flag，例如 -arg7=5
	intArgPtr := flag.Int("arg7", 0, "This is a int argument")

	var intVarArg int
	flag.IntVar(&intVarArg, "arg8", 0, "This is a int argument")

	// 带默认值和提示语句的int64类型flag，例如 -arg9=5
	int64ArgPtr := flag.Int64("arg9", 0, "This is a int64 argument")

	var int64VarArg int64
	flag.Int64Var(&int64VarArg, "arg10", 0, "This is a int64 argument")

	flag.Parse()

	println("The value of 'arg1': ", *boolArgPtr)
	println("The value of 'arg2': ", boolVarArg)
	println("The value of 'arg3': ", *durationArgPtr)
	println("The value of 'arg4': ", durationVarArg)
	fmt.Printf("The value of 'arg5':  %.2f\n", *float64ArgPtr)
	println("The value of 'arg5': ", strconv.FormatFloat(*float64ArgPtr, 'f', 2, 64))
	println("The value of 'arg6': ", float64VarArg)
	println("The value of 'arg7': ", *intArgPtr)
	println("The value of 'arg8': ", intVarArg)
	println("The value of 'arg9': ", *int64ArgPtr)
	println("The value of 'arg10': ", int64VarArg)
}

// 测试使用自定义命令
// 注意：flag.Func 只能用在非测试函数中
func TestFlagFunc(t *testing.T) {
	// First we need to declare variables to hold the values from the
	// command-line flags. Notice that we also need to set any defaults,
	// which will be used if the relevant flag is not provided at runtime.
	var (
		urls  []string                    // Default of the empty slice
		pause time.Duration = time.Second // Default of one second
	)

	// The flag.Func() function takes three parameters: the flag name,
	// descriptive help text, and a function with the signature
	// `func(string) error` which is called to process the string value
	// from the command-line flag at runtime and assign it to the necessary
	// variable. In this case, we use strings.Fields() to split the string
	// based on whitespace and store the resulting slice in the urls
	// variable that we declared above. We then return nil from the
	// function to indicate that the flag was parsed without any errors.
	flag.Func("urls", "List of URLs to print", func(flagValue string) error {
		urls = strings.Fields(flagValue)
		return nil
	})

	// Likewise we can do the same thing to parse the pause duration. The
	// time.ParseDuration() function may throw an error here, so we make
	// sure to return that from our function.
	flag.Func("pause", "Duration to pause between printing URLs", func(flagValue string) error {
		var err error
		pause, err = time.ParseDuration(flagValue)
		return err
	})

	// Importantly, call flag.Parse() to trigger actual parsing of the
	// flags.
	flag.Parse()

	// Print out the URLs, pausing between each iteration.
	for _, u := range urls {
		t.Log(u)
		time.Sleep(pause)
	}
}
