package tree_sitter_dot_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-dot"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_dot.Language())
	if language == nil {
		t.Errorf("Error loading Dot grammar")
	}
}
