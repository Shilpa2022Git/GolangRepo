package main

import (
	"fmt"
	"time"
)

//sends data into channel
func genMessage(ch chan<- string){
	//send message on channel ch
	ch <- "message"
}

func relayMessage(sendCh <-chan string, recCh chan<- string){

	//receive message on sendCh
	m := <-sendCh

	//send on recCh
	recCh <- m
}

func main(){

	//channel creation
	ch := make(chan int)  //syntax make(chan type)
	ch1 := make(chan int)
	go func(a, b int){
		c := a+b
		ch <- c
		ch1 <- c
	}(1,2)

	var res int
	res = <-ch

	//short hand operation to recive data from channel, once data recieved from channel, data is lost
	r := <-ch1

	fmt.Println("Received value from channel ", res)
	fmt.Println("Short hand op ", r)

	//Buffered channel
	chCount := make(chan int, 6) //make(chan type, capacity)
	go func(){		
		defer close(chCount)
		for i:=0; i<6;i++{
			chCount <- i
		}
		//close(chCount)  //if we don't close the channel receiving end went into deadlock situation.
						//fatal error: all goroutines are asleep - deadlock!, goroutine 1 [chan receive]:
						//best practice to use defer to close channel. goroutine will create the channel and return receiver channel.
	}()

	//range over channel to receive values
	for v := range chCount{
		fmt.Println("Value from buffered channel ", v)
	}

	//Channel direction. in channel which send data only and out to receive data
	//syntax in <-chan [send]  out chan<- [receive]

	chS := make(chan string)
	chR := make(chan string)
	go genMessage(chS)
	go relayMessage(chS, chR)
	
	r1 := <-chR

	fmt.Println("Relayed ", r1)

	//Ownership of channel
	owner := func() (ch <-chan int){
		
		chO := make(chan int, 5)
		
		go func () {
			defer close(chO)
			for i:=0;i<5;i++{
				chO <- i
			}
		}()

		return chO
	}

	retCh := owner()
	for v := range retCh{
		fmt.Println("Returned ch ", v)
	}

	chSel1 := make(chan string)
	chSel2 := make(chan string)

	go func(){
		time.Sleep(1*time.Second)
		chSel1 <- "One"		
	}()

	go func(){
		time.Sleep(1*time.Second)
		chSel2 <- "Two"		
	}()

	//select function is to wait on channels. when channel gt some value select stmt will execute and read teh data fro channel and perform
	for i:=0;i<2;i++{
		select{
			case m1 := <-chSel1:
				fmt.Println("Message 1:", m1)				
			case m2 := <-chSel2:
				fmt.Println("Message 2:", m2)
		}
	}

	//use of timeout in select

	chTime := make(chan string)
	go func(){
		time.Sleep(2*time.Second)
		chTime <- "Filling"
	}()

	select{
	case m1 := <-chTime:
		fmt.Println("printed after given time ",m1)
	case <-time.After(1*time.Second):  //waiting time, after that it will release and execute instructions
		fmt.Println("Timeout")
	}

	//We can implement non blocking communication using select and default option. Altough message is not received
	//communication will not block and  execute defualt block untill it receives any message

	chDef := make(chan string)
	go func (){
		time.Sleep(2*time.Second)
		chDef <- "messsage"
	}()

	for i:=0;i<2;i++{
		select{
		case res1 := <-chDef:
			fmt.Println(res1)
		default:
			fmt.Println("do something else untill get the data")
		}
	}

	fmt.Println("After default select")
	time.Sleep(5*time.Second)
	fmt.Println("Exiting application")
}