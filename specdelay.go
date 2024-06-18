package cron

import "time"

// based on SpecSchedule, extend to include delay
type SpecDelaySchedule struct {
	Delay        uint64
	SpecSchedule *SpecSchedule
}

func (s *SpecDelaySchedule) Next(t time.Time) time.Time {
	next := s.SpecSchedule.Next(t)
	if next.IsZero() {
		return next
	}

	return next.Add(time.Second * time.Duration(s.Delay))
}
