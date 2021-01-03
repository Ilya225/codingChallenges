package main

// https://leetcode.com/problems/time-based-key-value-store/

type PairElement struct {
	Value string
	Timestamp int
}

type TimeMap struct {
	KVStore map[string][]*PairElement
}


/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		KVStore: map[string][]*PairElement{},
	}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
	queue := this.KVStore[key]

	if queue == nil {
		queue = []*PairElement{}
	}

	queue = append(queue, &PairElement{Value: value, Timestamp: timestamp})

	this.KVStore[key] = queue

}


func (this *TimeMap) Get(key string, timestamp int) string {
	queue, ok := this.KVStore[key]
	if !ok {
		return ""
	}
	l := len(queue)
	if l == 0 {
		return ""
	}

	if queue[0].Timestamp > timestamp {
		return ""
	}

	mid := l / 2
	end := l-1

	nxt := mid
	for nxt < end && nxt > 0 {
		if queue[nxt].Timestamp == timestamp {
			return queue[nxt].Value
		} else if queue[nxt].Timestamp > timestamp {
			end = nxt
			nxt = nxt/2
		} else {
			if nxt == 1 {
				nxt += 1
				continue
			}
			nxt = nxt + nxt/2
		}
	}
	if nxt >= l {
		return queue[l-1].Value
	}

	if queue[nxt].Timestamp > timestamp {
		return queue[nxt-1].Value
	}
	return queue[nxt].Value
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

func timeMapApp( ) {
	timeMap := Constructor()
	timeMap.Set("foo", "bar2" , 4)
	timeMap.Set("foo", "bar2" , 6)
	timeMap.Set("foo", "bar2" , 7)
	timeMap.Set("foo", "bar2" , 12)
	timeMap.Set("foo", "bar2" , 15)
	println(timeMap.Get("foo", 4))
	println(timeMap.Get("foo", 5))
	println(timeMap.Get("foo", 8))
}