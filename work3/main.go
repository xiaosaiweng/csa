package main

/**
程序：程序只是一组指令的有序集合，它本身没有任何运行的含义，它只是一个静态的实体
进程：执行中的程序叫做进程，是一个动态的概念。
线程：
协程：指在后台中运行的轻量级执行线程，go 协程是 Go 中实现并发的关键组成部分
*/
//启动协程----go 函数名(参数)
//func main(){
//	go repeat("gugugu",100)
//	repeat("阿巴阿巴",5)
//}
//func repeat(str string,count int){
//	for i :=0;i<count;i++{
//		fmt.Println(str)
//		time.Sleep(time.Microsecond)
//	}
//}

//------------------未加锁出现问题-------------------
//var myRes=make(map[int]int,20)
//
//func main(){
//	for i:=1;i<=100;i++{
//		go factorial(i)
//	}
//}
//func factorial(n int){
//	var res=1
//	for i:=1;i<=n;i++{
//		res *=i
//	}
//	myRes[n]=res
//}

//------------------加锁-----------------------
//var (
//	myRes=make(map[int]int,20)
//	lock sync.Mutex
//)
//
//func main(){
//	for i:=1;i<=100;i++{
//		go factorial(i)
//	}
//}
//func factorial(n int){
//	var res=1
//	for i:=1;i<=n;i++{
//		res *=i
//	}
//	lock.Lock()
//	myRes[n]=res
//	lock.Unlock()
//}

/**
channel
创建channel:   ch := make(chan type)
使用channel：
		ch <- x  // 将变量x发送给管道ch
		x = <-ch // 变量x接受管道ch发送过来的值
		<-ch     // 接受管道ch发送过来的值，但不使用
不带缓存的channel的声明方式：
		ch = make(chan int) //可以接收和发送类型为 type 的数据
		var ch chan int  // 可以接收和发送类型为 type 的数据
*/
//-------------------------主协程等待的---------------------------------
//func main(){
//		var channel =make(chan int)
//		var send=6666
//		// 关键字go后跟的是需要被并发执行的代码块，它由一个匿名函数代表
//		// 在这里，我们只要知道在花括号中的就是将要被并发执行的代码就可以
//		go func(){
//			channel<-send
//			fmt.Println("数据已发送")
//		}()
//		// 这里让主线程休眠1秒钟
//		// 以使上面的goroutine有机会被执行
//		time.Sleep(1 *time.Second)
//}
//----------------------------协程执行后，再通知主协程退出------------------------
//func main(){
//	var channel=make(chan int)
//	var send=6666
//	var receive int
//	go func() {
//		channel <-send
//		fmt.Println("数据已发送")
//	}()
//	receive = <-channel
//	fmt.Println(receive)
//}
/**
带缓存的channel
创建了一个可以持有三个整型元素的带缓存Channel:
	ch = make(chan int, 3)
无阻塞的向`ch`传入三个int:
	ch <- 1
	ch <- 2
	ch <- 3
传入第四个int类型的时候，当前协程就会被阻塞
*/
//---------------------带缓存的channel-------------------------
//func main(){
//	var channel=make(chan int ,3)
//	go func() {
//		channel <- 1
//		channel <- 2
//		channel <- 3
//		fmt.Println("我发送了3个数据")
//		channel <- 4
//		fmt.Println("我发送了第4个数据")
//	}()
//	time.Sleep(time.Second)
//}
//------------------不带缓存的channel，在调用`ch<-`的时候会阻塞，直到`<-ch`被调用继续-------------------------
//func main()  {
//	st := time.Now()
//	ch := make (chan bool)
//	go func() {
//		time.Sleep(time.Second *2)
//		<- ch
//	}()
//	ch <-true
//		fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds())
//	time.Sleep(time.Second*5)
//}
//---------------------稍微改动一点点-------------------
//func main()  {
//	st := time.Now()
//	ch := make (chan bool,1)
//	go func() {
//		time.Sleep(time.Second *2)
//		<- ch
//	}()
//	ch <-true
//	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds())
//	time.Sleep(time.Second*5)
//}
/**
`range`，可以用在channel中
`close`也就是golang的内置函数，可以用于关闭一个channel，
	注意:
		1.对一个已经被close过的channel进行接收操作依然可以接受到之前已经成功发送的数据
		2.对于一个已经关闭的管道，我们不能进行发送，也不能再次关闭，不然程序会直接抛出panic
 */
//func fibonacci(n int ,c chan int){
//	x ,y := 1,1
//	for i:=0;i<n;i++{
//		c <-x
//		x,y=y,x+y
//	}
//	close(c)
//}
//func main()  {
//	c:=make(chan int,10)
//	go fibonacci(cap(c),c)
//	for i:=range c{
//		fmt.Println(i)
//	}
//}
//-----------------------通道死锁--------------------------------
//func main()  {
//	out :=make(chan int)
//	out <-2
//	go fl(out)
//}
//func fl(in chan int)  {
//	fmt.Println(<-in)
//}
/**
	select:
	如果多个case同时就绪时，select会随机地选择一个执行，这样来保证每一个channel都有平等的被select的机会

	如果没有case需要处理，则会选择`default`去处理，若没有`default`语句时则`select`语句会被阻塞（`select`默认为阻塞），直到某个case需要处理
 */
//-------------------------------select------------------------------
//func fibonacci(ch,quit chan int)  {
//	x,y:=0,1
//	for{
//		select {
//		case ch<-x:
//			x,y=y,x+y
//		case <-quit:
//			fmt.Println("quit")
//			return
//		}
//	}
//}
//func main()  {
//	c:=make(chan int)
//	quit:=make(chan int)
//	go func() {
//		for i:=0;i<10;i++{
//			fmt.Println(<-c)
//		}
//		quit<-0
//	}()
//	fibonacci(c,quit)
//}
//-------------------------------- select 完成随机选择0，1的的程序--------------------------------
//func main()  {
//	for i:=range random(100){
//		fmt.Println(i)
//	}
//}
//func random(n int) <-chan int  {
//	c:=make(chan int)
//	go func() {
//		defer close(c)
//		for i:=0;i<n;i++{
//			select {
//			case c<-0:
//			case c<-1:
//			}
//		}
//	}()
//	return c
//}
//----------------等待所有子协程1、使用channel------------------------
//const count = 5
//func main(){
//	waitCh :=make(chan struct{},count)// 缓存要和并发数一样大
//	for i:=0;i<count;i++{
//		go func(x int) {
//			fmt.Println("finishing",x)
//			waitCh<- struct{}{}
//		}(i)
//	}
//	for i:=0;i<count;i++{
//		<-waitCh
//	}
//	fmt.Println("main goroutine finishing")
//}
//----------------等待所有子协程2、sync.WaitGroup------------------------
//var wg sync.WaitGroup
//const count = 5
//
//func main()  {
//	for i:=0;i<count;i++{
//		wg.Add(1)
//		go func(x int) {
//			defer wg.Done()
//			fmt.Println("finishing",x)
//		}(i)
//	}
//	wg.Wait()
//	fmt.Println("main groutine finish")
//}
//--------------------------唤醒所有子协程1、使用channel----------------------------
//var wg sync.WaitGroup
//const count =5
//
//func main()  {
//	doneCh :=make(chan struct{})
//	for i:=0;i<count;i++{
//		wg.Add(1)
//		go func(x int) {
//			defer wg.Done()
//			<-doneCh
//			fmt.Println("finishing",x)
//		}(i)
//	}
//	for i:=0;i<count;i++{
//		doneCh<- struct{}{}
//	}
//	wg.Wait()
//	fmt.Println("main groutine finish")
//}
//--------------------------唤醒所有子协程2、done channel pattern----------------------------
//var wg sync.WaitGroup
//const count=5
//
//func main()  {
//	doneCh:=make(chan struct{})
//	for i:=0;i<count;i++{
//		wg.Add(1)
//		go func(x int) {
//			defer wg.Done()
//			<-doneCh
//			fmt.Println("finishing",x)
//		}(i)
//	}
//	close(doneCh)
//	wg.Wait()
//	fmt.Println("main goroutine finish")
//}
//--------------------------唤醒所有子协程3、条件变量sync.Cond----------------------------
//var wg sync.WaitGroup
//const count =5
//
//func main()  {
//	doneCh:=make(chan struct{})
//	lock:=sync.Mutex{}
//	cond:=sync.NewCond(&lock)
//	for i:=0;i<count;i++{
//		wg.Add(1)
//		go func(x int) {
//			defer wg.Done()
//			defer fmt.Println("finishing",x)
//			for{
//				cond.L.Lock()
//				cond.Wait()
//				cond.L.Unlock()
//				select {
//				case <-doneCh:
//					return
//				default:
//				}
//				fmt.Printf("%d wakeup\n\n", x)
//			}
//		}(i)
//	}
//	for i:=0;i<3;i++{
//		time.Sleep(1*time.Second)
//		cond.Broadcast()
//		fmt.Println("broadcast",i)
//	}
//	close(doneCh)
//	cond.Broadcast()
//	wg.Wait()
//	fmt.Println("main goroutine finish")
//}
//--------------------------唤醒所有子协程3、条件变量sync.Cond、增加一个间接层----------------------------
//var wg sync.WaitGroup
//const count =5
//
//func main()  {
//	doneCh:=make(chan struct{})
//	lock:=sync.Mutex{}
//	cond:=sync.NewCond(&lock)
//	waitCond := func() (condCh chan struct{}){
//		condCh =make(chan struct{})
//		wg.Add(1)
//		go func(){
//			defer wg.Done()
//			for{
//				cond.L.Lock()
//				cond.Wait()
//				cond.L.Unlock()
//				select {
//				case <-doneCh:
//					return
//				case condCh <- struct{}{}:
//				}
//			}
//		}()
//		return
//	}
//	for i:=0;i<count;i++{
//		wg.Add(1)
//		go func(x int) {
//			defer wg.Done()
//			defer fmt.Println("finishing",x)
//			condCh:=waitCond()
//			for{
//				select {
//				case <-doneCh:
//					return
//				case <-condCh:
//				}
//				fmt.Printf("%d wakeup\n", x)
//			}
//		}(i)
//	}
//	for i:=0;i<3;i++{
//		time.Sleep(1*time.Second)
//		cond.Broadcast()
//		fmt.Println("broadcast",i)
//	}
//	close(doneCh)
//	cond.Broadcast()
//	wg.Wait()
//	fmt.Println("main goroutine finish")
//}
//---------------------------关闭所有子协程1、使用退出channel退出------------------------------------
//func worker(stopCh <-chan struct{},t *time.Ticker){
//	go func() {
//		defer fmt.Println("worker exit")
//		for{
//			select {
//			case <-stopCh:
//				fmt.Println("Recv stop signal")
//				return
//			case <-t.C:
//				fmt.Println("Working .")
//			}
//		}
//	}()
//	return
//}
//---------------------------关闭所有子协程2、使用context------------------------------------
//func worker(ctx context.Context,t *time.Ticker){
//	go func() {
//		defer fmt.Println("worker exit")
//		for{
//			select {
//			case <-ctx.Done():
//				fmt.Println("recv stop signal")
//				return
//			case <-t.C:
//				fmt.Println("Working .")
//			}
//		}
//	}()
//	return
//}
//func main()  {
//	ticker:=time.NewTicker(time.Second)
//	ctx ,cancel:=context.WithCancel(context.Background())
//	go worker(ctx,ticker)
//	time.Sleep(3*time.Second)
//	cancel()
//	time.Sleep(50*time.Millisecond)
//}
//---------------------gin-------------------------------------------
//import "github.com/gin-gonic/gin"
//func main(){
//	router :=gin.Default()
//	router.GET("/strong",func(ctx *gin.Context){
//		ctx.String(200,"xpy")
//	})
//	router.Run(":8080")
//}