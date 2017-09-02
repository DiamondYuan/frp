package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
)

var confMap map[string](map[string]string)

func main() {
	prefix := "FRP_"
	filename := "frp.ini"
	f, err := os.Create("/frp/"+filename)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
	}
	defer f.Close()
	confMap = make(map[string](map[string]string))
	w := bufio.NewWriter(f)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		rayKey := pair[0];
		if len(rayKey) > len(prefix) && strings.Compare(rayKey[0:len(prefix)], prefix) == 0 {
			temp := strings.ToLower(rayKey[len(prefix):])
			temps :=strings.Split(temp,"_")
			head := temps[0]
			key := temp[len(head)+1:]
			tempMap := make(map[string]string)
			if confMap[head] != nil {
				tempMap = confMap[head]
			}
			tempMap[key]=pair[1]
			confMap[head]=tempMap
		}
	}
	for k, v := range confMap {
		fmt.Fprintln(w,"["+k+"]")
		for k1,v1 :=range v{
			fmt.Fprintln(w,k1 +" = "+v1)
		}
		fmt.Fprintln(w)
	}
	w.Flush()
}
