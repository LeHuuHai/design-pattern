package model

type ProxyProfile struct {
	ID      uint
	profile *Profile
}

func (p *ProxyProfile) Get() []byte {
	if p.profile == nil {
		p.profile = &Profile{ID: p.ID}
		p.profile.Load()
	}
	return p.profile.Get()
}
