package blah


type Person struct {
	Age int // public
	Name string // public
	aadhar string // private
}

/*
local -- cannot touch
global -- possibly touch
      |-- Capital, Smol

func Gu() {
	var ; // local
}

// import "blah"

// var keo blah.Person
// keo.Name, keo.Age, keo.aadhar
// keo.Gu(), keo.Gu.k