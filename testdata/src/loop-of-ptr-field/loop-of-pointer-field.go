package main

func main() {
	var intSlice []*int

	println("loop expecting 10, 11, 12, 13")
	type valueWithin struct {
		i int
	}
	for _, p := range []*valueWithin{{i: 10}, {i: 11}, {i: 12}, {i: 13}} {
		intSlice = append(intSlice, &p.i)
	}

	println(`slice expecting "10, 11, 12, 13" is "10, 11, 12, 13"`)
	for _, p := range intSlice {
		printp(p)
	}

}

func printp(p *int) {
	println(*p)
}
