type TrieNode struct  {
    Child map[rune] *TrieNode
    IsEnd bool
    Suggestions []string
}

func NewTrieNode () *TrieNode {
    return &TrieNode {
        Child: make(map[rune] *TrieNode),
        IsEnd: false,
        Suggestions: make([]string, 0, 3),
    }
}

type Trie struct {
    Root *TrieNode
}

func Constructor () Trie {
    return Trie{
        Root: NewTrieNode(),
    }
}

func (this *Trie) Insert(word string) {
    cur := this.Root

    for _, char := range word {
        if _, exists := cur.Child[char]; !exists {
            cur.Child[char] = NewTrieNode()
        }

        cur = cur.Child[char]

        if len(cur.Suggestions) < 3 {
            cur.Suggestions = append(cur.Suggestions, word)
        }
    }

    cur.IsEnd = true
}

func (this *Trie) Search(word string) bool {
    cur := this.Root

    for _, char := range word {
        if child, exists := cur.Child[char]; exists {
            cur = child
        } else {
            return false
        }
    }

    return cur.IsEnd
}

func (this *Trie) StartWith(prefix string) bool {
    cur := this.Root

    for _, char := range prefix {
        if child, exists := cur.Child[char]; exists {
            cur = child
        } else {
            return false
        }
    }

    return true
}

func (this *Trie) GetSuggestions(searchWord string) [][]string {
    res := make([][]string, len(searchWord))
    cur := this.Root

    isBroken := false

    for i, char := range searchWord {
        if isBroken || cur.Child[char] == nil {
            isBroken = true
            res[i] = []string{}
            continue
        }

        cur = cur.Child[char]
        res[i] = cur.Suggestions
    }

    return res
}


func suggestedProducts(products []string, searchWord string) [][]string {
    trie := Constructor()
    sort.Strings(products)

    for _, product := range products {
        trie.Insert(product)
    }

    return trie.GetSuggestions(searchWord)
}