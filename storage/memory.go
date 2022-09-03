package storage

var _ Storage = (*DefaultStorage)(nil)
var emptyStruct struct{}

// DefaultStorage defines a default Storage interface.
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
	s.Truncate()
}

func (s *DefaultStorage) Truncate() {
	dummyHashTable := map[string]struct{}{
		"": {},
	}
	cleanFrequencyLookupTable := make([]map[string]struct{}, 1)
	s.FrequencyLookupTable = append(cleanFrequencyLookupTable, dummyHashTable)
	s.IpAddressTallyMap = make(map[string]int)
}

// Insert implements the Storage Insert interface
func (s *DefaultStorage) Insert(ip_address string) {
	var initialValue int

	if val, ok := s.IpAddressTallyMap[ip_address]; ok {
		initialValue = val
	}
	finalValue := initialValue + 1
	s.IpAddressTallyMap[ip_address] = finalValue

	if len(s.FrequencyLookupTable) == finalValue {
		newHashTable := map[string]struct{}{
			ip_address: {},
		}
		s.FrequencyLookupTable = append(s.FrequencyLookupTable, newHashTable)
	} else {
		s.FrequencyLookupTable[finalValue][ip_address] = emptyStruct
	}
	if finalValue > 1 {
		delete(s.FrequencyLookupTable[initialValue], ip_address)
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
	return output
}
