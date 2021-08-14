package publisher

// ConvertToProducerHeaders helps to parse headers defined as map to the struct
// representing a header in this pkg.
func ConvertToProducerHeaders(hs map[string]string) []EventHeader {
	headers := make([]EventHeader, 0)
	for k, v := range hs {
		headers = append(headers, EventHeader{
			Key:   k,
			Value: []byte(v),
		})
	}

	return headers
}
