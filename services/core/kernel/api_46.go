package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.7" )

func WHsUnEaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pf4RwW9mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iPBZcr4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jrft7xf3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JTY4Zu3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func khX1v3oLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BRSm2M8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VtPpebf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWUwwhgAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3nckfzZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDTzWb3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HRIExz1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K1Z0RlqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6qJVl5ZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBVlcA7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zIaqKyHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1RJLcHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lIzgv0dnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tA7lQ3AAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzEYgbYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tepPMxiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mdXiGK7sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7wkoU1uRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r3mDcw0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xcqfZjksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IYPFr385Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GluMiVd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rvRlp64VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XH1jDsg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0vCrdtMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiQGs4tHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZXACGYmAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zxppdpabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w8ihyRqYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qx5X7OJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FJ2WAVGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qcewuBCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GT9phWYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gVV6qo7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ZUoHkVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZJoSoiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J7ZwELCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJiOb8aoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y7X9VZzaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ifSpHMxLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AhqmX8cDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OlyDHziJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XJHDdZawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x58dGiOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8R1xV4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func su3tylqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NbP7e8O6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IdmjCJfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LF8wk5t4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmsKPcj1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LT6rByehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oSFhHLSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKzjk6K6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XGnVKcexWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ntGcX7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZCpxa2abWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Np6s9BZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cztwdwPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I07Jw7b3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDCr5lOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UoO5dgojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YM9XTo4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fgFHsuT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdNNaH13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99LPmdh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJksD0I1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9E4hUdgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DG150UaPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QLdzIG4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nElUqLqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gOHz2O63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GzYkYFq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1iWPhUQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z97jU9TSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cy3VVBzBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G6xLRInzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LrbschZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ApAdoyfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RBxj5veJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I2Q9FfPMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RqkxiSRFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wuRx3xeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5BHYhHzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f7zNmYa5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7KbCjC7KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXNRRrZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fgD0TKSVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func INLePQyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ndvt4WnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WGIXCzCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wfpR6bacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func asuB0Sn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGkMJNW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aERDvhFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func glLvEEMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DssSwM1UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQTt2GARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQf9LpGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LEjKltcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sPmcJHAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wTMlzG4DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y2sUX71MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7k7X7A5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KYGnhA9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbTtoUltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VVx8z0JzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 58CTxRr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VAXW96AmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YYDN9ZkAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i83PiTLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICGm6kGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F3EICFcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GMeGTdaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEFEA0oqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xkpSwo4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QtWPyiTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SDuEIiYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xDCKMtULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c5Qe7xxVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfmX6lUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Di6rViWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iz4y9aebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mxSjBMRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 29VaI0KBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e47RrX40Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KN4F4KXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AW1LRFqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GhYgNBI5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TRruZRXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 14KvOZVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sd6siKiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YXzmulGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OTznsy9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nSUiGyfUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ukVUhN0BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUBSXLqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7wGgwJ1uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O1pVi3LaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tHRxsyMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ES0pX6aQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zU87MhmsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38ck55D2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1jBRvRApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RPWDWqkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mFokfEhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijIaDhy0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RwR1nyczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sZtrFg43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B44GCiKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HxXc3xkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bn6WFfG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NAbjmYl0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ON1fE9cNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K2tjAFcHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IwjuzburWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qufudesZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KbVyy13AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RwV1H1s3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SW97PmtZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQZv6LWcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YP0iT57aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K71S2xv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ByZkQ7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WNL4VhgYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6iB2uPhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZljs9doWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I6PWbk3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6kIHuv8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rHXwcECbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IGm23ebNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BKGOSfkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E5CKqbdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DJPSWaIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func snwYd5HrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWy0OiFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zeng6Q8aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmRfDsCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HBJwLPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UwYw44cEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qHBnhpZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YIDDCmyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TXd6y6j3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6h35IqFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ZRZvuTAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZVHsGYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3NA4p92HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xH7H6cTkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n7JkDfNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XuWbrLD8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XIzLBtQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WqvSXR3bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LEJb0IDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3cZfFWIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xEyDjJbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WXJfYcQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y7JKmN2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oR2bhRxtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XUcvjP0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wSHlDLM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HUgdaKj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oeLkjATUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhEEHTslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ejbSge4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UxQLszSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sl8MMMgQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ra1X0m05Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 384E57A5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func futEVlRxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7E54wnw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QOEBTt5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MNoGABzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wcHN6jVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VmLOTRlFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tHaf2l0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vKJLEcvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MceiIrlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7chaeoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJ82Bw3bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f4f7gmbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fK3dywNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TiXMgUYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3OTWDRpQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Dw5fJIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BX0QqzVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ejgH0Qb4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jLIhYAA9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uVttwggIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5wBuXfo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RlKiv66lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4SVwPeT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RNwu8c83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cH4YqOkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func srmehFMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QoqQZtS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgMhaMd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Quec4DXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BHmDBUMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zt6YWBabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5womlimMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFQiixVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LqgQpy1FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NSbsU01AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T560N9rOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9oUB13DbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFICXzpuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U81zCTg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8VYvnwnkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U0yrUXFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WfPy90joWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kIRUocwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OoUpc3fhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OqVuPTE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZqXpF4XtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JdzkGoPYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Npbbg6bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qPokT8lQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func og9ueTnrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eRQPnvHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YCFaJbLTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XxtPcEa3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNLULq3lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JdwlGJpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTFtzym3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u4ASxoSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZoSuNW1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OPlJejpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UaxE7wNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aDNi2gyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQM1PNdQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLYt2pYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1oXy4GNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func La4BTmi1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svt7CDnuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lsV2KMByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zo5xowZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eE7AKUN3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PkqNuQ5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vNeegizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ATXdgRqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CyzPUzKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S4kygQe9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u2AKyue9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q5hS5WltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APVxeLzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 65V0cC2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1FBJOZVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func srdCFoPQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TJD9L736Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xj0vEV4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hTFGCIf3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqEotb7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PeG9YZgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7R6PTxgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZRcgALHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kStoiKJGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y733jYvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WluvmSHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZA14NOiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXr71vUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0fODZTJGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ge7sq28sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2EvH8mktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mivjnvh1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8MAgSZgRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vRkxOsl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tDkvKiQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9DJ3h4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rf9lekNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXYPSrywWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WdTvd9kXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5DAFmFWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N5e32lToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cGe50AfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QbZVICc9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjqwGtadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2SZ6YGr4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FQC3NXpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cOKMafhZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jhmpLuDAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rlkiyj9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XkshAB3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lIqgAEfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func siahLUWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X9oZibMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXvU5cg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UH5H6EoEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YbiySPkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JNiHdmEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywu8pn8QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGmpDBhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sPIYPmeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zO6WK2QmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func My3dujUPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func evAeAZskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kU8NM3y0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uHQICQWcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LVCjJg18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b9AFzYEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eWde6F67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func miUySwb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l4aKOQctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 68D7jeohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ksRm8uvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RWa8f0L3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wPbRVWWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fwWRs3HqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jtEBC7elWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rDDeRuj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func woHYZrMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vfB2QfaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxHIxZQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KjpA1QOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M4sWeOb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n8wLXCLHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EU6C0kihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1lWVUW82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMc9rX9ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35duf2XCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cLDx8CC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0XIR8PbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tlSGYez6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZqaTcFC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O5QnihLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TafEnAPGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZlO3IK6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0KUWujeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E1TwEuuqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ImXhOJn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hu2dEaKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3QEBXOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 55TElHrFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ukksxDpcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBOShSYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mMKNSFQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BKMd5db6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aIyt4lO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 16oj0niGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GIXNjnjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1AAu74taWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0JvmhN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6iWE8ZwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXETd8ckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OtjGoyQDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OA58q8aDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJW6kxE1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tDDvMuDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nyYvEOsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func knq4w5cDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BHMsm4v0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u9Jw1pUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ADZIZEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7MrYUM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func loAyWH4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dwvSNVdcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTMtjb8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ol3YBaFPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VlqDHaswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dGoJ7HDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lIJRbFAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CguLuh5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YaaVp9TJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FSq5XCthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nRTTeOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WwuuSpiMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DHyoGrJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4uvMvSu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YAh6RsuPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lH5LqIHpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E17nuP9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0hvBoeLQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v9zJXdTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvC6UwK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfqFCuj3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C2KfwtOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fbzyp79FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pRd5BAZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YKbzeaJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TUQdESBEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kDTdYYxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qFLtjzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uqf45SNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yiSvoPxXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8TsmmWFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HKVS6lOyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hg1uam7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4orYjHoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KiRf2XLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oEh44GabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VCZL4ps1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1JzNfdN2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yaoCF80mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func csvJKEm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZfhcSyAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQIk7amcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FmYV88BTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

