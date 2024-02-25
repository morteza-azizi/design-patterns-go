package proxy

type Subject interface {
	Operation() string
}

type RealSubject struct{}

func (rs *RealSubject) Operation() string {
	return "RealSubject: Handling operation"
}

type Proxy struct {
	realSubject *RealSubject
}

func (proxy *Proxy) Operation() string {
	if proxy.realSubject == nil {
		proxy.realSubject = &RealSubject{}
	}

	result := "Proxy: Pre-processing\n"
	result += proxy.realSubject.Operation()
	result += "\nProxy: Post-processing"
	return result
}
