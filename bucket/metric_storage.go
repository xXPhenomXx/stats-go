package bucket

type metricStorage struct {
	threshold uint
	metrics   map[string]map[string]uint
}

func newMetricStorage(threshold uint) *metricStorage {
	return &metricStorage{threshold: threshold, metrics: make(map[string]map[string]uint)}
}

func (s *metricStorage) LooksLikeID(firstSection, secondSection string) bool {
	if _, ok := s.metrics[firstSection]; !ok {
		s.metrics[firstSection] = make(map[string]uint, s.threshold)
	}

	// avoid storing all values to avoid memory loss
	if uint(len(s.metrics[firstSection])) < s.threshold {
		s.metrics[firstSection][secondSection]++
	}

	return uint(len(s.metrics[firstSection])) >= s.threshold
}
