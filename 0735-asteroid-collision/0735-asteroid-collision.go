type Stack[T any] []T

func (s *Stack[T]) IsEmpty() bool {
    return len(*s) == 0
}

func (s *Stack[T]) Push(data T) {
    *s = append(*s, data)
}

func (s *Stack[T]) Pop() T {
    if s.IsEmpty() {
        var zero T
        return zero
    } 

    topIndex := len(*s) - 1
    data := (*s)[topIndex]

    var zero T
    (*s)[topIndex] = zero
    *s = (*s)[:topIndex]

    return data
}

func (s *Stack[T]) Top() T {
        if s.IsEmpty() {
        var zero T
        return zero
    } 
    topIndex := len(*s) - 1
    data := (*s)[topIndex]

    return data
}

func direction(val int) bool {
    if val > 0 {
        return true
    } else {
        return false
    }
}

func Abs[T int](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func asteroidCollision(asteroids []int) []int {
    s := Stack[int]{}

    for _, elem := range asteroids {
        alive := true

        for !s.IsEmpty() && s.Top() > 0 && elem < 0 {
            top := s.Top()

            if Abs(top) < Abs(elem) {
                s.Pop()
                continue
            } else if Abs(top) == Abs(elem) {
                s.Pop()
                alive = false
                break
            } else {
                alive = false
                break
            }
        }

        if alive {
            s.Push(elem)
        }
    }

    return []int(s)
}