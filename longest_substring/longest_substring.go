// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package sequence

func predecessor(s string, start int, stop int) int {
	for i := start - 1; i >= stop; i-- {
		if s[i] == s[start] {
			return i
		}
	}
	return -1
}

func max(v1, v2 int) int {
	if v1 >= v2 {
		return v1
	}
	return v2
}

func lengthOfLongestSubstring(s string) int {
	count := 1
	result := 1
	i := 1
	left_border_in_string := 0
	if len(s) < 1 {
		return 0
	}
	for i < len(s) {
		same_char_position := predecessor(s, i, left_border_in_string)
		if same_char_position == -1 {
			count++
			i++
		} else {
			result = max(result, count)
			count = 1
			left_border_in_string = same_char_position + 1
			i = left_border_in_string + 1
		}
	}
	return max(result, count)
}

func lengthOfLongestSubstring2(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	seenChar := make(map[rune]int)
	left := 0
	longest := 0
	for right := 0; right < len(s); right++ {
		currentChar := s[right]
		prevSeenChar := seenChar[rune(currentChar)]
		if (prevSeenChar - 1) >= left {
			left = (prevSeenChar - 1) + 1
		}
		seenChar[rune(currentChar)] = right + 1
		longest = max(longest, right-left+1)

	}
	return longest
}
