package main

import (
	"fmt"
)

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)
	d := make([]int, N, N)
	for i := range d {
		fmt.Scanf("%d", &(d[i]))
	}
	c := make([]int, M+1, M+1)
	for i := 1; i <= M; i++ {
		fmt.Scanf("%d", &(c[i]))
	}
	dp := make([]map[int]int, M+1, M+1)
	dp[0] = map[int]int{0: 0}
	for i := 1; i <= M; i++ {
		dp[i] = map[int]int{}
		for j := range dp[i-1] {
			v, ok := dp[i][j]
			if ok {
				dp[i][j] = min(v, dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
			if j < N {
				v, ok := dp[i][j+1]
				if ok {
					dp[i][j+1] = min(v, dp[i-1][j]+c[i]*d[j])
				} else {
					dp[i][j+1] = dp[i-1][j] + c[i]*d[j]
				}
			}
		}
	}
	fmt.Println(dp[M][N])

}
