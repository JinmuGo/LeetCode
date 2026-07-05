
type Elem struct {
    Score int
    Count int
}

const MOD = 1000000007

func pathsWithMaxScore(board []string) []int {

    n := len(board)
    dp := make([][]Elem, n)

    for i := range dp {
        dp[i] = make([]Elem, n)
    }

    dp[0][0] = Elem{Score: 0, Count: 1}

    for i, row := range board {
        for j, _ := range row {
            if board[i][j] == 'X' || i == 0 && j == 0 {
                continue
            }

            prevCases := [][]int{{i - 1, j}, {i, j - 1}, {i -1, j - 1}}
            // maxArray := make([]int, 0)

            maxScore := -1
            totalCount := 0

            for _, coord := range prevCases {
                py, px := coord[0], coord[1]

                if px >= 0 && px < n && py >= 0 && py < n {
                    if dp[py][px].Score > maxScore {
                        maxScore = dp[py][px].Score
                        totalCount = dp[py][px].Count
                    } else if dp[py][px].Score == maxScore {
                        totalCount = (totalCount + dp[py][px].Count) % MOD
                    }
                }
            }

            if maxScore != -1 {
                curScore := 0

                if board[i][j] != 'S' && board[i][j] != 'E' {
                    curScore = int(board[i][j] - '0')
                }
                dp[i][j] = Elem{ Score: maxScore + curScore, Count: totalCount}
            }
        }
    }

    finalCell := dp[n-1][n-1]
    if finalCell.Count == 0 {
        return []int{0, 0}
    }

    return []int{finalCell.Score, finalCell.Count}
}