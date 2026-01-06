package pattern

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Pattern represents a gf search pattern configuration
type Pattern struct {
	Flags    string   `json:"flags,omitempty"`
	Pattern  string   `json:"pattern,omitempty"`
	Patterns []string `json:"patterns,omitempty"`
	Engine   string   `json:"engine,omitempty"`
}

// GetPatternDir returns the directory where pattern files are stored
func GetPatternDir() (string, error) {
	usr, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(usr, ".config/gf")
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return path, nil
	}
	return filepath.Join(usr, ".gf"), nil
}

// Load loads a pattern from a file
func Load(name string) (*Pattern, error) {
	patDir, err := GetPatternDir()
	if err != nil {
		return nil, fmt.Errorf("unable to open user's pattern directory: %w", err)
	}

	filename := filepath.Join(patDir, name+".json")
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("no such pattern")
	}
	defer f.Close()

	var p Pattern
	dec := json.NewDecoder(f)
	if err := dec.Decode(&p); err != nil {
		return nil, fmt.Errorf("pattern file '%s' is malformed: %w", filename, err)
	}

	return &p, nil
}

// Compile returns the compiled pattern string
func (p *Pattern) Compile() (string, error) {
	if p.Pattern != "" {
		return p.Pattern, nil
	}

	if len(p.Patterns) == 0 {
		return "", errors.New("pattern contains no pattern(s)")
	}

	return "(" + strings.Join(p.Patterns, "|") + ")", nil
}

// Engine returns the search engine to use (default: grep)
func (p *Pattern) GetEngine() string {
	if p.Engine != "" {
		return p.Engine
	}
	return "grep"
}

// Save saves a pattern to a file
func Save(name, flags, pat string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	if pat == "" {
		return errors.New("pattern cannot be empty")
	}

	p := &Pattern{
		Flags:   flags,
		Pattern: pat,
	}

	patDir, err := GetPatternDir()
	if err != nil {
		return fmt.Errorf("failed to determine pattern directory: %w", err)
	}

	if err := os.MkdirAll(patDir, 0755); err != nil {
		return fmt.Errorf("failed to create pattern directory: %w", err)
	}

	path := filepath.Join(patDir, name+".json")
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return fmt.Errorf("failed to create pattern file: %w", err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "    ")

	if err := enc.Encode(p); err != nil {
		return fmt.Errorf("failed to write pattern file: %w", err)
	}

	return nil
}

// List returns all available pattern names
func List() ([]string, error) {
	out := []string{}

	patDir, err := GetPatternDir()
	if err != nil {
		return out, fmt.Errorf("failed to determine pattern directory: %w", err)
	}

	files, err := filepath.Glob(filepath.Join(patDir, "*.json"))
	if err != nil {
		return out, err
	}

	for _, f := range files {
		name := f[len(patDir)+1 : len(f)-5]
		out = append(out, name)
	}

	return out, nil
}
