package main

// 真正做初始化
//func newTimingWheelWithClock(interval time.Duration, numSlots int, execute Execute, ticker timex.Ticker) (
//	*TimingWheel, error) {
//	tw := &TimingWheel{
//		interval:      interval,                     // 单个时间格时间间隔
//		ticker:        ticker,                       // 定时器，做时间推动，以interval为单位推进
//		slots:         make([]*list.List, numSlots), // 时间轮
//		timers:        NewSafeMap(),                 // 存储task{key, value}的map [执行execute所需要的参数]
//		tickedPos:     numSlots - 1,                 // at previous virtual circle
//		execute:       execute,                      // 执行函数
//		numSlots:      numSlots,                     // 初始化 slots num
//		setChannel:    make(chan timingEntry),       // 以下几个channel是做task传递的
//		moveChannel:   make(chan baseEntry),
//		removeChannel: make(chan interface{}),
//		drainChannel:  make(chan func(key, value interface{})),
//		stopChannel:   make(chan lang.PlaceholderType),
//	}
//	// 把 slot 中存储的 list 全部准备好
//	tw.initSlots()
//	// 开启异步协程，使用 channel 来做task通信和传递
//	go tw.run()
//
//	return tw, nil
//}
//
//func (tw *TimingWheel) run() {
//	for {
//		select {
//		// 定时器做时间推动 -> scanAndRunTasks()
//		case <-tw.ticker.Chan():
//			tw.onTick()
//			// add task 会往 setChannel 输入task
//		case task := <-tw.setChannel:
//			tw.setTask(&task)
//			...
//		}
//	}
//}
//
//func (c *Cache) Set(key string, value interface{}) {
//	c.lock.Lock()
//	_, ok := c.data[key]
//	c.data[key] = value
//	c.lruCache.add(key)
//	c.lock.Unlock()
//
//	expiry := c.unstableExpiry.AroundDuration(c.expire)
//	if ok {
//		c.timingWheel.MoveTimer(key, expiry)
//	} else {
//		c.timingWheel.SetTimer(key, value, expiry)
//	}
//}
//
//
//func (tw *TimingWheel) moveTask(task baseEntry) {
//	// timers: Map => 通过key获取 [positionEntry「pos, task」]
//	val, ok := tw.timers.Get(task.key)
//	if !ok {
//		return
//	}
//
//	timer := val.(*positionEntry)
//	// {delay < interval} => 延迟时间比一个时间格间隔还小，没有更小的刻度，说明任务应该立即执行
//	if task.delay < tw.interval {
//		threading.GoSafe(func() {
//			tw.execute(timer.item.key, timer.item.value)
//		})
//		return
//	}
//	// 如果 > interval，则通过 延迟时间delay 计算其出时间轮中的 new pos, circle
//	pos, circle := tw.getPositionAndCircle(task.delay)
//	if pos >= timer.pos {
//		timer.item.circle = circle
//		// 记录前后的移动offset。为了后面过程重新入队
//		timer.item.diff = pos - timer.pos
//	} else if circle > 0 {
//		// 转移到下一层，将 circle 转换为 diff 一部分
//		circle--
//		timer.item.circle = circle
//		// 因为是一个数组，要加上 numSlots [也就是相当于要走到下一层]
//		timer.item.diff = tw.numSlots + pos - timer.pos
//	} else {
//		// 如果 offset 提前了，此时 task 也还在第一层
//		// 标记删除老的 task，并重新入队，等待被执行
//		timer.item.removed = true
//		newItem := &timingEntry{
//			baseEntry: task,
//			value:     timer.item.value,
//		}
//		tw.slots[pos].PushBack(newItem)
//		tw.setTimerPosition(pos, newItem)
//	}
//}
