package main

func fx() {
    tf = (*myType).myMethod
    for i := 0; i < 1000; i++ {
        tf(myObj)
        tf(myObj2)
    }
}

func fy() {
    var ff func()
    for i := 0; i < 1000; i++ {
        ff = myObj.myMethod // myObj.myMethod escapes to heap
        //ff()
        ff = myObj2.myMethod // myObj.myMethod escapes to heap
        //ff()
    }
    _ = ff
}

func fy2() {
    for i := 0; i < 1000; i++ {
        ff := myObj.myMethod
        ff()
        ff = myObj2.myMethod
        ff()
    }
}

func fz() {
    for i := 0; i < 1000; i++ {
        vf = myObj.myMethod // myObj.myMethod escapes to heap
        //vf()
        vf = myObj2.myMethod // myObj.myMethod escapes to heap
        //vf()
    }
}

var (
    myObj = &myType{}
    myObj2 = &myType{}
    vf func()
    tf func(*myType)
)

type myType struct{}

func (m *myType) myMethod() {}

func main() {}
