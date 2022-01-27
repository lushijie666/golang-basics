package main

func main()  {
	a := 0
	switch a {
		case 1:
		case 2:
			fallthrough
		case 3:
		default:
	}
}
