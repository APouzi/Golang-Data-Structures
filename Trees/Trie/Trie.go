package bst

type Trie struct {
	Head *TrieNode
}

type TrieNode struct {
	Nodes map[byte]*TrieNode
}

func InitializeTrie() {

}