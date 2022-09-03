package storage

var _ Storage = (*DefaultStorage)(nil)
var emptyStruct struct{}

// A DefaultStorage defines a default Storage interface.
// This is the starting point of extending the library to work with other storage.
type DefaultStorage struct {
	IpAddressTallyMap    map[string]int
	FrequencyLookupTable []map[string]struct{}
}

// NewStorage initializes a DefaultStorage
func NewStorage() *DefaultStorage {
	storage := &DefaultStorage{}
	storage.Init()
	return storage
}

// Init implements the Storage Init interface
func (s *DefaultStorage) Init() {
	cleanFrequencyLookupTable := make([]map[string]struct{}, 1)
	s.FrequencyLookupTable = cleanFrequencyLookupTable
	s.IpAddressTallyMap = make(map[string]int)
}

func (s *DefaultStorage) Truncate() {
	s.Init()
}

// Insert implements the Storage Insert interface
func (s *DefaultStorage) Insert(ipAddress string) {
	var initialValue int

	if val, ok := s.IpAddressTallyMap[ipAddress]; ok {
		initialValue = val
	}
	finalValue := initialValue + 1
	s.IpAddressTallyMap[ipAddress] = finalValue

	if len(s.FrequencyLookupTable) == finalValue {
		newHashTable := map[string]struct{}{
			ipAddress: {},
		}
		s.FrequencyLookupTable = append(s.FrequencyLookupTable, newHashTable)
	} else {
		s.FrequencyLookupTable[finalValue][ipAddress] = emptyStruct
	}
	if finalValue > 1 {
		delete(s.FrequencyLookupTable[initialValue], ipAddress)
	}
}

// Fetch implements the Storage Fetch interface
func (s *DefaultStorage) Fetch(limit int64) []string {
	output := make([]string, limit) // pre-allocate needed memory to ensure amortized insert time.

	outputIterator := 0
	maxArrayCursor := 0
	for outputIterator < int(limit) && len(s.FrequencyLookupTable) > maxArrayCursor {
		freqHashTable := s.FrequencyLookupTable[len(s.FrequencyLookupTable)-maxArrayCursor-1]
		if len(freqHashTable) != 0 {
			for k := range freqHashTable {
				output[outputIterator] = k
				outputIterator++
				if outputIterator == int(limit) {
					break
				}
			}
		}
		maxArrayCursor++
	}
	return output[0:outputIterator] //removes unused pre-allocated space
}
