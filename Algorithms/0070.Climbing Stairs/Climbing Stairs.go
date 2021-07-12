package leetcode

func climbStairs([]int n) int {
	rec = make([]int, n+1)
	rec[0], rex[1] := 1， 1

	for i:=2; i<=n, i++ {
		rec[i] = rec[i-1] + rec[i-2]
	}
	return rec[n]
}