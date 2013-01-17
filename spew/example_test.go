/*
 * Copyright (c) 2013 Dave Collins <dave@davec.name>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package spew_test

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type Flag int

const (
	flagOne Flag = iota
	flagTwo
)

var flagStrings = map[Flag]string{
	flagOne: "flagOne",
	flagTwo: "flagTwo",
}

func (f Flag) String() string {
	if s, ok := flagStrings[f]; ok {
		return s
	}
	return fmt.Sprintf("Unknown flag (%d)", int(f))
}

type Bar struct {
	flag Flag
	data uintptr
}

type Foo struct {
	unexportedField Bar
	ExportedField   map[interface{}]interface{}
}

// This example demonstrates how to use Dump to dump variables to stdout.
func ExampleDump() {
	// The following package level declarations are assumed for this example:
	/*
		type Flag int

		const (
			flagOne Flag = iota
			flagTwo
		)

		var flagStrings = map[Flag]string{
			flagOne: "flagOne",
			flagTwo: "flagTwo",
		}

		func (f Flag) String() string {
			if s, ok := flagStrings[f]; ok {
				return s
			}
			return fmt.Sprintf("Unknown flag (%d)", int(f))
		}

		type Bar struct {
			flag Flag
			data uintptr
		}

		type Foo struct {
			unexportedField Bar
			ExportedField   map[interface{}]interface{}
		}
	*/

	// Setup some sample data structures for the example.
	bar := Bar{Flag(flagTwo), uintptr(0)}
	s1 := Foo{bar, map[interface{}]interface{}{"one": true}}
	f := Flag(5)

	// Dump!
	spew.Dump(s1, f)

	// Output:
	// (spew_test.Foo) {
	//  unexportedField: (spew_test.Bar) {
	//   flag: (spew_test.Flag) flagTwo,
	//   data: (uintptr) <nil>
	//  },
	//  ExportedField: (map[interface {}]interface {}) {
	//   (string) "one": (bool) true
	//  }
	// }
	// (spew_test.Flag) Unknown flag (5)
	//
}

// This example demonstrates how to use Printf to display a variable with a
// format string and inline formatting.
func ExamplePrintf() {
	// Create a double pointer to a uint 8.
	ui8 := uint8(5)
	pui8 := &ui8
	ppui8 := &pui8

	// Create a circular data type.
	type circular struct {
		ui8 uint8
		c   *circular
	}
	c := circular{ui8: 1}
	c.c = &c

	// Print!
	spew.Printf("ppui8: %v\n", ppui8)
	spew.Printf("circular: %v\n", c)

	// Output:
	// ppui8: <**>5
	// circular: {1 <*>{1 <*><shown>}}
}

// This example demonstrates how to use a SpewState.
func ExampleSpewState() {
	// A SpewState does not need initialization.
	ss := new(spew.SpewState) // or var ss spew.SpewState

	// Modify the indent level of the SpewState only.  The global configuration
	// is not modified.
	ssc := ss.Config()
	ssc.Indent = "\t"

	// Output using the SpewState instance.
	v := map[string]int{"one": 1}
	ss.Printf("v: %v\n", v)
	ss.Dump(v)

	// Output:
	// v: map[one:1]
	// (map[string]int) {
	// 	(string) "one": (int) 1
	// }
}

// This example demonstrates how to use a SpewState.Dump to dump variables to
// stdout
func ExampleSpewState_Dump() {
	// See the top-level Dump example for details on the types used in this
	// example.

	// A SpewState does not need initialization.
	ss := new(spew.SpewState)  // or var ss spew.SpewState
	ss2 := new(spew.SpewState) // or var ss2 spew.SpewState

	// Modify the indent level of the first SpewState only.
	ssc := ss.Config()
	ssc.Indent = "\t"

	// Setup some sample data structures for the example.
	bar := Bar{Flag(flagTwo), uintptr(0)}
	s1 := Foo{bar, map[interface{}]interface{}{"one": true}}

	// Dump using the SpewState instances.
	ss.Dump(s1)
	ss2.Dump(s1)

	// Output:
	// (spew_test.Foo) {
	// 	unexportedField: (spew_test.Bar) {
	// 		flag: (spew_test.Flag) flagTwo,
	// 		data: (uintptr) <nil>
	// 	},
	// 	ExportedField: (map[interface {}]interface {}) {
	//		(string) "one": (bool) true
	// 	}
	// }
	// (spew_test.Foo) {
	//  unexportedField: (spew_test.Bar) {
	//   flag: (spew_test.Flag) flagTwo,
	//   data: (uintptr) <nil>
	//  },
	//  ExportedField: (map[interface {}]interface {}) {
	//   (string) "one": (bool) true
	//  }
	// }
	//
}

// This example demonstrates how to use SpewState.Printf to display a variable
// with a format string and inline formatting.
func ExampleSpewState_Printf() {
	// See the top-level Dump example for details on the types used in this
	// example.

	// A SpewState does not need initialization.
	ss := new(spew.SpewState)  // or var ss spew.SpewState
	ss2 := new(spew.SpewState) // or var ss2 spew.SpewState

	// Modify the method handling of the first SpewState only.
	ssc := ss.Config()
	ssc.DisableMethods = true

	// This is of type Flag which implements a Stringer and has raw value 1.
	f := flagTwo

	// Dump using the SpewState instances.
	ss.Printf("f: %v\n", f)
	ss2.Printf("f: %v\n", f)

	// Output:
	// f: 1
	// f: flagTwo
}
