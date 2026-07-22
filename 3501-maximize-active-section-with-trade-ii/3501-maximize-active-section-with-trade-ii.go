type SparseTable struct {
	sparse [][]int
}

func NewSparseTable(nums []int) *SparseTable {
	n := len(nums)
	sparse := make([][]int, 21)
	for i := range sparse {
		sparse[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		sparse[0][i] = nums[i]
	}

	for base := 1; base <= 20; base++ {
		pow2 := 1 << (base - 1)
		for i := 0; i < n; i++ {
			if i+pow2 < n {
				sparse[base][i] = max(sparse[base-1][i], sparse[base-1][i+pow2])
			} else {
				sparse[base][i] = sparse[base-1][i]
			}
		}
	}

	return &SparseTable{sparse: sparse}
}

func (st *SparseTable) Query(l, r int) int {
	if l > r {
		return 0
	}
	base := 0
	for base <= 20 {
		if (1 << base) > (r - l + 1) {
			break
		}
		base++
	}
	base--
	if base < 0 {
		return 0
	}
	return max(st.sparse[base][l], st.sparse[base][r-(1<<base)+1])
}

type SegmentTree struct {
	n   int
	arr []int
	seg []int
}

func NewSegmentTree(nums []int) *SegmentTree {
	n := len(nums)
	st := &SegmentTree{
		n:   n,
		arr: nums,
		seg: make([]int, 4*n),
	}
	if n > 0 {
		st.build(1, 0, n-1)
	}
	return st
}

func (st *SegmentTree) build(node, l, r int) {
	if l == r {
		st.seg[node] = st.arr[l]
		return
	}
	mid := (l + r) >> 1
	st.build(2*node, l, mid)
	st.build(2*node+1, mid+1, r)
	st.seg[node] = max(st.seg[2*node], st.seg[2*node+1])
}

func (st *SegmentTree) internalQuery(node, start, end, l, r int) int {
	if l <= start && end <= r {
		return st.seg[node]
	}
	mid := (start + end) >> 1
	res := 0
	if mid >= l {
		res = max(res, st.internalQuery(node*2, start, mid, l, r))
	}
	if r > mid {
		res = max(res, st.internalQuery(node*2+1, mid+1, end, l, r))
	}
	return res
}

func (st *SegmentTree) Query(l, r int) int {
	if l > r || st.n == 0 {
		return 0
	}
	return st.internalQuery(1, 0, st.n-1, l, r)
}

func lowerBound(zeroRight []int, seg, l int) int {
	left, right := 0, seg
	for left < right {
		mid := (left + right) >> 1
		if zeroRight[mid] >= l {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func upperBound(zeroLeft []int, seg, r int) int {
	left, right := 0, seg
	for left < right {
		mid := (left + right) >> 1
		if zeroLeft[mid] <= r {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func maxActiveSectionsAfterTrade(s string, q [][]int) []int {
	n := len(s)
	cnt1 := 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			cnt1++
		}
	}

	var zeroBlocks []int
	var zeroLeft []int
	var zeroRight []int

	idx := 0
	for idx < n {
		r := idx
		for r < n && s[idx] == s[r] {
			r++
		}
		curBlockLen := r - idx
		if s[idx] == '0' {
			zeroBlocks = append(zeroBlocks, curBlockLen)
			zeroLeft = append(zeroLeft, idx)
			zeroRight = append(zeroRight, r-1)
		}
		idx = r
	}

	m := len(zeroBlocks)
	seg := m
	ans := make([]int, 0, len(q))

	// base case
	if m <= 1 {
		for i := 0; i < len(q); i++ {
			ans = append(ans, cnt1)
		}
		return ans
	}

	nums := make([]int, m-1)
	for bl := 0; bl < m-1; bl++ {
		nums[bl] = zeroBlocks[bl] + zeroBlocks[bl+1]
	}

	sp := NewSparseTable(nums)

	for i := 0; i < len(q); i++ {
		l, r := q[i][0], q[i][1]

		lIdx := lowerBound(zeroRight, seg, l)
		rIdx := upperBound(zeroLeft, seg, r) - 1

		if lIdx > m-1 || rIdx < 0 || lIdx >= rIdx {
			ans = append(ans, cnt1)
			continue
		}

		leftLen := zeroRight[lIdx] - max(zeroLeft[lIdx], l) + 1
		rightLen := min(r, zeroRight[rIdx]) - zeroLeft[rIdx] + 1

		if lIdx+1 == rIdx {
			contribution := leftLen + rightLen
			ans = append(ans, cnt1+contribution)
			continue
		}

		leftContri := leftLen + zeroBlocks[lIdx+1]
		rightContri := rightLen + zeroBlocks[rIdx-1]
		middleContri := sp.Query(lIdx+1, rIdx-2)

		ans = append(ans, cnt1+max(leftContri, max(rightContri, middleContri)))
	}

	return ans
}