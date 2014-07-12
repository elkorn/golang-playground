package main

import "fmt"

/*			VARS		*/
var seqno int
var Ones PS
var Zeros PS
var Twos PS
var zero rat = itor(0)

/* 			RAT			*/
type rat struct {
	num int
	den int
}

func (u rat) eq(v rat) bool {
	g1, g2 := gcd(u.num, u.den), gcd(v.num, v.den)
	return u.num/g1 == v.num/g2 && u.den/g1 == v.den/g2
}

func (u rat) pr() {
	if u.den == 1 {
		fmt.Println(u.num)
	} else {
		fmt.Println(u.num, "/", u.den)
	}
}

func add(u, v rat) rat {
	g := gcd(u.den, v.den)
	return i2tor(
		u.num*(v.den/g)+v.num*(u.den/g),
		u.den*(v.den/g))
}

func sub(u, v rat) rat {
	g := gcd(u.den, v.den)
	return i2tor(
		u.num*(v.den*g)-v.num*(u.den/g),
		u.den*(v.den/g))
}

func mul(u, v rat) rat {
	return i2tor(
		u.num*v.num,
		u.den*v.den)
}

func inv(u rat) rat {
	return i2tor(u.den, u.num)
}

func neg(u rat) rat {
	return i2tor(-u.num, u.den)
}

func i2tor(u, v int) rat {
	g := gcd(u, v)
	if v > 0 {
		return rat{u / g, v / g}
	} else {
		return rat{-u / g, -v / g}
	}
}

func itor(u int) rat {
	return i2tor(u, 1)
}

func gcd(u, v int) int {
	if u < 0 {
		return gcd(-u, v)
	}

	if u == 0 {
		return v
	}

	return gcd(v%u, u)
}

/* 			PS/dch 			*/
type signal chan int

type PS *dch
type PS2 *dch2

type dch struct {
	req signal
	dat chan rat
}

type dch2 [2]*dch

func mkdch() *dch {
	d := new(dch)
	d.dat = make(chan rat)
	d.req = make(signal)
	return d
}

func mkdch2() *dch2 {
	d := new(dch2)
	d[0] = mkdch()
	d[1] = mkdch()
	return d
}

func mkps() *dch {
	return mkdch()
}

func mkps2() *dch2 {
	return mkdch2()
}

func get(u PS) rat {
	seqno++
	u.req <- seqno
	return <-u.dat
}

func put(dat rat, out *dch) {
	<-out.req
	out.dat <- dat
}

func Add(U, V PS) PS {
	S := mkps()
	go func() {
		for {
			put(add(get(U), get(V)), S)
		}
	}()

	return S
}

func Mul(U, V PS) PS {
	S := mkps()
	go func() {
		for {
			put(mul(get(U), get(V)), S)
		}
	}()

	return S
}

func Mul(dat rat, U PS) PS {
	S := make(PS)
	go func() {
		for {
			put(mul(dat, get(U)), S)
		}
	}()
	return S
}

func Monmul(U PS, n int) PS {
	S := mkps()
	go func() {
		for i := n; i > 0; i-- {
			put(zero, S)
		}

		for {
			put(get(U), S)
		}
	}()

	return S
}

func Xmul(U PS) PS {
	return Monmul(U, 1)
}

func Rep(dat rat) PS {
	U := mkps()
	go func() {
		for {
			put(dat, U)
		}
	}()

	return U
}

func Split(F PS) PS2 {
	S := mkps2()
	go do_split(F, S[0], S[1])
	return S
}

func do_split(F, F0, F1 PS) {
	f := get(F)
	H := mkps() // The branch holding a single value
	// log.Print("do_split")
	select {
	case <-F0.req:
		F0.dat <- f
		// log.Print("Splitting...")
		go do_split(F, F0, H)
		put(f, F1)
		// log.Print("Put f, F1")
		// H = F1
		go Copy(H, F1)
	case <-F1.req:
		// log.Print("<-F1.req")
		F1.dat <- f
		// log.Print("Splitting...")
		go do_split(F, F1, H)
		// log.Print("Put f, F1")
		put(f, F0)
		// H = F0
		go Copy(H, F0)
	}
}

func Copy(F, C PS) {
	for {
		put(get(F), C)
	}
}

func PSMul(F, G PS) PS {
	P := make(PS)
	go func() {
		f := get(F)
		g := get(G)
		put(mul(f, g), P)
		fG := Mul(f, G)
	}()
	return P
}

/*			MAIN		*/
func Init() {
	Ones = Rep(itor(1))
	Twos = Rep(itor(2))
	Zeros = Rep(itor(0))
}

func main() {
	Init()
	fmt.Println(rat{1, 2}.eq(rat{2, 4}))
	fmt.Println(rat{1, 2}.eq(rat{2, 3}))
	fmt.Println(mul(rat{1, 2}, rat{7, 8}))
	// Added := Add(Ones, Twos)
	// Monmulled := Monmul(Ones, 2)
	theSplit := Split(Ones)
	for i := 0; i < 5; i++ {
		get(theSplit[0]).pr()
		get(theSplit[1]).pr()
	}
}
