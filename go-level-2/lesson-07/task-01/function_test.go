package main

import "testing"

func TestAssignErrors(t *testing.T) {
	var cases = []struct {
		name   string
		source map[string]any
		target any
	}{
		{
			name:   "nil target",
			source: make(map[string]any),
			target: nil,
		}, {
			name:   "not structure target",
			source: make(map[string]any),
			target: "not structure",
		}, {
			name:   "bad type of value",
			source: map[string]any{"Integer": "Not integer"},
			target: &struct{ Integer int }{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if err := assign(c.target, c.source); err == nil {
				t.Fatal("expeted error, got", c.target)
			}
		})
	}
}

func TestAssignHuman(t *testing.T) {
	var cases = []struct {
		name    string         // Название теста.
		isError bool           // Ожидается ошибка.
		source  map[string]any // Мапа со значениями.
		target  *Human         // Целевая структура.
		expect  *Human         // Проверочная структура.
	}{
		{
			name:   "only name",
			source: map[string]any{"Name": "Ivan"},
			target: &Human{},
			expect: &Human{Name: "Ivan"},
		}, {
			name:   "age and height",
			source: map[string]any{"Age": 14, "Height": 111.2},
			target: &Human{},
			expect: &Human{Age: 14, Height: 111.2},
		}, {
			name: "all fields",
			source: map[string]any{
				"Name":   "Peter",
				"Age":    14,
				"Height": 111.2,
				"Weight": 33.4,
			},
			target: &Human{},
			expect: &Human{
				Name:   "Peter",
				Age:    14,
				Height: 111.2,
				Weight: 33.4,
			},
		}, {
			name:    "bad age",
			isError: true,
			source:  map[string]any{"Age": "ABC"},
			target:  &Human{},
			expect:  &Human{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if err := assign(c.target, c.source); err == nil && c.isError {
				t.Error("expected error but didn't get one")
			} else if err != nil && !c.isError {
				t.Error("unexpected error:", err)
			} else if !c.target.IsEqualFields(c.expect) {
				t.Error("expeted:", c.expect, "got", c.target)
			}
		})
	}
}

func TestAssignStructure(t *testing.T) {
	t.Run("zero fields", func(t *testing.T) {
		var trg = struct{}{}
		var src = map[string]any{}

		if err := assign(&trg, src); err != nil {
			t.Error("unexpected error:", err)
		}
	})

	t.Run("nil-value field", func(t *testing.T) {
		var trg = struct{ String string }{}
		var src = map[string]any{"String": nil}

		if err := assign(&trg, src); err != nil {
			t.Error("unexpected error:", err)
		}
	})

	t.Run("int field", func(t *testing.T) {
		var trg = struct{ Integer int }{}
		var src = map[string]any{"Integer": int(123)}

		if err := assign(&trg, src); err != nil {
			t.Error("unexpected error:", err)
		} else if trg.Integer != src["Integer"] {
			t.Error("got:", trg.Integer, "expected:", src["Integer"])
		}
	})

	t.Run("float field", func(t *testing.T) {
		var trg = struct{ Float float32 }{}
		var src = map[string]any{"Float": float32(123.456)}

		if err := assign(&trg, src); err != nil {
			t.Error("unexpected error:", err)
		} else if trg.Float != src["Float"] {
			t.Error("got:", trg.Float, "expected:", src["Float"])
		}
	})

	t.Run("string field", func(t *testing.T) {
		var trg = struct{ String string }{}
		var src = map[string]any{"String": "ABC"}

		if err := assign(&trg, src); err != nil {
			t.Error("unexpected error:", err)
		} else if trg.String != src["String"] {
			t.Error("got:", trg.String, "expected:", src["String"])
		}
	})

	t.Run("bad field type", func(t *testing.T) {
		var trg = struct{ Integer int }{}
		var src = map[string]any{"Integer": "ABC"}

		if err := assign(&trg, src); err == nil {
			t.Error("expected error but didn't get one")
		}
	})
}
