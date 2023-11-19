package adapter

import "testing"

func TestAdapter(t *testing.T) {
	message := "Hello World!"

	adapter := PrinterAdapter{HpLegacyPrinter: &MyLegacyPrinter{}, Message: message}
	returnMessage := adapter.PrintMessage()
	if returnMessage != "Legacy Printer: Adapter: Hello World!\n" {
		t.Errorf("Message didn't match: %s\n", returnMessage)
	}

	adapter = PrinterAdapter{HpLegacyPrinter: nil, Message: message}
	returnMessage = adapter.PrintMessage()
	if returnMessage != "Hello World!" {
		t.Errorf("Message didn't match: %s\n", returnMessage)
	}
}
