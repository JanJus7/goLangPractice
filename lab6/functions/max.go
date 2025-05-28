package functions

func MaxI(a, b, c int) int {
    m := a
    if b > m {
        m = b
    }
    if c > m {
        m = c
    }
    return m
}

func MaxF(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}