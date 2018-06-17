package main

type myType struct{}

func (m *myType) myMethod() {for false {}}

var (
    myObj = &myType{}
    myObj2 = &myType{}
)

func fx() {
    var tf = (*myType).myMethod
    for i := 0; i < 1000; i++ {
        tf(myObj)
        tf(myObj2)
    }
}

func fy() {
    var ff func()
    for i := 0; i < 1000; i++ {
        ff = myObj.myMethod // myObj.myMethod escapes to heap
        ff()
        ff = myObj2.myMethod // myObj.myMethod escapes to heap
        ff()
    }
}

func fz() {
    for i := 0; i < 1000; i++ {
        ff := myObj.myMethod
        ff()
        ff = myObj2.myMethod
        ff()
    }
}

func fw() {
    for i := 0; i < 1000; i++ {
        myObj.myMethod()
        myObj2.myMethod()
    }
}

func main() {}
