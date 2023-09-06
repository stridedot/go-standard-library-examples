package sort_test

import (
	"sort"
	"strings"
	"testing"
)

// 二分查找法查找并返回 cmp(i) <= 0 的最小索引
func TestSortFind(t *testing.T) {
	//str1 := []string{"foo"}
	str2 := []string{"ab", "ca"}
	tests := []struct {
		data      []string
		target    string
		wantPos   int
		wantFound bool
	}{
		//	{[]string{}, "foo", 0, false},
		//	{[]string{}, "", 0, false},

		//	{str1, "foo", 0, true},
		{str2, "ca", 0, false},
	}

	for _, tt := range tests {
		cmp := func(i int) int {
			t.Logf("tt.target = %v, tt.data[%d] = %v\n", tt.target, i, tt.data[i])
			res := strings.Compare(tt.target, tt.data[i])
			t.Logf("res = %v\n", res)
			return res
		}
		pos, found := sort.Find(len(tt.data), cmp)
		t.Logf("pos = %d, found = %t\n", pos, found)
	}
}

// sort.Float64s() 用于对 float64 类型的切片进行排序
func TestSortFloat64(t *testing.T) {
	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6}
	sort.Float64s(s)
	t.Logf("Sorted s = %v\n", s)
}

// sort.Float64sAreSorted() 用于判断 float64 类型的切片是否已经排序
func TestSortFlat64sAreSorted(t *testing.T) {
	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6}
	sort.Float64s(s)
	t.Logf("Flat64sAreSorted = %v\n", sort.Float64sAreSorted(s))
}

// sort.Ints() 用于对 int 类型的切片进行排序
func TestSortInts(t *testing.T) {
	s := []int{5, -1, 0, -3, 2}
	sort.Ints(s)
	t.Logf("Sorted s = %v\n", s)
}

// sort.IntsAreSorted() 用于判断 int 类型的切片是否已经排序
func TestSortIntsAreSorted(t *testing.T) {
	s := []int{5, -1, 0, -3, 2}
	sort.Ints(s)
	t.Logf("IntsAreSorted = %v\n", sort.IntsAreSorted(s))
}

// sort.IsSorted() 用于判断切片是否已经排序
func TestSortIsSorted(t *testing.T) {
	s := []int{5, -1, 0, -3, 2}
	t.Logf("IsSorted = %v\n", sort.IsSorted(sort.IntSlice(s)))
}

// sort.Search() 用于在已排序的切片中查找元素 x，返回 x 所在的索引
func TestSortSearch(t *testing.T) {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) && a[i] == x {
		t.Logf("Found %d at index %d in %v\n", x, i, a)
	} else {
		t.Logf("%d not found in %v\n", x, a)
	}
}

// sort.SearchFloat64s() 用于在已排序的 float64 切片中查找元素 x，返回 x 所在的索引
func TestSortSearchFloat64s(t *testing.T) {
	a := []float64{1.1, 3.3, 6.6, 10.1, 15.5, 21.2, 28.8, 36.3, 45.4, 55.5}
	x := 6.6

	i := sort.SearchFloat64s(a, x)
	if i < len(a) && a[i] == x {
		t.Logf("Found %f at index %d in %v\n", x, i, a)
	} else {
		t.Logf("%f not found in %v\n", x, a)
	}
}

// sort.SearchInts() 用于在已排序的 int 切片中查找元素 x，返回 x 所在的索引
func TestSortSearchInts(t *testing.T) {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i := sort.SearchInts(a, x)
	if i < len(a) && a[i] == x {
		t.Logf("Found %d at index %d in %v\n", x, i, a)
	} else {
		t.Logf("%d not found in %v\n", x, a)
	}
}

// sort.SearchStrings() 用于在已排序的 string 切片中查找元素 x，返回 x 所在的索引
func TestSortSearchStrings(t *testing.T) {
	a := []string{"a", "b", "c", "d", "e", "f", "g"}
	x := "c"

	i := sort.SearchStrings(a, x)
	if i < len(a) && a[i] == x {
		t.Logf("Found %s at index %d in %v\n", x, i, a)
	} else {
		t.Logf("%s not found in %v\n", x, a)
	}
}

// sort.Slice() 用于对切片进行排序
func TestSortSlice(t *testing.T) {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	t.Logf("Sorted by Name = %v\n", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	t.Logf("Sorted by Age = %v\n", people)
}

// sort.SliceIsSorted() 用于判断切片是否已经排序
func TestSortSliceIsSorted(t *testing.T) {
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 55},
		{"Bob", 75},
		{"Gopher", 7},
		{"Vera", 24},
	}
	t.Logf("IsSorted = %v\n", sort.SliceIsSorted(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	}))
}

// sort.SliceStable() 用于对切片进行稳定排序
func TestSortSliceStable(t *testing.T) {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	t.Logf("Sorted by Name = %v\n", people)

	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	t.Logf("Sorted by Age = %v\n", people)
}

// sort.Sort 用于对实现了 sort.Interface 接口的对象进行排序
func TestSortSort(t *testing.T) {
	d := []int{5, 2, 6, 3, 1, 4}
	sort.Sort(sort.IntSlice(d))
	t.Logf("sorted d = %v\n", d)
}

// sort.Stable() 用于对实现了 sort.Interface 接口的对象进行稳定排序
func TestSortStable(t *testing.T) {
	d := []int{5, 2, 6, 3, 1, 4}
	sort.Stable(sort.IntSlice(d))
	t.Logf("sorted d = %v\n", d)
}

// sort.Strings() 用于对 string 类型的切片进行排序
func TestSortStrings(t *testing.T) {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	t.Logf("Sorted s = %v\n", s)
}

// sort.StringsAreSorted() 用于判断 string 类型的切片是否已经排序
func TestSortStringsAreSorted(t *testing.T) {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	t.Logf("StringsAreSorted = %v\n", sort.StringsAreSorted(s))
}

// sort.Float64Slice() 用于对 float64 类型的切片进行排序
func TestSortFloat64Slice(t *testing.T) {
	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6}
	sort.Float64Slice(s).Sort()

	t.Logf("Sorted s = %v\n", s)
}

// sort.Reverse 用于对实现了 sort.Interface 接口的对象进行逆序排序
func TestSortReverse(t *testing.T) {
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	t.Logf("s = %v\n", s)
}

// sort.StringSlice() 用于对 string 类型的切片进行排序
func TestSortStringSlice(t *testing.T) {
	p := sort.StringSlice([]string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"})
	t.Logf("Len = %v\n", p.Len())
	t.Logf("Less = %v\n", p.Less(0, 1))
	t.Logf("Swap = %v\n", p.Search("Go"))

	p.Swap(0, 1)
	t.Logf("p = %v\n", p)

	p.Sort()
	t.Logf("p = %v\n", p)
}
