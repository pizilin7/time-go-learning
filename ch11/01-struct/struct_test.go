package struct_test

import (
    "fmt"
    "testing"
)

type Employee struct {
    Id string
    Name string
    Age int
}

// 打印Employee信息
// 传入的是指针
func (e *Employee) String1() string {
    return fmt.Sprintf("ID:%s/Name:%s/Age:%d, Employee:%p\n", e.Id, e.Name, e.Age, &e)
}

// 传入的是struct,会拷贝一份
func (e Employee) String2() string {
    return fmt.Sprintf("ID:%s/Name:%s/Age:%d, Employee:%p\n", e.Id, e.Name, e.Age, &e)
}

func TestEmployee(t *testing.T) {
    e := Employee{"0001", "Bob", 18}
    e1 := Employee{Name:"Bob", Age : 18, Id : "0002"}

    t.Log(e)
    t.Log(e1)

    // 创建的对象是指针
    e2 := new(Employee)
    e2.Id = "003"
    e2.Name = "Rose"
    e2.Age = 16
    t.Log(e2)

    t.Logf("e is %T", e)
    t.Logf("e2 is %T", e2)
}

func TestCreateEmployeeObj(t *testing.T) {
    e := Employee{"0", "Bob", 18}
    e1 := Employee{Name:"Bob", Age : 18, Id : "0002"}
    e2 := new(Employee)
    e2.Id = "003"
    e2.Name = "Rose"
    e2.Age = 16

    t.Logf("e = %p\n", &e)
    t.Logf("%s", e.String1())
    t.Logf("%s", e.String2())


    t.Logf("e1 = %p\n", &e1)
    t.Logf("%s", e1.String1())
    t.Logf("%s", e1.String2())


    t.Logf("e2 = %p\n", &e2)
    t.Logf("%s", e2.String1())
    t.Logf("%s", e2.String2())
}