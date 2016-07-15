# metrics
forked from https://github.com/facebookgo/metrics and added some useful types

When I use facebook/metrics I find the type of it support just for int64, but I should take all the events and error-logs in my developed modular. So I add the flow(流) type.

#Usage

##Fiisio := metrics.NewFlow(N)

创建一个可以最大容纳N个字符串的Flow类型，N是int64类型，当push进去超过N个则，第一个会被挤下来，Flow的数据个数始终维持<=N的状态。


##Fiisio.RPush(str)/Fiisio.LPush(str)

从左边或者右边Push进去str


##Fiisio.RPop()/Fiisio.LPop()

从左边或者右边Pop出去一条数据


##Fiisio.Values()

获取Flow中所有的数据，返回值类型[][]byte


##Fiisio.Size()

获取Fiisio中数据个数


##Fiisio.Clear()

清楚Fiisio中所有数据

