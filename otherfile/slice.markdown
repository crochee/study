# 关于指针切片
源码位于runtime/slice.go
## 指针切片的常见bug
```go
func main() {
	m := parse()
	for k, stu := range m {
		fmt.Println(k, stu)
	}
}

type student struct {
	Name string
	Age  int
}

func parse() map[string]*student {
	var m = make(map[string]*student)
	stus := []student{
		{Name: "1", Age: 19},
		{Name: "2", Age: 20},
		{Name: "3", Age: 21},
	}
	for _, stu := range stus { // 在index=0时，创建了一个stu变量，存的是值
		m[stu.Name] = &stu // 存的时候如果，直接取地址的话，会取到stu变量的地址，
		// 只是它的值经过覆盖之后变成最后一个，所以最后输出的时候，因为指针指向都是stu的指针
		fmt.Println(stu.Name, stu, &stu, m, &stus)
	}
	return m
}
```
上面的代码会得到以下结果：
```
1 &{3 21}
2 &{3 21}
3 &{3 21}
```
那么这是为什么呢？在第一次遍历stus的时候创建了一个临时变量stu，stu存的是student的值，当存入map时，取的是stu的指针，尽管stu的值变化了，但是指针未变化，所以map中存的全是stu的指针，经过几轮赋值之后stu指向的是最后一个值，故当打印的时候，取该地址全是stu的指针地址，只能是值经过几轮赋值后变为最后一个；

正确的写法：
```go
for i := 0; i < len(stus); i++ {
    m[stus[i].Name] = &stus[i]
}
```
这样就能实时的取切片中每个元素的值，并取地址，不需要通过中间变量来赋值
```
2 &{2 20}
3 &{3 21}
1 &{1 19}
```
## 切片的原理

关于详细的原理参考：https://www.jianshu.com/p/030aba2bff41

关于切片的清空的两个个写法
* s[0:0] 相当于s[:0],也相当于s[:0:cap(s)] 只是将长度置零，底层数组没有改变
* s[0:0:0] 相当于s[:0:0] 将容量变为0 相当于底层数组需要重新申请


