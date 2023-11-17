package calc

import "testing"

func TestRandom(t *testing.T) {
	n := 10000
	for i := 0; i < n; i++ {
		t.Run("random", func(t *testing.T) {
			t.Parallel()
			if data := random(0.9, 1.1); data < 0.9 || data > 1.1 {
				t.Errorf("random() = %f, want 0.9 <= data <= 1.1", data)
			}
		})
	}
}
