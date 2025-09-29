package wrapt

import (
	"testing"
)

func TestWrapt_Wrap(t *testing.T) {
	tests := []struct {
		name         	string
		input        	string
		terminalWidth 	int
		expected     	string
	}{
		{
			name:         "No Breaks",
			input:        "hello world",
			terminalWidth: 20,
			expected:     "hello world",
		},
		{
			name:         "One Break",
			input:        "this is a longer string that needs to be wrapped",
			terminalWidth: 20,
			expected:     "this is a longer \nstring that needs \nto be wrapped",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Wrap(tt.input, tt.terminalWidth)
		
			if result != tt.expected {
				t.Errorf("\nResult: %v\n, Expected: %v", result, tt.expected)
			}	
		})
	}
}

func TestWrapt_WrapArray(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		marginSize   int
		terminalWidth int
		expected     []string
	}{
		{
			name:         "No Margin, One Line",
			input:        "hello world",
			marginSize:   0,
			terminalWidth: 20,
			expected:     []string{"hello world"},
		},
		{
			name:         "No Margin, Array Wrap",
			input:        "this is a longer string that needs to be wrapped",
			marginSize:   0,
			terminalWidth: 20,
			expected:     []string{"this is a longer", "string that needs", "to be wrapped"},
		},
		{
			name:         "Margin 4, Array Wrap",
			input:        "this is a longer string that needs to be wrapped",
			marginSize:   4,
			terminalWidth: 24,
			expected:     []string{"this is a longer", "string that needs", "to be wrapped"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WrapArray(tt.input, tt.marginSize, tt.terminalWidth)
		
			for i, _ := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("\nResult: %v\n, Expected: %v", result[i], tt.expected[i])
				}	
			}
		})
	}
}
