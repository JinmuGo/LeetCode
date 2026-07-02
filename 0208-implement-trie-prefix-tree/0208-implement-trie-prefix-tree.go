type TrieNode struct {
    Children map[rune] *TrieNode
    IsEnd bool
}

func NewTrieNode() *TrieNode {
    return &TrieNode {
        Children: make(map[rune]*TrieNode),
        IsEnd: false,
    }
}

type Trie struct {
    Root *TrieNode
}


func Constructor() Trie {
    return Trie{
        Root: NewTrieNode(),
    }
}


func (this *Trie) Insert(word string)  {
    current := this.Root

    for _, char := range word {
        if _, exists := current.Children[char]; !exists {
            current.Children[char] = NewTrieNode()
        }

        current = current.Children[char]
    }

    current.IsEnd = true
}

func (this *Trie) Search(word string) bool {
    current := this.Root

    for _, char := range word {
        if nextNode, exists := current.Children[char]; exists {
            current = nextNode
        } else {
            return false
        }
    }

    return current.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
    current := this.Root

    for _, char := range prefix {
        if nextNode, exists := current.Children[char]; exists {
            current = nextNode
        } else {
            return false
        }
    }

    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */