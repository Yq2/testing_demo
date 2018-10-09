package test

import (
	"fmt"
	"testing"
)

func print1to20() int {
	res := 0
	for i := 1; i <= 20; i++ {
		res += i
	}
	return res
}

func TestPrint2(t *testing.T) {
	//t.Run可以控制功能测试case的执行顺序

	t.Run("a1", func(t *testing.T) {
		fmt.Println("a1")
	})
	t.Run("a2", func(t *testing.T) {
		fmt.Println("a2")
	})
	t.Run("a3", func(t *testing.T) {
		fmt.Println("a3")
	})
}

func TestPrint1(t *testing.T) {
	res := print1to20()
	if res != 210 {
		t.Errorf("Return value not valid")
	}
	fmt.Println("TestPrint1 h1")
	t.SkipNow()
	fmt.Println("TestPrint1 h2")
}

func TestAll(t *testing.T) {
	t.Run("TestPrint1", TestPrint1)
	//t.Run("TestPrint2", TestPrint2)
}

func TestMain(m *testing.M) {
	fmt.Println("Tests begings...")
	m.Run()
}

//一个test case中只能有一个，且前缀相同也不行
//func TestMain2(m *testing.M) {
//
//}

//执行时间不稳定的测试函数
func aaa(n int) int {
	for n>0 {
		n--
	}
	return n
}

//go test -bench=.  执行当前测试文件所有基准测试case
func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		print1to20()

		//aaa(n)  //会导致当前基准测试卡死
	}
}
