package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/buildbarn/go-xdr/pkg/compiler/model"
	"github.com/buildbarn/go-xdr/pkg/compiler/parser"
)

type simpleSink struct {
	w *bufio.Writer
}

type typeInfo struct {
	pkg string
	t   model.Type
}

type simpleResolver struct {
	constants      map[string]*big.Int
	types          map[string]typeInfo
	pendingImports []string
}

func (r *simpleResolver) RegisterConstant(name string, c *big.Int) error {
	if _, ok := r.constants[name]; ok {
		return fmt.Errorf("Constant defined twice: %#v", name)
	}
	r.constants[name] = c
	return nil
}

func (r *simpleResolver) RegisterType(name, pkg string, t model.Type) error {
	if _, ok := r.types[name]; ok {
		return fmt.Errorf("Type defined twice: %#v", name)
	}
	r.types[name] = typeInfo{
		pkg: pkg,
		t:   t,
	}
	return nil
}

func (r *simpleResolver) RegisterImport(path string) error {
	r.pendingImports = append(r.pendingImports, path)
	return nil
}

func (r *simpleResolver) ResolveConstant(name string) (*big.Int, error) {
	c, ok := r.constants[name]
	if !ok {
		return nil, fmt.Errorf("Unknown constant: %#v", name)
	}
	return c, nil
}

func (r *simpleResolver) ResolveType(name string) (string, model.Type, error) {
	ti, ok := r.types[name]
	if !ok {
		return "", nil, fmt.Errorf("Unknown type: %#v", name)
	}
	return ti.pkg, ti.t, nil
}

type errorDetector struct {
	base      antlr.ErrorListener
	gotErrors bool
}

func (ed *errorDetector) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	ed.gotErrors = true
	ed.base.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)
}

func (ed *errorDetector) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	ed.gotErrors = true
	ed.base.ReportAmbiguity(recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

func (ed *errorDetector) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	ed.gotErrors = true
	ed.ReportAttemptingFullContext(recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

func (ed *errorDetector) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	ed.gotErrors = true
	ed.ReportContextSensitivity(recognizer, dfa, startIndex, stopIndex, prediction, configs)
}

type stringList []string

func (sl *stringList) String() string {
	return strings.Join(*sl, ":")
}

func (sl *stringList) Set(v string) error {
	*sl = append(*sl, v)
	return nil
}

func main() {
	var importPaths stringList
	flag.Var(&importPaths, "I", "Add search path for source files")
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("Usage: [-I import_path] in.x out.go")
	}

	r := simpleResolver{
		constants:      map[string]*big.Int{},
		types:          map[string]typeInfo{},
		pendingImports: []string{args[0]},
	}

	processedImports := map[string]struct{}{}
	var inputTopLevel *model.TopLevel
	for len(r.pendingImports) > 0 {
		path := r.pendingImports[0]
		r.pendingImports = r.pendingImports[1:]

		// Prevent duplicate processing of the same source file.
		if _, ok := processedImports[path]; ok {
			continue
		}
		processedImports[path] = struct{}{}

		var input *antlr.FileStream
		for _, importPath := range importPaths {
			var err error
			input, err = antlr.NewFileStream(filepath.Join(importPath, path))
			if err == nil {
				goto FoundSourceFile
			}
			if !os.IsNotExist(err) {
				log.Fatalf("Failed to open file %#v: %s", path, err)
			}
		}
		log.Fatalf("File %#v not found in any of the import paths", path)
	FoundSourceFile:
		p := parser.NewXDRParser(
			antlr.NewCommonTokenStream(
				parser.NewXDRLexer(input),
				0))

		errorDetector := errorDetector{
			base: antlr.NewDiagnosticErrorListener(true),
		}
		p.AddErrorListener(&errorDetector)
		topLevel := p.TopLevel().GetV()
		if errorDetector.gotErrors {
			os.Exit(1)
		}
		if inputTopLevel == nil {
			inputTopLevel = &topLevel
		}

		if err := topLevel.Register(&r); err != nil {
			log.Fatal(err)
		}
	}

	// Emit the resulting Go source file.
	f, err := os.OpenFile(args[1], os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		log.Fatal("Failed to open output file: ", err)
	}

	w := bufio.NewWriter(f)
	if err := inputTopLevel.Emit(w, &r); err != nil {
		log.Fatal(err)
	}
	if err := w.Flush(); err != nil {
		log.Fatal("Failed to write to output file: ", err)
	}
}
