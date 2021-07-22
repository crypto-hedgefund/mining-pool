package stats

type StatisticsService struct {
	miners map[string]*Entry
}

type Entry struct {
}

func NewStatsService() *StatisticsService {
	return &StatisticsService{}
}

func (s *StatisticsService) Start() {

}
