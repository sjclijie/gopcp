package main

import "testing"

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {

	t.Logf("hello %s\n", checkMark)

	t.Errorf("aaaaa %s\n", ballotX)

}
