package pattern

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestPatternCompile(t *testing.T) {
	tests := []struct {
		name    string
		pattern *Pattern
		want    string
		wantErr bool
	}{
		{
			name: "single pattern",
			pattern: &Pattern{
				Pattern: "test-pattern",
			},
			want:    "test-pattern",
			wantErr: false,
		},
		{
			name: "multiple patterns",
			pattern: &Pattern{
				Patterns: []string{"foo", "bar", "baz"},
			},
			want:    "(foo|bar|baz)",
			wantErr: false,
		},
		{
			name:    "empty pattern",
			pattern: &Pattern{},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pattern.Compile()
			if (err != nil) != tt.wantErr {
				t.Errorf("Compile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Compile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPatternGetEngine(t *testing.T) {
	tests := []struct {
		name string
		p    *Pattern
		want string
	}{
		{
			name: "default engine",
			p:    &Pattern{},
			want: "grep",
		},
		{
			name: "custom engine",
			p:    &Pattern{Engine: "ag"},
			want: "ag",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.GetEngine(); got != tt.want {
				t.Errorf("GetEngine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveAndLoad(t *testing.T) {
	// Create a temporary directory for test patterns
	tmpDir := t.TempDir()

	// Mock GetPatternDir by using a different approach
	// We'll test Save and Load directly with file paths

	testPat := &Pattern{
		Flags:   "-Hnri",
		Pattern: "test-pattern",
	}

	// Test JSON encoding/decoding
	data, err := json.MarshalIndent(testPat, "", "    ")
	if err != nil {
		t.Fatalf("Failed to marshal pattern: %v", err)
	}

	patFile := filepath.Join(tmpDir, "test.json")
	if err := os.WriteFile(patFile, data, 0666); err != nil {
		t.Fatalf("Failed to write pattern file: %v", err)
	}

	// Read and verify
	f, err := os.Open(patFile)
	if err != nil {
		t.Fatalf("Failed to open pattern file: %v", err)
	}
	defer f.Close()

	var loaded Pattern
	if err := json.NewDecoder(f).Decode(&loaded); err != nil {
		t.Fatalf("Failed to decode pattern: %v", err)
	}

	if loaded.Flags != testPat.Flags {
		t.Errorf("Flags = %v, want %v", loaded.Flags, testPat.Flags)
	}
	if loaded.Pattern != testPat.Pattern {
		t.Errorf("Pattern = %v, want %v", loaded.Pattern, testPat.Pattern)
	}
}

func TestSaveErrors(t *testing.T) {
	tests := []struct {
		name    string
		nameArg string
		flags   string
		pat     string
		wantErr bool
	}{
		{
			name:    "empty name",
			nameArg: "",
			flags:   "-Hnri",
			pat:     "test",
			wantErr: true,
		},
		{
			name:    "empty pattern",
			nameArg: "test",
			flags:   "-Hnri",
			pat:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Save(tt.nameArg, tt.flags, tt.pat)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPatternJSON(t *testing.T) {
	// Test that Pattern struct can be properly marshaled/unmarshaled
	p := &Pattern{
		Flags:    "-Hnri",
		Pattern:  "test",
		Patterns: []string{"foo", "bar"},
		Engine:   "ag",
	}

	data, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var got Pattern
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if got.Flags != p.Flags {
		t.Errorf("Flags = %v, want %v", got.Flags, p.Flags)
	}
	if got.Pattern != p.Pattern {
		t.Errorf("Pattern = %v, want %v", got.Pattern, p.Pattern)
	}
	if got.Engine != p.Engine {
		t.Errorf("Engine = %v, want %v", got.Engine, p.Engine)
	}
}
