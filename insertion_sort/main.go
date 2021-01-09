package main

func main() {
	a := []int{5, 2, 1, 3, 4}
	for i := 0; i < len(a); i++ {
		temp := a[i]
		j := i - 1
		for j = i; j > 0 && a[j-1] > temp; j-- {
			a[j] = a[j-1]
		}
		a[j] = temp
	}
}

/*func main() {
	a := []int{5,4,3,2,1}
	j:=0
	for i:=0;i<len(a);i++{
		temp:=a[i]
		fmt.Println(temp)
		for j=i;j>0 && a[j-1]>temp;j--{
			a[j]=a[j-1]
		}
		fmt.Println(a)
		a[j]=temp
	}
	fmt.Println(a)
}*/
