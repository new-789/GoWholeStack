package split_string

import (
	"reflect"
	"testing"
)

// testSplit4 测试组示例
/*
func TestSplit(t *testing.T) {
	type testCase struct {
		str string
		sep string
		want []string
	}
	// 声明测试用例组
	testGroup := []testCase{
		testCase{str: "babcbef", sep: "b", want: []string{"", "a", "c", "ef"}},
		testCase{str: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		testCase{str: "abcef", sep: "bc", want: []string{"a", "ef"}},
		testCase{str: "沙河有沙又有河", sep: "有", want: []string{"沙河", "沙又","河"}},
	}
	for _, value := range testGroup {
		got := Split(value.str, value.sep)
		want := value.want
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("want:%v got:%v\n", want, got)
		}
	}
}
*/

// TestSplit 子测试用例示例
func TestSplit(t *testing.T) {
	type testCase struct {
		str string
		sep string
		want []string
	}
	// 创建测试组: map 类型
	testGroup := map[string]testCase{
		"case_1": testCase{str: "babcbef", sep: "b", want: []string{"", "a", "c", "ef"}},
		"case_2": testCase{str: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"case_3": testCase{str: "abcef", sep: "bc", want: []string{"a", "ef"}},
		"case_4": testCase{str: "沙河有沙又有河", sep: "有", want: []string{"沙河", "沙又","河"}},
	}
	for k, v := range testGroup {
		// t.Run 子测试使用方法
		t.Run(k, func(t *testing.T) {
			got := Split(v.str, v.sep)
			if !reflect.DeepEqual(got, v.want) {
				t.Fatalf("name:%v want:%v got:%v\n",k, v.want, got)
			}
		})
	}
}

// 基准测试,测试命令参数：
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}

// 性能对比测试
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 1)
}

func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}

func BenchmarkFib3(b *testing.B) {
	benchmarkFib(b, 3)
}

func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 10)
}

func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}

// 测试命令总结
// 基本测试命令：go test
// 基本测试命令带详细信息：go test -v
// 测试覆盖率：go test -cover
// 测试覆盖率输出到文件：go test -cover -coverprofile=文件名
// 将测试覆盖率输出到文件的内容在浏览器显示：go tool cover -html=文件名
// 测试运行时间：go test -bench=需要测试的函数名
// 测试内存占用情况：go test -bench=需要测试的函数 -benchmem
// 并行测试：go test -bench=需要测试的函数名 -cpu=指定的cpu个数