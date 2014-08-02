package algorithms

//import "fmt"
type Ord interface {
    Swap(i, j int)  
    Less(i,j int) bool 
    Len() int
} 

type intArray []int
type strArray []string

//methods for intArray
func (p intArray) Swap(i int, j int){
    p[i], p[j] = p[j], p[i]
}

func (p intArray) Len() int{
    return len(p)
}

func (p intArray) Less(i,j int) bool{
    return p[i] < p[j]
}

//methods for strArray
func (p strArray) Swap(i, j int){
    p[i], p[j] = p[j], p[i]
}
func (p strArray) Len() int{
    return len(p)
}
func (p strArray) Less(i, j int) bool{
    return p[i] < p[j]
}

func Patition(A Ord, l int, r int) int {
    var i int = l-1
    for j:=l; j<r; j++ {
        if A.Less(j, r) {
            i = i + 1
            A.Swap(j, i)
        } 
    }
    A.Swap(r, i+1)
    return i+1
}

func QuickSort(A Ord, l, r int) {
    if l < r {
        m := Patition(A, l, r)       
        //fmt.Println(l, r, m)
        QuickSort(A, l, m-1)
        QuickSort(A, m+1, r)
    }
}

