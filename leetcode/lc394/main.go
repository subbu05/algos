package main

import (
   "fmt"
   "strconv"
)

func main() {
   decodeString("3[a2[cc]]")
}


func decodeString(s string) string {
   temp := []string{}
   res := ""
   st := 0
   for i := 0; i < len(s); i++ {
      if s[i] == '[' || s[i] == ']' {
         temp = append(temp, string(s[st:i]))
         st = i+1
      }
   }
   fmt.Println(temp)

   j:=0

   for len(temp) > 0 {
      str := temp[len(temp)-1]
      temp = temp[0:len(temp)-1]
      for j=0;j<len(str);{
         if str[j] >= 48 && str[j] <= 57 {
            j++
         } else {
            break
         }
      }
   }
   for i:=len(temp)-1;i>=0;i--{
       fmt.Println(temp[i])

      num,_ := strconv.Atoi(string(temp[i][0:j]))
      str := string(temp[i][j:])
      for k:=0;k<num;k++{
         res+=str
      }
   }
   return res
}
