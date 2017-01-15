// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package gogenerate

import (
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestFilesContaining(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	checks := []struct {
		d       string
		cmds    []string
		matches []string
	}{
		{"_testFiles/eg01", []string{"ls", "/bin/ls"}, []string{"a.go", "b.go", "c.go", "d.go"}},
	}

Checks:
	for _, c := range checks {

		path := filepath.Join(cwd, c.d)

		res, err := FilesContainingCmd(path, c.cmds...)
		if err != nil {
			t.Errorf("Got unexpected error find matches in %v: %v", c.d, err)
			continue Checks
		}

		if len(res) != len(c.matches) {
			t.Errorf("Matches not up to expectations: %v vs %v", res, c.matches)
			continue Checks
		}

		// just in case we were sloppy in the test table
		sort.Sort(byBase(c.matches))
		for i := range res {
			if res[i] != c.matches[i] {
				t.Errorf("Matches not up to expectations: %v vs %v", res, c.matches)
				continue Checks
			}
		}
	}
}

func TestNameFile(t *testing.T) {
	checks := []struct {
		n string
		c string
		r string
	}{
		{"a", "bananaGen", "gen_a_bananaGen.go"},
	}

	for _, c := range checks {
		r := NameFile(c.n, c.c)
		if r != c.r {
			t.Errorf("Expected NameFile(%q, %q) to be %v got %v", c.n, c.c, c.r, r)
		}
	}
}

func TestNameTestFile(t *testing.T) {
	checks := []struct {
		n string
		c string
		r string
	}{
		{"a", "bananaGen", "gen_a_bananaGen_test.go"},
	}

	for _, c := range checks {
		r := NameTestFile(c.n, c.c)
		if r != c.r {
			t.Errorf("Expected NameFile(%q, %q) to be %v got %v", c.n, c.c, c.r, r)
		}
	}
}

func TestFileIsGenerated(t *testing.T) {
	checks := []struct {
		p string
		r bool
	}{
		{"is_a_gen_file", false},
		{"gen_file", true},
		{"gen_file.go", true},
		{"genfile.go", false},
		{"/path/to/gen_file", true},
		{"/path/to/gen_file.go", true},
		{"/path/to/genfile.go", false},
		{"/", false},
	}

	for _, c := range checks {
		r := FileIsGenerated(c.p)
		if r != c.r {
			t.Errorf("Expected FileIsGenerated(%q) to be %v got %v", c.p, c.r, r)
		}
	}
}

func TestFileGeneratedBy(t *testing.T) {
	checks := []struct {
		n string
		c string
		r bool
	}{
		{"gen_bananaGen.go", "bananaGen", true},
		{"gen_bananaGen_test.go", "bananaGen", true},
		{"gen_a_bananaGen.go", "bananaGen", true},
		{"gen_a_bananaGen_test.go", "bananaGen", true},
		{"gen_abananaGen.go", "bananaGen", false},
		{"gen_", "bananaGen", false},
	}

	for _, c := range checks {
		r := FileGeneratedBy(c.n, c.c)
		if r != c.r {
			t.Errorf("Expected FileGeneratedBy(%q, %q) to be %v got %v", c.n, c.c, c.r, r)
		}
	}
}

func TestNameFileFromFile(t *testing.T) {
	checks := []struct {
		n  string
		o  string
		ok bool
	}{
		{"/path/to/a.txt", "", false},
		{"/path/to/a.go", "/path/to/gen_a_bananaGen.go", true},
		{"path/to/a.go", "path/to/gen_a_bananaGen.go", true},
		{"a.go", "gen_a_bananaGen.go", true},
		{"/path/to/a_test.go", "/path/to/gen_a_bananaGen_test.go", true},
		{"path/to/a_test.go", "path/to/gen_a_bananaGen_test.go", true},
		{"a_test.go", "gen_a_bananaGen_test.go", true},
	}

	for _, c := range checks {
		o, ok := NameFileFromFile(c.n, "bananaGen")
		if o != c.o || ok != c.ok {
			t.Errorf("Expected NameFileFromFile(%q) to be %v got %v", c.n, c.o, o)
		}
	}
}

func TestCommentLicenseHeader(t *testing.T) {
	fn := "_testFiles/licenseFile.txt"

	res, err := CommentLicenseHeader(&fn)

	if err != nil {
		t.Fatalf("CommentLicenseHeader failed when it should not have: %v", err)
	}

	exp := `// Copyright (c) Bananaman 2016
// Line 2

`

	if res != exp {
		t.Errorf("Actual output %q was not as expected %q", res, exp)
	}
}
