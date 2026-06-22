func predictPartyVictory(senate string) string {
    d := make([]int, 0)
    r := make([]int, 0)
    n := len(senate)

    for i, val := range senate {
        switch val {
            case 'D':
                d = append(d, i)
            case 'R':
                r = append(r, i)
        }
    }

    for len(d) != 0 && len(r) != 0 {
        D := d[0]
        R := r[0]

        d = d[1:]
        r = r[1:]

        if D < R {
            d = append(d, D + n)
        } else {
            r = append(r, R + n)
        }
    }

    if len(d) > 0 {
        return "Dire"
    }

    return "Radiant"
}