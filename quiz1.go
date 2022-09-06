package main

func main()  {
	var a, b, k, x []int
	var z int	
	
	for k := 1; k < 10; k++ {
		x[k]= (a[k+1]+ b[k+1] )/2
			if a[k-1]*x[k]<0{
				b[k] =x[k]
				a[k]=a [k+1]
			}else{
				a[k]=a [k-1]
				b[k]= b[k+1] 
			}
		|b[k]-a[k]| <= z
	}
}
