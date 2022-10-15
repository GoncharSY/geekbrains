package main

import "testing"

func TestIsEqualFields(t *testing.T) {
	var etalon = Human{
		Name:   "Ivanov Ivan",
		Age:    uint(32),
		Height: centimeter(177.123),
		Weight: kilogram(72.234),
	}

	var cases = []struct {
		name   string
		expect bool
		object Human
	}{
		{
			name:   "All fields are equal",
			expect: true,
			object: Human{
				Name:   "Ivanov Ivan",
				Age:    uint(32),
				Height: centimeter(177.123),
				Weight: kilogram(72.234),
			},
		}, {
			name:   "Name isn't equal",
			expect: false,
			object: Human{
				Name:   "Petrov Petr",
				Age:    uint(32),
				Height: centimeter(177.123),
				Weight: kilogram(72.234),
			},
		}, {
			name:   "Age isn't equal",
			expect: false,
			object: Human{
				Name:   "Ivanov Ivan",
				Age:    uint(22),
				Height: centimeter(177.123),
				Weight: kilogram(72.234),
			},
		}, {
			name:   "Height isn't equal",
			expect: false,
			object: Human{
				Name:   "Ivanov Ivan",
				Age:    uint(32),
				Height: centimeter(177),
				Weight: kilogram(72.234),
			},
		}, {
			name:   "Weight isn't equal",
			expect: false,
			object: Human{
				Name:   "Ivanov Ivan",
				Age:    uint(32),
				Height: centimeter(177.123),
				Weight: kilogram(72),
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if result := etalon.IsEqualFields(&c.object); result != c.expect {
				t.Fatal("got:", result, "expected:", c.expect)
			}
		})
	}
}
