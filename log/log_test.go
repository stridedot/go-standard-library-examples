package log_test

import (
	"log"
	"os"
	"testing"
)

func TestLogFatal(t *testing.T) {
	log.Fatal("fatal")
}

func TestLogFatalf(t *testing.T) {
	log.Fatalf("fatal %s", "error")
}

func TestLogFatalln(t *testing.T) {
	log.Fatalln("fatalln", "error")
}

func TestLogFlags(t *testing.T) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	t.Logf("log.Flags: %v", log.Flags())
	t.Logf("log.Ldate: %v", log.Ldate)
	t.Logf("log.Ltime: %v", log.Ltime)
	t.Logf("log.Lmicroseconds: %v", log.Lmicroseconds)
	t.Logf("log.Llongfile: %v", log.Llongfile)
	t.Logf("log.Lshortfile: %v", log.Lshortfile)
	t.Logf("log.LstdFlags: %v", log.LstdFlags) //LstdFlags     = Ldate | Ltime
}

// 自定义log输出
func TestLogOutput(t *testing.T) {
	l := log.New(os.Stdout, "log->", log.Ldate)
	l.Output(2, "log output1")

	log.SetOutput(os.Stdout)
	log.Output(2, "log output2")

	w := log.Writer()
	log.SetOutput(w)
	log.Output(2, "log output3")
}

func TestLogPanic(t *testing.T) {
	log.Panic("panic")
}

func TestLogPanicf(t *testing.T) {
	log.Panicf("panic %s", "error")
}

func TestLogPanicln(t *testing.T) {
	log.Panicln("panicln", "error")
}

func TestLogPrefix(t *testing.T) {
	log.SetPrefix("log->")
	t.Logf("log.Prefix: %v", log.Prefix())
}

func TestLoggerDefault(t *testing.T) {
	logger := log.Default()
	logger.Println("log default")
}

// 测试log.New
func TestLoggerNew(t *testing.T) {
	file, err := os.Create("test.log")
	if err != nil {
		t.Fatalf("create file error: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "log->", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("log new2")
}

func TestLoggerFatal(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "logger fatal->", log.Ldate|log.Ltime|log.Llongfile)
	logger.Fatal("logger fatal")
}

func TestLoggerFatalf(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "logger fatalf->", log.Ldate|log.Ltime|log.Llongfile)
	logger.Fatalf("logger fatalf %s", "error")
}

func TestLoggerFatalln(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "logger fataln->", log.Ldate|log.Ltime|log.Llongfile)
	logger.Fatalln("logger fataln", "error")
}

func TestLoggerFlags(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "logger flags->", log.Ldate|log.Ltime|log.Llongfile)
	t.Logf("logger.Flags: %v", logger.Flags())
}

func TestLoggerOutput(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "logger output->", log.Ldate|log.Ltime|log.Llongfile)
	logger.Output(2, "logger output test")
}

func TestLoggerPanic(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "logger panic->", log.Ldate|log.Ltime|log.Llongfile)
	logger.Panic("logger panic")
}

func TestLoggerPanicf(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "logger panicf->", log.Ldate|log.Ltime|log.Llongfile)
	logger.Panicf("logger panicf %s", "error")
}

func TestLoggerPanicln(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "logger panicln->", log.Ldate|log.Ltime|log.Llongfile)
	logger.Panicln("logger panicln", "error")
}

func TestLoggerPrefix(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "", log.Ldate|log.Ltime|log.Llongfile)
	logger.SetPrefix("[logger prefix->]")
	logger.Println("logger prefix")
}

func TestLoggerPrint(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "logger print->", log.Ldate|log.Ltime|log.Llongfile)
	logger.Print("logger print")
}

func TestLoggerPrintf(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "logger printf->", log.Ldate|log.Ltime|log.Llongfile)
	logger.Printf("logger printf %s", "error")
}

func TestLoggerPrintln(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("logger println")
}

func TestLoggerWriter(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	logger := log.New(file, "logger writer->", log.Ldate|log.Ltime|log.Llongfile)
	t.Logf("logger.Writer: %v", logger.Writer())
}
