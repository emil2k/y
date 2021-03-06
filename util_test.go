package main

import (
	"go/build"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"testing"
)

// testBuild tests that the package in the passed directory builds without
// errors and returns the package.
func testBuild(t *testing.T, dir string) *build.Package {
	pkg, err := build.ImportDir(dir, 0)
	if err != nil {
		t.Errorf("error during build : %s", err.Error())
	}
	return pkg
}

// testImports tests that the package in the passed directory has the expected
// imports. Includes imports from test files and external test files based on
// `includeTests` bool.
func testImports(t *testing.T, dir string, imports []string, includeTests bool) {
	pkg := testBuild(t, dir)
	gotImports := getImports(pkg, includeTests) // sorts it already
	sort.Strings(imports)
	if !reflect.DeepEqual(gotImports, imports) {
		t.Errorf("imports not equal : got %v, expected %v", gotImports, imports)
	}
}

// getTestContextCopy creates a temporary directory, copies the contents from
// src into it, and returns a context with it as set as the GOPATH.
// In case of error immediately failst the test.
func getTestContextCopy(t *testing.T, src string) *build.Context {
	ctx := getTestContext(t)
	if err := copyDir(src, ctx.GOPATH, true); err != nil {
		t.Errorf("error while copying GOPATH : %s", err.Error())
		t.FailNow()
	}
	return ctx
}

// getTestContext creates a temporary directory and returns a context with it as
// set as the GOPATH.
// In case of error immediately fails the test.
func getTestContext(t *testing.T) *build.Context {
	dst, err := ioutil.TempDir(os.TempDir(), "vendtest")
	if err != nil {
		t.Errorf("error getting temporary directory : %s", err.Error())
		t.FailNow()
	}
	ctx := build.Default
	ctx.GOPATH = dst
	return &ctx
}
