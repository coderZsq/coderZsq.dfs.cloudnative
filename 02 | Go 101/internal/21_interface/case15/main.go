package main

type Ia interface {
	fa()
}

type Ib = interface {
	fb()
}

type Ic interface {
	fa()
	fb()
}

type Ix interface {
	Ia
	Ic
}

type Iy = interface {
	Ib
	Ic
}

type Iz interface {
	Ic
	fa()
}

func main() {

}