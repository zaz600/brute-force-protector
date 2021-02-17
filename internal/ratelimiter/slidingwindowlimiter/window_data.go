package slidingwindowlimiter

import "time"

type windowData struct {
	limit          int64
	window         time.Duration
	timestamps     []int64
	lastAccessTime time.Time
}

// add добавляет текущий таймстамп в коллекцию
func (s *windowData) add() {
	s.timestamps = append(s.timestamps, time.Now().UnixNano())
	s.lastAccessTime = time.Now()
}

// shrinkLeft удаляет элементы из слайса левее заданного.
// используется для удаления таймстампов выпавших из окна
func (s *windowData) shrinkLeft(idx int) {
	s.timestamps = s.timestamps[idx : len(s.timestamps)-1]
}

// leftBorder определяем первый слева элемент входящий в окно
func (s *windowData) leftBorder() int {
	windowLeft := time.Now().UnixNano() - s.window.Nanoseconds()
	for i, value := range s.timestamps {
		if value >= windowLeft {
			return i
		}
	}
	return 0
}

// currentSize количество элементов внутри текущего окна
// Заодно удаляет элементы выпавшие из окна
func (s *windowData) currentSize() int64 {
	if leftBorder := s.leftBorder(); leftBorder > 0 {
		s.shrinkLeft(leftBorder)
	}
	return int64(len(s.timestamps))
}

func newWindowData(limit int64, window time.Duration) *windowData {
	return &windowData{
		limit:          limit,
		window:         window,
		timestamps:     make([]int64, 0, limit),
		lastAccessTime: time.Time{},
	}
}
