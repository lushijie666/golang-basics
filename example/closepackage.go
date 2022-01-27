package main

func main()  {
	nextInt := initSeq()
	println(nextInt())
}

func initSeq() func() int{
	i := 0
	return func() int {
		i++
		return i
	}
}

