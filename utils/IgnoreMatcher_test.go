package utils

import "testing"

func TestIgnoreMatcherDefaults(t *testing.T) {
	matcher := NewIgnoreMatcher(t.TempDir())

	tests := []struct {
		path  string
		isDir bool
	}{
		{"__pycache__", true},
		{"src/__pycache__", true},
		{"node_modules", true},
		{"app/dist", true},
		{"src/main.pyc", false},
		{"assets/logo.png", false},
	}

	for _, tt := range tests {
		if !matcher.ShouldIgnore(tt.path, tt.isDir) {
			t.Fatalf("expected %q to be ignored", tt.path)
		}
	}
}

func TestIgnoreMatcherCustomPatterns(t *testing.T) {
	matcher := &IgnoreMatcher{}
	matcher.addRule("tmp/")
	matcher.addRule("*.log")
	matcher.addRule("/root-only.txt")
	matcher.addRule("!keep.log")

	if !matcher.ShouldIgnore("src/tmp", true) {
		t.Fatal("expected nested tmp directory to be ignored")
	}
	if !matcher.ShouldIgnore("logs/app.log", false) {
		t.Fatal("expected log file to be ignored")
	}
	if matcher.ShouldIgnore("keep.log", false) {
		t.Fatal("expected negated file to be allowed")
	}
	if !matcher.ShouldIgnore("root-only.txt", false) {
		t.Fatal("expected root anchored file to be ignored at root")
	}
	if matcher.ShouldIgnore("src/root-only.txt", false) {
		t.Fatal("did not expect root anchored file to be ignored in nested directory")
	}
}
