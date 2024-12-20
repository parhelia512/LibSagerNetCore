package libcore

type LocalResolver interface {
	LookupIP(network string, domain string) (string, error)
}
