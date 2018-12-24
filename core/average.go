package core

type Average struct {
	Sum   float32
	Count uint64
}

func (avg *Average) Calculate() float32 {
	if avg.Count == 0 {
		return 0
	}
	return avg.Sum / float32(avg.Count)
}
