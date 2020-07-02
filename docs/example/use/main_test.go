package main

import (
	"fmt"
	"testing"

	"bou.ke/monkey"
	"github.com/roger-russel/smallbox/docs/example/use/box"
)

func Example_main() {
	main()
	// output:
	// simple
}

func Example_mainStubed() {

	// Clear all patchs at end
	defer monkey.UnpatchAll()

	// If you want to stub the function box.Get to return what you want.
	monkey.Patch(box.Get, func(name string) ([]byte, error) {
		return []byte("stubed"), nil
	})

	main()
	// output:
	// stubed
}

func Test_foo(t *testing.T) {

	t.Run("foo got boo", func(t *testing.T) {

		defer monkey.UnpatchAll()
		// If you want to stub the function box.Get to return what you want.
		monkey.Patch(box.Get, func(name string) ([]byte, error) {
			return []byte("boo"), nil
		})

		out := foo("boo")

		if out != "boo" {
			t.Errorf("Expecting boo but got: %v", out)
		}
	})

	t.Run("foo got panic", func(t *testing.T) {
		defer monkey.UnpatchAll()
		// If you want to stub the function box.Get to return what you want.
		monkey.Patch(box.Get, func(name string) ([]byte, error) {
			return []byte{}, fmt.Errorf("panic")
		})

		defer func() {
			if r := recover(); r != nil {
				if fmt.Sprint(r) != "panic" {
					t.Errorf("Expecting: %v, got: %v", "panic", fmt.Sprint(r))
				}
			}
		}()

		foo("panic")

	})

}
