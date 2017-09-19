package trie

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type Node struct {
	Key   string
	Term  string
	Val   string
	Child *Trie
}

type Trie map[string]Node

const space = " "

func New() *Trie {
	t := make(Trie)
	return &t
}

func Split(k string) []string {
	var words []string
	var word string

	for _, r := range k {
		if unicode.IsSpace(r) {
			if len(word) > 0 {
				words = append(words, strings.ToLower(word))
				word = ""
			}
		} else {
			word += string(r)
		}
	}
	if len(word) > 0 {
		words = append(words, strings.ToLower(word))
	}
	// fmt.Println("key=[" + strings.Join(words, " ") + "]")
	return words
}

func (t *Trie) Insert(k string, v string) error {
	return t.insert(k, k, v)
}

func (t *Trie) insert(k string, term string, v string) error {
	if t == nil {
		return NullPointer
	}

	keys := Split(k)
	k = keys[0]

	if len(keys) == 1 {
		if node, ok := (*t)[k]; ok {
			node.Term = term
			node.Val = v
		} else {
			(*t)[k] = Node{k, term, v, New()}
		}
		return nil
	}

	key := strings.Join(keys[1:], space)

	if _, ok := (*t)[k]; !ok {
		(*t)[k] = Node{key, "", "", New()}
	}
	c := (*t)[k].Child
	return c.insert(key, term, v)
}

func (t *Trie) Value(k string) (string, error) {
	if t == nil {
		return "", NullPointer
	}

	keys := Split(k)

	if len(keys) == 0 {
		return "", KeyNotFound
	}
	if len(keys) == 1 {
		k = keys[0]
		if node, ok := (*t)[k]; ok {
			return node.Val, nil
		} else {
			return "", KeyNotFound
		}
	}

	k = keys[0]
	if node, ok := (*t)[k]; ok {
		return node.Child.Value(strings.Join(keys[1:], space))
	}
	return "", KeyNotFound
}

func (t *Trie) TermValue(k string) (string, string, error) {
	if t == nil {
		return "", "", NullPointer
	}

	keys := Split(k)

	if len(keys) == 0 {
		return "", "", KeyNotFound
	}
	if len(keys) == 1 {
		k = keys[0]
		if node, ok := (*t)[k]; ok {
			return node.Term, node.Val, nil
		} else {
			return "", "", KeyNotFound
		}
	}

	k = keys[0]
	if node, ok := (*t)[k]; ok {
		return node.Child.TermValue(strings.Join(keys[1:], space))
	}
	return "", "", KeyNotFound
}

func (t *Trie) Child(k string) (*Trie, error) {
	if t == nil {
		return nil, NullPointer
	}

	keys := Split(k)

	if len(k) == 0 {
		return nil, KeyNotFound
	}
	if len(keys) == 1 {
		k = keys[0]
		if node, ok := (*t)[k]; ok {
			return node.Child, nil
		} else {
			return nil, KeyNotFound
		}
	}

	k = keys[0]
	if node, ok := (*t)[k]; ok {
		return node.Child.Child(strings.Join(keys[1:], space))
	}
	return nil, KeyNotFound
}

func (t *Trie) Traverse(initial string, suffix string) {
	for k, n := range *t {
		// fmt.Printf("%s[%s], {[%s] = [%s]}\n", indent, k, n.Key, n.Value)
		fmt.Printf("%s[%s] = [%s]\n", initial, k, n.Val)
		if n.Child != nil {
			n.Child.Traverse(initial+suffix, suffix)
		}
	}
}

var KeyNotFound = errors.New("Key Not Found")
var NullPointer = errors.New("Null Pointer")
