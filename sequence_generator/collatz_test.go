package sequence_generator_test

import (
	"collatz/sequence_generator"
	"testing"
)

func TestCollatzNextValue(t *testing.T) {

	t.Run("when 0 or less is passed, error is returned", func(t *testing.T) {

		want := "invalid value"
		_, have := sequence_generator.Next(0)

		if have == nil {
			t.Error("should have return a error")
		} else if want != have.Error() {
			t.Errorf("wanted %s, have %s", want, have.Error())
		}

	})

	t.Run("when 1 is passed, error is returned", func(t *testing.T) {

		want := "already finished"
		_, have := sequence_generator.Next(1)

		if have == nil {
			t.Error("should have return a error")
		} else if want != have.Error() {
			t.Errorf("wanted %s, have %s", want, have.Error())
		}

	})

	t.Run("test that next sequence from 2 is 1", func(t *testing.T) {
		want := 1
		have, err := sequence_generator.Next(2)

		if err != nil {
			t.Fatal(err)
		}

		if want != have {
			t.Errorf("wanted %d, have %d", want, have)
		}
	})

	t.Run("test that next sequence from 3 is 10", func(t *testing.T) {
		want := 10
		have, err := sequence_generator.Next(3)

		if err != nil {
			t.Fatal(err)
		}

		if want != have {
			t.Errorf("wanted %d, have %d", want, have)
		}
	})

	t.Run("when 4 is passed, the 2 is returned", func(t *testing.T) {
		want := 2
		have, err := sequence_generator.Next(4)

		if err != nil {
			t.Fatal(err)
		}

		if want != have {
			t.Errorf("wanted %d, have %d", want, have)
		}
	})
}

func TestCollatzSequenceGeneration(t *testing.T) {
	t.Run("when 2 is passed, then sequence is 1", func(t *testing.T) {
		want := []int{1}
		have, err := sequence_generator.Sequence(2)

		if err != nil {
			t.Fatal(err)
		}

		if !equal(want, have) {
			t.Errorf("wanted %d, have %d", want, have)
		}
	})

	t.Run("when 3 is passed, then sequence is 10, 5, 16, 8, 4, 2, 1", func(t *testing.T) {
		want := []int{10, 5, 16, 8, 4, 2, 1}
		have, err := sequence_generator.Sequence(3)

		if err != nil {
			t.Fatal(err)
		}

		if !equal(want, have) {
			t.Errorf("wanted %d, have %d", want, have)
		}
	})

	t.Run("when 4 is passed, then sequence is 2, 1", func(t *testing.T) {
		want := []int{2, 1}
		have, err := sequence_generator.Sequence(4)

		if err != nil {
			t.Fatal(err)
		}

		if !equal(want, have) {
			t.Errorf("wanted %d, have %d", want, have)
		}
	})
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
