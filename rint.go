package rint

func rint(x int) int {
    rx := 0

    for x > 0 {
        rx = rx * 10 + x % 10
        x /= 10
    }

    return rx
}