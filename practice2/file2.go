package main

import (
    "fmt"
    "golang.org/x/sys/windows"
)

func main() {
    var free, total, avail uint64

    path:= "c:\\"
    pathPtr, err := windows.UTF16PtrFromString(path)
    if err != nil{
        panic(err)
    }
    err = windows.GetDiskFreeSpaceEx(pathPtr, &free, &total, &avail)


    fmt.Println("Free:", free, "Total:", total, "Available:", avail)
}
