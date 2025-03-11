// package main

// import "fmt"

// type KBTU interface {
// 	Study() bool
// 	Retake() bool
// }

// type Student struct {
// 	Name string
// 	ID   int
// 	Scholarship bool
// }

// func (s Student) Study() bool {
// 	switch s.Scholarship {
// 	case true:
// 		fmt.Println("Student is studying")
// 		return true
// 	default:
// 		fmt.Println("Student is not studying")
// 	}
// 	return false
// }

// type Teacher struct {
// 	Name string
// 	ID   int
// 	Salary int
// }

// func (t Teacher) Study() bool {
// 	fmt.Println("Teacher is studying")
// 	return true
// }

// func (t Teacher) Retake() bool {
// 	fmt.Println("Teacher is retaking")
// 	return true
// }


// func main(){
// 	student := Student{
// 		Name: "John",
// 		ID: 1,
// 		Scholarship: true,
// 	}

// 	teacher := Teacher{
// 		Name: "Jane",
// 		ID: 2,
// 		Salary: 1000,
// 	}
// 	student.Study()
// 	teacher.Study()
// 	teacher.Retake()

// 	var KBTU = []KBTU{student, teacher}

// 	for _, kbtu := range KBTU {
// 		kbtu.Study()
// 	}
// }

package main

// Разделяем интерфейсы
type Printer interface {
    Print()
}

type Scanner interface {
    Scan()
}

type Fax interface {
    Fax()
}

// Теперь структуры реализуют только нужные интерфейсы
type MultiFunctionPrinter struct{}

func (m *MultiFunctionPrinter) Print() { /* Реализация */ }
func (m *MultiFunctionPrinter) Scan()  { /* Реализация */ }
func (m *MultiFunctionPrinter) Fax()   { /* Реализация */ }

type SimplePrinter struct{}

func (s *SimplePrinter) Print() { /* Реализация */ } // Теперь нет ненужных методов
// func (s *SimplePrinter) Scan()  { /* Реализация */ } // Нет метода Scan
// func (s *SimplePrinter) Fax()   { /* Реализация */ } // Нет метода Fax

