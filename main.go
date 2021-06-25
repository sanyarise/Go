//1. С интерфейсами мы столкнулись когда обрабатывали ошибки при сканировании
//из потока ввода
//2. На первых домашних заданиях код писался без проверки захватываемого потока ввода на
//ошибки. Соответственно, там хочется добавить такую проверку, чтобы исключить
//неправильную работу программы.

package main

import (
	"fmt"
)

type pointerMethodInterface interface {
	pointerMethod()
}
type valueMethodInterface interface {
	valueMethod()
}

type Struct struct {
	s string
}

func (st *Struct) pointerMethod() {
	fmt.Println("Field:", st.s, "Method: pointerMethod")
}

func (st Struct) valueMethod() {
	fmt.Println("Field:", st.s, "Method: valueMethod")
}
func main() {

	pointerStruct := &Struct{
		s: "pointerStruct",
	}
	valueStuct := Struct{
		s: "valueStruct",
	}
	pointerStruct.pointerMethod()
	valueStuct.valueMethod()
	pointerStruct.valueMethod()
	valueStuct.pointerMethod()

	var emptyInterfaceInstancePointerStruct pointerMethodInterface = pointerStruct
	var emptyInterfaceInstanceValueStruct valueMethodInterface = valueStuct

	emptyInterfaceInstancePointerStruct.pointerMethod()
	//emptyInterfaceInstancePointerStruct.valueMethod()
	//emptyInterfaceInstanceValueStruct.pointerMethod()
	emptyInterfaceInstanceValueStruct.valueMethod()
	valueStuct.s = "newValue"
	valueStuct.valueMethod()
	emptyInterfaceInstanceValueStruct.valueMethod()

}
