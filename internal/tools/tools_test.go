package tools

import (
	"strings"
	"testing"
)

func TestDir(t *testing.T) {
	res := Dir("c:\\Users\\Administrator\\Downloads\\logins\\@InfectLogs #2025 50.txt")

	if (!strings.HasSuffix(res, "logins")) {
		t.Fatal(res)
	}

	res = Dir("c:\\Users\\Administrator\\Downloads\\logins\\")

	if (!strings.HasSuffix(res, "logins")) {
		t.Fatal(res)
	}
}
