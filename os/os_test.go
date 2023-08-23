package os_test

import (
	"os"
	"testing"
)

// 改变当前工作目录
func TestOSChdir(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal("Getwd failed")
	}
	t.Logf("Current dir: %s", dir)

	err = os.Chdir("D:\\projects\\go")
	if err != nil {
		t.Fatal("Chdir failed")
	}
	dir, err = os.Getwd()
	if err != nil {
		t.Fatal("Getwd failed")
	}
	t.Logf("Current dir: %s", dir)
}

func TestOSChmod(t *testing.T) {

}
