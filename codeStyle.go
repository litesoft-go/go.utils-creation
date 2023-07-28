package go_utils_creation

// LinesToString -- is designed to make adding linefeed seperated text
// in code, where every desired line is in alignment from an indentation
// perspective.  To accomplish this "line alignment", the first and last line
// are treated special; in that if there is 3 or more lines, then the
// first and/or last lines are removed IF they are empty.
// See tests for examples.
// -
func LinesToString(expectedLines ...string) (expected string) {
	lines := len(expectedLines)
	firstAt, lastAt := 0, lines-1
	switch lines {
	case 0:
		return
	case 1:
		return expectedLines[0]
	case 2:
		break
	default: // 3 or more
		if expectedLines[firstAt] == "" {
			firstAt++
		}
		if expectedLines[lastAt] == "" {
			lastAt--
		}
		break
	}
	expected = expectedLines[firstAt]
	for nextAt := firstAt + 1; nextAt <= lastAt; nextAt++ {
		expected += "\n" + expectedLines[nextAt]
	}
	return
}
