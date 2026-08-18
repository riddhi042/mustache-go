package scanner

import "testing"

func TestEOS(t *testing.T) {
	s := New("hello")

	if s.EOS() {
		t.Fatal("expected EOS() to be false")
	}

	s.Scan("hello")

	if !s.EOS() {
		t.Fatal("expected EOS() to be true")
	}
}

func TestScan(t *testing.T) {
	s := New("{{name}}")

	got := s.Scan("{{")

	if got != "{{" {
		t.Fatalf("got %q, want %q", got, "{{")
	}
}

func TestScanUntil(t *testing.T) {
	s := New("Hello {{name}}")

	got := s.ScanUntil("{{")

	if got != "Hello " {
		t.Fatalf("got %q, want %q", got, "Hello ")
	}
}

func TestScanFailure(t *testing.T) {
	s := New("Hello")

	got := s.Scan("{{")

	if got != "" {
		t.Fatalf("got %q, want empty string", got)
	}
}

func TestScanUntilNotFound(t *testing.T) {
	s := New("Hello")

	got := s.ScanUntil("{{")

	if got != "Hello" {
		t.Fatalf("got %q, want %q", got, "Hello")
	}

	if !s.EOS() {
		t.Fatal("expected EOS after scanning entire string")
	}
}
