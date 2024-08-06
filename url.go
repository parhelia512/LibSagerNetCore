package libcore

import (
	"net"
	"net/url"
	"strconv"
	"strings"
	_ "unsafe"
)

type URL interface {
	GetScheme() string
	SetScheme(scheme string)
	GetOpaque() string
	SetOpaque(opaque string)
	GetUsername() string
	SetUsername(username string)
	GetPassword() string
	SetPassword(password string) error
	GetHost() string
	SetHost(host string)
	GetPort() int32
	SetPort(port int32)
	GetPath() string
	SetPath(path string)
	GetRawPath() string
	SetRawPath(rawPath string) error
	QueryParameterNotBlank(key string) string
	AddQueryParameter(key, value string)
	GetFragment() string
	SetRawFragment(rawFragment string) error
	GetString() string
}

var _ URL = (*netURL)(nil)

type netURL struct {
	url.URL
	url.Values
}

func NewURL(scheme string) URL {
	u := new(netURL)
	u.Scheme = scheme
	u.Values = make(url.Values)
	return u
}

//go:linkname setFragment net/url.(*URL).setFragment
func setFragment(u *url.URL, fragment string) error

//go:linkname setPath net/url.(*URL).setPath
func setPath(u *url.URL, fragment string) error

func ParseURL(rawURL string) (URL, error) {
	u := &netURL{}
	ru, frag, _ := strings.Cut(rawURL, "#")
	uu, err := url.Parse(ru)
	if err != nil {
		return nil, newError("failed to parse url: ", rawURL).Base(err)
	}
	u.URL = *uu
	u.Values = u.Query()
	if u.Values == nil {
		u.Values = make(url.Values)
	}
	if frag == "" {
		return u, nil
	}
	if err = u.SetRawFragment(frag); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *netURL) GetScheme() string {
	return u.Scheme
}

func (u *netURL) SetScheme(scheme string) {
	u.Scheme = scheme
}

func (u *netURL) GetOpaque() string {
	return u.Opaque
}

func (u *netURL) SetOpaque(opaque string) {
	u.Opaque = opaque
}

func (u *netURL) GetUsername() string {
	if u.User != nil {
		return u.User.Username()
	}
	return ""
}

func (u *netURL) SetUsername(username string) {
	if u.User != nil {
		if password, ok := u.User.Password(); !ok {
			u.User = url.User(username)
		} else {
			u.User = url.UserPassword(username, password)
		}
	} else {
		u.User = url.User(username)
	}
}

func (u *netURL) GetPassword() string {
	if u.User != nil {
		if password, ok := u.User.Password(); ok {
			return password
		}
	}
	return ""
}

func (u *netURL) SetPassword(password string) error {
	if u.User == nil {
		return newError("set username first")
	}
	u.User = url.UserPassword(u.User.Username(), password)
	return nil
}

func (u *netURL) GetHost() string {
	return u.Hostname()
}

func (u *netURL) SetHost(host string) {
	_, port, err := net.SplitHostPort(u.Host)
	if err == nil {
		u.Host = net.JoinHostPort(host, port)
	} else {
		u.Host = host
	}
}

func (u *netURL) GetPort() int32 {
	portStr := u.Port()
	if portStr == "" {
		return 0
	}
	port, _ := strconv.Atoi(portStr)
	return int32(port)
}

func (u *netURL) SetPort(port int32) {
	host, _, err := net.SplitHostPort(u.Host)
	if err == nil {
		u.Host = net.JoinHostPort(host, strconv.Itoa(int(port)))
	} else {
		u.Host = net.JoinHostPort(u.Host, strconv.Itoa(int(port)))
	}
}

func (u *netURL) GetPath() string {
	return u.Path
}

func (u *netURL) SetPath(path string) {
	u.Path = path
	u.RawPath = ""
}

func (u *netURL) GetRawPath() string {
	return u.RawPath
}

func (u *netURL) SetRawPath(rawPath string) error {
	return setPath(&u.URL, rawPath)
}

func (u *netURL) QueryParameterNotBlank(key string) string {
	return u.Get(key)
}

func (u *netURL) AddQueryParameter(key, value string) {
	u.Add(key, value)
}

func (u *netURL) GetFragment() string {
	return u.Fragment
}

func (u *netURL) SetRawFragment(rawFragment string) error {
	return setFragment(&u.URL, rawFragment)
}

func (u *netURL) GetString() string {
	u.RawQuery = u.Encode()
	return u.String()
}
