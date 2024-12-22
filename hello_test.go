package hello

import (
	"testing"
)

// TestHello は Hello 関数をテストします
func TestHello(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Hello() panicked: %v", r)
			}
		}()
		Hello()
	}()
}
