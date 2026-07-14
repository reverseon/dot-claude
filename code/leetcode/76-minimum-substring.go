package main

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	// Step 1: Build a frequency map of what we need from t
	// This lets us quickly check: "do we have enough of character X?"
	// Example: t="ABC" → required={A:1, B:1, C:1}
	required := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		required[t[i]]++
	}

	// Step 2: Track what we currently have in our sliding window
	// We'll compare this against `required` to know if we're valid
	window := make(map[byte]int)

	// Step 3: Track how many UNIQUE characters are "satisfied"
	// A char is satisfied when we have AT LEAST the required count
	// Example: if we need A:1, and window has A:2, that char is satisfied
	// Once all chars are satisfied (formed == len(required)), window is valid
	formed := 0
	required_len := len(required)

	// Step 4: Bookkeeping for the answer
	// We'll slide through s and record the minimum window we find
	min_len := len(s) + 1
	min_start := 0

	left := 0

	// Step 5: Expand-and-Shrink Loop
	// The key insight: scan right to find valid windows, then shrink from left
	// to find the SMALLEST valid window, then expand right again for the next one
	for right := 0; right < len(s); right++ {
		// EXPAND: Pull the right character into the window
		// This grows our window, hopefully getting us closer to having all required chars
		char := s[right]
		window[char]++

		// Did adding this character satisfy a requirement?
		// Check: is this char in t? AND do we now have enough of it?
		if required[char] > 0 && window[char] == required[char] {
			formed++
		}

		// SHRINK: Once we have a valid window, try to make it smaller
		// Why? Because we want the MINIMUM window. So once valid, remove noise from the left
		// Keep shrinking as long as we're still valid (formed == required_len)
		for formed == required_len {
			// Found a valid window! Is it smaller than our best so far?
			// If yes, record it
			if right-left+1 < min_len {
				min_len = right - left + 1
				min_start = left
			}

			// Remove the leftmost character
			// Why remove from left? Because we want to shrink the window to find the minimum
			// As we shrink, we might lose a required character (formed decreases)
			// When we lose a required char, the window becomes invalid, and we stop shrinking
			char := s[left]
			if required[char] > 0 && window[char] == required[char] {
				formed--
			}
			window[char]--
			left++
		}
		// When inner loop exits, we couldn't shrink anymore without losing validity
		// So we loop back to expand right and look for the next valid window
	}

	// If we never found a valid window, min_len stays len(s)+1
	if min_len == len(s)+1 {
		return ""
	}

	return s[min_start : min_start+min_len]
}
