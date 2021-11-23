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

type Id = interface {
	Ia // 内嵌Ia
	Ib // 内嵌Ib
}

type Ie interface {
	Ia // 内嵌Ia
	fb()
}

func main() {

}
