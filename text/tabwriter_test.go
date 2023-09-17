package text_test

import (
	"fmt"
	"os"
	"testing"
	"text/tabwriter"
)

// tabwriter.NewWriter
// minwidth 最小单元格宽度，包括任何 padding
// tabwidth 制表符的宽度（相当于空格数）
// padding 在计算单元格宽度之前添加到单元格的填充
func TestTabWriterNewWriter(t *testing.T) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\tc")
	fmt.Fprintln(w, "aa\tbb\tcc")
	fmt.Fprintln(w, "aaa\t") // 尾随制表符
	fmt.Fprintln(w, "aaaa\tdddd\teeee")
	w.Flush()
}

// tabwriter.Init
func TestTabwriterInit(t *testing.T) {
	w := new(tabwriter.Writer)

	// 使用 \t 分隔列格式，制表位为 8
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 0, '\t', tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\tc\td\t")
	fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t")
	fmt.Fprintln(w)
	w.Flush()

	// 在最小宽度为 5 的空格分隔列中设置右对齐格式
	// 和至少一个填充空白（因此更宽的列条目不会接触）。
	// 如果值长度大于 minwidth，单元格总长度 = minwidth + padding
	// 如果值长度小于 minwidth，单元格总长度 = minwidth
	w.Init(os.Stdout, 5, 0, 1, '-', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\tc\td\t.")
	fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t.")
	fmt.Fprintln(w)
	w.Flush()
}
