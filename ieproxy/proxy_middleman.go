package ieproxy

// Return original oKey or fallback fbKey if oKey doesn't exist in the map.
func mapFallback(oKey, fbKey string, m map[string]string) string {
	if v, ok := m[oKey]; ok {
		return v
	} else {
		return m[fbKey]
	}
}
