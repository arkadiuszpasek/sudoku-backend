package sudoku

import "testing"

func TestSerialize(t *testing.T) {
	board := BoardValues{
		{0, 1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8, 0},
		{2, 3, 4, 5, 6, 7, 8, 0, 1},
		{3, 4, 5, 6, 7, 8, 0, 1, 2},
		{4, 5, 6, 7, 8, 0, 1, 2, 3},
		{5, 6, 7, 8, 0, 1, 2, 3, 4},
		{6, 7, 8, 0, 1, 2, 3, 4, 5},
		{7, 8, 0, 1, 2, 3, 4, 5, 6},
		{8, 0, 1, 2, 3, 4, 5, 6, 7},
	}
	expected := "012345678123456780234567801345678012456780123567801234678012345780123456801234567"

	got := Serialize(board)

	if got != expected {
		t.Errorf("Serialize() = %q, want %q", got, expected)
	}
}

func TestDeserialize(t *testing.T) {
	serialized := "012345678123456780234567801345678012456780123567801234678012345780123456801234567"

	board, err := Deserialize(serialized)
	if err != nil {
		t.Fatalf("Deserialize() unexpected error: %v", err)
	}

	got := Serialize(board)
	if got != serialized {
		t.Errorf("Deserialize() board serializes to %q, want %q", got, serialized)
	}
}

func TestDeserializeInvalidLength(t *testing.T) {
	serialized := "0123454521321"

	_, err := Deserialize(serialized)
	if err == nil {
		t.Fatal("Deserialize() expected error for invalid length")
	}
}

func TestGenerateBoard(t *testing.T) {
	expected := BoardValues{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	expectedSerialized := Serialize(expected)

	got := Serialize(GenerateBoard())

	if got != expectedSerialized {
		t.Errorf("GenerateBoard() serializes to %q, want %q", got, expectedSerialized)
	}
}
