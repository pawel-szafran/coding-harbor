package chess

import (
	"testing"

	"github.com/pawel-szafran/coding-harbor/chess/board"
)

func TestCountSafeBoards(t *testing.T) {
	tests := []struct {
		size      board.Size
		pieces    Pieces
		wantCount int
	}{
		{
			size:      board.Size{Rows: 3, Cols: 3},
			pieces:    Pieces{King: 2, Rook: 1},
			wantCount: 4,
		},
		{
			size:      board.Size{Rows: 4, Cols: 4},
			pieces:    Pieces{Rook: 2, Knight: 4},
			wantCount: 8,
		},
	}
	for _, tt := range tests {
		count := CountSafeBoards(tt.size, tt.pieces)
		if count != tt.wantCount {
			t.Errorf("Want %d, got %d", tt.wantCount, count)
		}
	}
}
