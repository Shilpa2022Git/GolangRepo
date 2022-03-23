package main

import (
	"fmt"
	"time"
	"sync"
)

//instead of sleep, wait should be handled using sync.
	//var wg  sync.Waitgroup
	//wg.add(1) ... signal done using defer to ensure done at end of function wg.done
	//wg.wait till all done

var wg sync.WaitGroup

func goSample(str string){	
	defer wg.Done()	
	for i:=0;i<3;i++ {
		fmt.Println(str, i)
		time.Sleep(10)
	}
}

func main(){
	//goSample("details")

	//call go routine 
	wg.Add(3)	
	go goSample("goroutine---1")

	//call go routine using anonymous function call	
	//wg.Add(1)	
	go func(){
		goSample("goroutine---2")
	}()

	fv := goSample
	go fv("goroutine---3")

	//time.Sleep(3000*1000)
	wg.Wait()
	fmt.Println("done....")

	incr := func(wg1 *sync.WaitGroup){
		var i int

		fmt.Println("incr called")
		wg1.Add(1)

		go func(){
			defer wg1.Done()
			fmt.Println("Called func ", i)
			i++
			fmt.Println("value of i ", i)  //here after returning from outer function, goroutine still have access to i
		}()
		fmt.Println("Return from incr function")
		return
	}

	var wgSync sync.WaitGroup
	incr(&wgSync)

	//Passing specific value to go routine.

	printNums := func(wg1 *sync.WaitGroup){
		for j:=0; j<3;j++{
		wg1.Add(1)

		go func(j int){
			defer wg1.Done()						
			fmt.Println("value of j passed ", j)  //here after returning from outer function, goroutine still have access to i
		}(j)  //passing data to gorutine
			
		}
	}

	printNums(&wgSync)

	wgSync.Wait()
}