package exmaple2

// 友情链接 https://github.com/asong2020/go-algorithm/tree/master/snowFlake
import (
	"errors"
	"sync"
	"time"
)

const (
	workerIDBits     = uint64(5) // 10bit 工作机器ID中的 5bit workerID
	dataCenterIDBits = uint64(5) // 10 bit 工作机器ID中的 5bit dataCenterID
	sequenceBits     = uint64(12)

	maxWorkerID     = int64(-1) ^ (int64(-1) << workerIDBits) //节点ID的最大值 用于防止溢出
	maxDataCenterID = int64(-1) ^ (int64(-1) << dataCenterIDBits)
	maxSequence     = int64(-1) ^ (int64(-1) << sequenceBits)

	timeLeft = uint8(22) // timeLeft = workerIDBits + sequenceBits // 时间戳向左偏移量
	dataLeft = uint8(17) // dataLeft = dataCenterIDBits + sequenceBits
	workLeft = uint8(12) // workLeft = sequenceBits // 节点IDx向左偏移量
	// 2020-05-20 08:00:00 +0800 CST
	twepoch = int64(1529923200000) // 常量时间戳(毫秒)
)

type Worker struct {
	mu           sync.Mutex
	LastStamp    int64
	WorkerId     int64
	DataCenterId int64
	Sequence     int64
}

func NewWorker(workerId, dataCenterId int64) *Worker {
	return &Worker{
		LastStamp:    0,
		WorkerId:     workerId,
		DataCenterId: dataCenterId,
		Sequence:     0,
	}
}

func (w *Worker) NextID() (uint64, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.nextId()
}

func (w *Worker) getMilliSecond() int64 {
	return time.Now().UnixNano() / 1e6
}

func (w *Worker) nextId() (uint64, error) {

	timeStamp := w.getMilliSecond()
	if timeStamp < w.LastStamp {
		return 0, errors.New("time is moving backwards ,waiting until")
	}

	if w.LastStamp == timeStamp {
		w.Sequence = (w.Sequence + 1) & maxSequence
		if w.Sequence == 0 {
			for timeStamp <= w.LastStamp {
				timeStamp = w.getMilliSecond()
			}
		}
	} else {
		w.Sequence = 0
	}

	w.LastStamp = timeStamp
	w.LastStamp = timeStamp
	id := ((timeStamp - twepoch) << timeLeft) |
		(w.DataCenterId << dataLeft) |
		(w.WorkerId << workLeft) |
		w.Sequence
	return uint64(id), nil
}
