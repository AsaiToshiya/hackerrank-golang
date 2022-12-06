package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
    var minValue, maxValue int32 = arr[0], 0
    var sum int64
    for _, n := range arr {
        minValue = min(minValue, n)
        maxValue = max(maxValue, n)
        sum += int64(n)
    }
    fmt.Printf("%v %v", sum - int64(maxValue), sum - int64(minValue))
}

func min(a, b int32) int32 {
    if a < b {
        return a
    }
    return b
}

func max(a, b int32) int32 {
    if a > b {
        return a
    }
    return b
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
