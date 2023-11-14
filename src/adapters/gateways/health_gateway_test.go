package gateways

import "testing"

func TestPing(t *testing.T) {
	g := NewHealthGateway(dbconn)
	t.Run("success", func(t *testing.T) {
		err := g.Ping()
		if err != nil {
			t.Errorf("got %v, want nil", err)
		}
	})
}
