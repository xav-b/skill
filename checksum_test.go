package main

import "testing"

func TestMd5Checksum(t *testing.T) {
  var contentTests = []struct {
    content  string  // input
    expected string  // expected result
  }{
    {"foo", "acbd18db4cc2f85cedef654fccc4a4d8"},
    {"bar", "37b51d194a7513e45b56f6524f2d51f2"},
  }

  for _, tt := range contentTests {
    actual := Checksum([]byte(tt.content), "MD5")
    if actual != tt.expected {
      t.Errorf("Checksum(%s, 'MD5'): expected %s, actual %s", tt.content, tt.expected, actual)
    }
  }
}
