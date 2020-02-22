# 简易协程池部分代码
```go
package buslogic
/**
*
* @ pool
* @Author: 
* @Date: 2019-10-18 16:24
 */

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"getUrl/app"
	"getUrl/data"
	"getUrl/relog"
	"github.com/jinzhu/gorm"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

/**
* @ pool
* @Author: licongfu@mioji.com
* @Date: 2019-09-27 17:22
 */

//定义一个任务结构
type task struct {
	f interface{}
}

//创建一个任务对象
func NewTask(f interface{}) *task {
	return &task{
		f: f,
	}
}

//定义一个线程池
type pool struct {
	timeout time.Duration //超时时间
	cap     int           //线程池的数量
	mode    int           //该参数应为需要传进来 所以
	Entry   chan *task    //线程池入口
	out     chan *task    //线程池出口
	job     chan *task
	control chan bool
	wg      *sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
}

//创建一个线程池
func NewPool(timeout time.Duration, cap, mode int) *pool {
	ctx, cancel := context.WithCancel(context.Background())
	return &pool{
		timeout: timeout,
		cap:     cap,
		mode:    mode,
		ctx:     ctx,
		cancel:  cancel,
		wg:      new(sync.WaitGroup),
		Entry:   make(chan *task, 200),
		out:     make(chan *task, 200),
		job:     make(chan *task, 100),
		control: make(chan bool, 2*runtime.NumCPU()),
	}
}

//运行线程池
func (p *pool) Run() {
	//日志记录
	defer func() {
		if r := recover(); r != nil {
			errDebug := strings.Split(string(debug.Stack()), "\n\t")
			fmt.Println("break out please fix:", errDebug)
		}
	}()
	//起多个线程处理任务
	for i := 0; i < p.cap; i++ {
		p.wg.Add(1)
		fmt.Printf("pool goroutine %d start\n", i)
		go p.workStart(cnn, i)
	}
	//起多个线程消耗任务去除数据
	for i := 0; i < 2*runtime.NumCPU(); i++ {
		p.wg.Add(1)
		p.control <- true
		fmt.Printf("goroutine %d start\n", i)
		go p.workout(i)
	}

	//不断将任务放入工作池中
	for {
		select {
		case <-p.ctx.Done():
			p.wg.Wait() //等待线程池所有线程释放
			return
		case task := <-p.Entry:
			p.job <- task
		}
	}
}

func (p *pool) Stop() {
	close(p.out)
	close(p.Entry)
	close(p.job)
	close(p.control)
}

//起一个线程运作
func (p *pool) workStart(cnn *gorm.DB, num int) {
	defer p.wg.Done()
			//入管道
	p.out <- NewTask(OutImgMsg{
		PoiId:       v.Id,
		Mode:        p.mode,
		Url:         v1.Tp,
		Recource_id: v1.Id,
		Source:      v1.Sk + "|" + v1.Sn,
		Stat:        app.Request(v1.Tp),
	})
	return nil
}

```

```go
//15个协程
	p := NewPool(20*time.Second, 30, mode)
	//放任务·
	go func() {
		for i := 1; i <= 30; i++ {
			p.Entry <- NewTask(arg1)//将任务传入到
		}
	}()
	p.Run()
	p.Stop()
```
	
#  结构图
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200221173910217.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzI0MTk0MTcz,size_16,color_FFFFFF,t_70)

# goroutine调度
详见：https://www.cnblogs.com/wdliu/p/9272220.html