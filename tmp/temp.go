package tmp

import (
	"math/big"
	"sync"
)

type CList struct {
	mtx    sync.RWMutex
	wg     *sync.WaitGroup
	waitCh chan struct{}
	head   *CElement // first element
	tail   *CElement // last element
	len    int       // list length
	maxLen int       // max list length

}

type CElement struct {
	mtx        sync.RWMutex
	prev       *CElement
	prevWg     *sync.WaitGroup
	prevWaitCh chan struct{}
	next       *CElement
	nextWg     *sync.WaitGroup
	nextWaitCh chan struct{}
	removed    bool

	Nonce    uint64
	GasPrice *big.Int
	Address  string

	Value interface{} // immutable
}

func (l *CList) CutHeadN(n int) {
	ll := l.Len()
	if ll <= n {
		if ll > 0 {
			l.Init()
		}
		return
	}
	l.mtx.Lock()
	defer l.mtx.Unlock()
	ele := l.head
	for i := 0; i < n; i++ {
		ele = ele.next
	}
	l.head = ele
	l.head.prev = nil

	// Update l.len
	l.len -= n
}

func (l *CList) Init() *CList {
	l.mtx.Lock()

	l.wg = waitGroup1()
	l.waitCh = make(chan struct{})
	l.head = nil
	l.tail = nil
	l.len = 0
	l.mtx.Unlock()
	return l
}

func (l *CList) Len() int {
	l.mtx.RLock()
	len := l.len
	l.mtx.RUnlock()
	return len
}

func waitGroup1() (wg *sync.WaitGroup) {
	wg = &sync.WaitGroup{}
	wg.Add(1)
	return
}
