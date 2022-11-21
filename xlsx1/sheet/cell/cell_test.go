package cell

import "testing"

func TestTen2Col(t *testing.T) {
	if "A" != ten2col(1) {
		t.Fatal("测试失败")
	}

	if "Z" != ten2col(26) {
		t.Fatal("测试失败")
	}

	if "AA" != ten2col(27) {
		t.Fatal("测试失败")
	}

	if "BA" != ten2col(53) {
		t.Fatal("测试失败")
	}
}
