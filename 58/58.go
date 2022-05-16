package main

func lengthOfLastWord(s string) int {
    start := len(s) - 1
    end := -1

    for i := len(s) - 1; i >= 0; i-- {
        if string(s[i]) != " " {
            start = i
            break
        }
    }

    for i := start; i >= 0; i-- {
        if string(s[i]) == " " {
            end = i
            break
        }
    }

    return start - end
}