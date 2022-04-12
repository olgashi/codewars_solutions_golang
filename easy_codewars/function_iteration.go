package main

func CreateIterator(fn func(int) int, n int) func(int) int {
	return func(someNum int) int {
    res := someNum
    for index := 0; index < n; index++ {
			res = fn(res)
		}
		return res
	}
}

func main(){
// see kata Function iteration
}

