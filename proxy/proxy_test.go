package proxy

import (
	"strings"
	"testing"
)

func TestProxy(t *testing.T) {
	proxy := &Proxy{}
	result := proxy.Operation()

	expected := "Proxy: Pre-processing\nRealSubject: Handling operation\nProxy: Post-processing"
	if result != expected {
		t.Errorf("Proxy request result is incorrect, got: %s, want: %s", result, expected)
	}
}

func TestProxyLazyInitialization(t *testing.T) {
	proxy := &Proxy{}

	if proxy.realSubject != nil {
		t.Errorf("RealSubject is initialized before request")
	}

	proxy.Operation()

	if proxy.realSubject == nil {
		t.Errorf("RealSubject is not initialized after request")
	}
}

func TestRealSubjectOperation(t *testing.T) {
	realSubject := &RealSubject{}
	result := realSubject.Operation()

	expected := "RealSubject: Handling operation"
	if result != expected {
		t.Errorf("RealSubject request result is incorrect, got: %s, want: %s", result, expected)
	}
}

func TestProxyOperationWithAdditionalFunctionality(t *testing.T) {
	proxy := &Proxy{}
	result := proxy.Operation()

	if !strings.Contains(result, "Proxy: Pre-processing") || !strings.Contains(result, "Proxy: Post-processing") {
		t.Errorf("Additional functionality not added by the proxy")
	}
}
