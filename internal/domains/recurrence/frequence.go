package recurrence

const (
	FrequenceDaily    = "DAILY"
	FrequenceWeekly   = "WEEKLY"
	FrequenceBiweekly = "BIWEEKLY"
	FrequenceMonthly  = "MONTHLY"
	FrequenceYearly   = "YEARLY"
)

type Frequence string

func (f Frequence) Validate() bool {
	switch f {
	case FrequenceDaily, FrequenceWeekly, FrequenceBiweekly, FrequenceMonthly, FrequenceYearly:
		return true
	}

	return false
}
