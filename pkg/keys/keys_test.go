package keys

import "testing"

func TestAddNewKey(t *testing.T) {
	tests := []struct {
		name     string
		password string
	}{
		{
			name:     "TestAddNewKey - abc12356",
			password: "abc12356",
		},
		{
			name:     "TestAddNewKey - abc12357",
			password: "abc12357",
		},
		{
			name:     "TestListKeys",
			password: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.password != "" {
				AddNewKey(tt.password)
			}

			ListKeys()
		})
	}
}
