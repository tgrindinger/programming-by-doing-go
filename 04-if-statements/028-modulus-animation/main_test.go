package main

import (
	"strings"
	"testing"
)

func TestDoAnimation(t *testing.T) {
	writer := &strings.Builder{}
	DoAnimation(44, writer)
	got := writer.String()
	want := "*****           \r *****          \r  *****         \r   *****        \r    *****       \r     *****      \r      *****     \r       *****    \r        *****   \r         *****  \r          ***** \r           *****\r          ***** \r         *****  \r        *****   \r       *****    \r      *****     \r     *****      \r    *****       \r   *****        \r  *****         \r *****          \r*****           \r *****          \r  *****         \r   *****        \r    *****       \r     *****      \r      *****     \r       *****    \r        *****   \r         *****  \r          ***** \r           *****\r          ***** \r         *****  \r        *****   \r       *****    \r      *****     \r     *****      \r    *****       \r   *****        \r  *****         \r *****          \r"
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}
