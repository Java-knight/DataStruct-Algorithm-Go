package _044_wildcard_matching

// 通配符匹配
// https://leetcode.cn/problems/wildcard-matching/

// s由26个小写字母组成, p由小写字母和通配符组成(?, *)

// 动态规划(dp二维表)(斜率优化)
func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)
	// 递归终止条件也是index走到arr.length位置, 所以需要创建dp数组是size+1
	dp := make([][]bool, sLen+1)
	for i := range dp {
		dp[i] = make([]bool, pLen+1)
	}

	// base case1 s匹配完了
	dp[sLen][pLen] = true               // (1) s和p都匹配完了
	for pi := pLen - 1; pi >= 0; pi-- { // (2) pi位置的值依赖于自己(自己属于*)和pi+1位置的值([pi+1...]能否完全匹配)
		dp[sLen][pi] = p[pi] == '*' && dp[sLen][pi+1]
	}
	// base case2 p已经匹配完了(多余的, 为了方便理解)

	// base case3 s从si出发... p从pi出发...
	// dp[si][pi]依赖dp[si+1][pi+1]和dp[si+len][pi+1]的值
	// 从左到右, 内层从下到上
	for si := sLen - 1; si >= 0; si-- {
		for pi := pLen - 1; pi >= 0; pi-- {
			// (1) 小写字母
			if p[pi] != '?' && p[pi] != '*' {
				dp[si][pi] = s[si] == p[pi] && dp[si+1][pi+1]
				continue
			}
			// (2) ?
			if p[pi] == '?' {
				dp[si][pi] = dp[si+1][pi+1]
				continue
			}
			// (3) '*'
			// 因为dp[si][pi] = dp[si][pi + 1] || dp[si+1][pi + 1] || dp[si+2][pi + 1]...
			// 因为dp[si+1][pi] =  dp[si+1][pi + 1] || dp[si+2][pi + 1]...
			// 所以 dp[si][pi] = dp[si][pi + 1] || dp[si+1][pi]
			dp[si][pi] = dp[si][pi+1] || dp[si+1][pi]
		}
	}
	return dp[0][0]
}

//// 方法2: 动态规划(dp二维表)
//func isMatch(s string, p string) bool {
//	sLen := len(s)
//	pLen := len(p)
//	// 递归终止条件也是index走到arr.length位置, 所以需要创建dp数组是size+1
//	dp := make([][]bool, sLen+1)
//	for i := range dp {
//		dp[i] = make([]bool, pLen+1)
//	}
//
//	// base case1 s匹配完了
//	dp[sLen][pLen] = true               // (1) s和p都匹配完了
//	for pi := pLen - 1; pi >= 0; pi-- { // (2) pi位置的值依赖于自己(自己属于*)和pi+1位置的值([pi+1...]能否完全匹配)
//		dp[sLen][pi] = p[pi] == '*' && dp[sLen][pi+1]
//	}
//	// base case2 p已经匹配完了(多余的, 为了方便理解)
//
//	// base case3 s从si出发... p从pi出发...
//	// dp[si][pi]依赖dp[si+1][pi+1]和dp[si+len][pi+1]的值
//	// 从左到右, 内层从下到上
//	for si := sLen - 1; si >= 0; si-- {
//		for pi := pLen - 1; pi >= 0; pi-- {
//			// (1) 小写字母
//			if p[pi] != '?' && p[pi] != '*' {
//				dp[si][pi] = s[si] == p[pi] && dp[si+1][pi+1]
//				continue
//			}
//			// (2) ?
//			if p[pi] == '?' {
//				dp[si][pi] = dp[si+1][pi+1]
//				continue
//			}
//			// (3) '*'
//			for size := 0; size <= len(s)-si; size++ {
//				if dp[si+size][pi+1] {
//					dp[si][pi] = true
//					break
//				}
//			}
//		}
//	}
//	return dp[0][0]
//}

//// 方法1: 暴力递归
//func isMatch(s string, p string) bool {
//	return process(s, p, 0, 0)
//}
//
//// s[si...] 能否被 p[pi...]匹配出来
//func process(s, p string, si, pi int) bool {
//	// base case1 s已经匹配完了
//	if si == len(s) {  // si ——> ""
//		if pi == len(p) {  // pi ——> ""
//			return true
//		} else {  // pi ——> "..."
//			// p[pi] == '*' && p[pi+1....] ——> 也是空串(从pi+1, 后面也要是空串)
//			return p[pi] == '*' && process(s, p, si, pi + 1)
//		}
//	}
//	// base case2 p已经匹配完了(多余的, 为了方便理解)
//	if pi == len(p) {
//		return si == len(s)   // 这个就是false, 因为如果是true的话, 就走base case1分支了
//	}
//
//	// base case3 s从si出发... p从pi出发...
//	// s[si] ——> 小写字母
//	// p[pi] ——> 小写字母\?\*
//	if p[pi] != '?' && p[pi] != '*' {  // (1) 小写字母
//		return p[pi] == s[si] && process(s, p, si + 1, pi + 1)
//	}
//	if p[pi] == '?' {  // (2) ?
//		return process(s, p, si, pi + 1)
//	}
//	// p[pi] == '*'  len=0,p[pi]一个也没有匹配到字符串, 需要让p[pi+1...]去匹配s[si...]
//	// len=1,p[pi]匹配到一个字符串, 需要让p[pi+1...]去匹配s[si+1...]
//	// len=2,p[pi]匹配到两个字符串, 需要让p[pi+1...]去匹配s[si+2...]
//	for size := 0; size <= len(s) - si; size++ {  // (3) *
//		if process(s, p, si + size, pi + 1) {
//			return true
//		}
//	}
//	return false
//}
