package leetcode_go

/*
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

Trie() Initializes the trie object.
void insert(String word) Inserts the string word into the trie.
boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.


Example 1:

Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True
*/

type Trie struct {
	children map[string]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{
		children: make(map[string]*Trie),
		isEnd:    false,
	}
}

func (this *Trie) Insert(word string) {
	root := this
	for _, c := range word {
		if _, ok := root.children[string(c)]; !ok {
			root.children[string(c)] = &Trie{children: make(map[string]*Trie)}
		}
		root = root.children[string(c)]
	}
	root.isEnd = true
}

func (this *Trie) Search(word string) bool {
	root := this.SearchPrefix(word)
	return root != nil && root.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}

func (this *Trie) SearchPrefix(prefix string) *Trie {
	root := this
	for _, ch := range prefix {
		if root.children[string(ch)] == nil {
			return nil
		}
		root = root.children[string(ch)]
	}
	return root
}
