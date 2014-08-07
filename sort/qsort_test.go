package sort

import (
    "testing"
    "math/rand"
    //"fmt"
)

func TestPatition(t *testing.T) {
    Max := 1000000;
    var x = make(intArray, Max)

    for i := 0; i < Max; i++ {
        x[i] = rand.Intn(Max)
    }

    m := Patition(x, 0, Max-1)    

    for i:=0; i<m; i++ {
        if x[i] > x[m] {
            t.FailNow()
       }
    }
    for i:=m; i<Max; i++ {
        if x[i] < x[m] {
            t.FailNow()
        }
    }
}

func TestIntArray(t *testing.T){
    Max := 1000000;
    var x = make(intArray, Max)

    for i := 0; i < Max; i++ {
        x[i] = rand.Intn(Max)
    }

    QuickSort(x, 0, Max-1)

    for i:=0; i<(Max-1); i++ {
        if x[i+1] < x[i] {
            t.Log("QuickSort Error") 
            t.FailNow() 
        }
    }
}

func TestStrArray(t *testing.T){
    x := strArray {"cd","de","ab","bc"}
    QuickSort(x,0, x.Len()-1)
    t.Log(x)
}
