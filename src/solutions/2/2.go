
/**
 * Project Euler
 * Author: Brendon Crawford
 *
 * Problem: 2
 *
 * Description:
 *   Each new term in the Fibonacci sequence is generated by adding the
 *   previous two terms. By starting with 1 and 2, the first 10 terms will be:
 *
 *     1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
 *
 *   By considering the terms in the Fibonacci sequence whose values do not
 *   exceed four million, find the sum of the even-valued terms.
 */

package main

import "os"
import "fmt"


func main () () {
    const m uint32 = 4000000
    var v1 uint32 = 1
    var v2 uint32 = 2
    var v uint32 = v1
    var a uint32 = 0
    for v <= m {
        v = v1 + v2
        v1 = v2
        v2 = v
        if v1 % 2 == 0 {
            a += v1
        }
    }
    fmt.Printf("%d\n", a)
    os.Exit(0)
}


