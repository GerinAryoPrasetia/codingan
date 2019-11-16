package main
import "fmt"

const N = 2019
var batas int
type RecType struct {
    f1 string
    f2 int
    f3 float64
}
type ArrType [N]RecType

func main() {
    var arr ArrType
    var find string
    i:= 0
    c := true

 for c {
    fmt.Scan(&arr[i].f1, &arr[i].f2, &arr[i].f3)
        if arr[i].f1 == "0" && arr[i].f2 == 0 && arr[i].f3 == 0 {
            c = false
        } else {
        batas++
        }
        i++
    }
 fmt.Println("rmax = ", arr[rmax(arr)].f3)
 fmt.Println("imin = ", imin(arr))
 fmt.Print("masukkan kata yang ingin dicari = ")
 fmt.Scan(&find)
 fmt.Println("found = ", found(arr, find))
 fmt.Println("pos = ", pos(arr, find))
 fmt.Println("Nama : Gerin Aryo Prasetia\nNIM : 1301194479")
}

func rmax(data ArrType) int {
    var max = 0
    i := 0
        for i < batas {
            if data[max].f3 < data[i].f3 {
            max = i
            }
            i++
        }
        return max
}

func imin(data ArrType) int {
    var min = 0
    i := 0
        for i < batas {
            if data[min].f2 > data[i].f2 {
                min = i
            }
        i++
}
    return min
}

func found(data ArrType, key string) bool {
    var found bool
    found = false
    i := 0
    for i < batas {
        if key == data[i].f1 {
        found = true
        }
    i++
 }
    return found
}

func pos(data ArrType, key string) int {
    var o, top, btm, mid int
    btm = 0
    top = batas - 1
    mid = (top + btm) / 2
    for top >= btm && key != data[mid].f1 {
        if key > data[mid].f1 {
            btm = mid + 1
        } else {
            top = mid - 1
  }
    mid = (top + btm) / 2
}
    if key == data[mid].f1 {
        o = mid
    } else {
        o = -1
    }
return o
}