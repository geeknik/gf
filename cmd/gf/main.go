package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/geeknik/gf/internal/util"
	"github.com/geeknik/gf/pkg/pattern"
)

func main() {
	var saveMode bool
	flag.BoolVar(&saveMode, "save", false, "save a pattern (e.g: gf -save pat-name -Hnri 'search-pattern')")

	var listMode bool
	flag.BoolVar(&listMode, "list", false, "list available patterns")

	var dumpMode bool
	flag.BoolVar(&dumpMode, "dump", false, "prints the grep command rather than executing it")

	flag.Parse()

	if listMode {
		pats, err := pattern.List()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}

		fmt.Println(strings.Join(pats, "\n"))
		return
	}

	if saveMode {
		name := flag.Arg(0)
		flags := flag.Arg(1)
		pat := flag.Arg(2)

		err := pattern.Save(name, flags, pat)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
		return
	}

	patName := flag.Arg(0)
	files := flag.Arg(1)
	if files == "" {
		files = "."
	}

	pat, err := pattern.Load(patName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	searchPat, err := pat.Compile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", patName, err)
		return
	}

	if dumpMode {
		fmt.Printf("grep %v %q %v\n", pat.Flags, searchPat, files)
		return
	}

	engine := pat.GetEngine()

	var cmd *exec.Cmd
	if util.StdinIsPipe() {
		cmd = exec.Command(engine, pat.Flags, searchPat)
	} else {
		cmd = exec.Command(engine, pat.Flags, searchPat, files)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}
