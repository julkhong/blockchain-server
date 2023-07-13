package logic

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func Test_generateWithDict(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		length   int
// 		dict     string
// 		expected string
// 	}{
// 		{
// 			"happy path - 6, letters",
// 			6,
// 			"abcdefg",
// 			"",
// 		},
// 		{
// 			"happy path - 10 - letters and numbers",
// 			10,
// 			"abcdefg12345",
// 			"",
// 		},
// 		{
// 			"happy path - 6 - one letter only",
// 			6,
// 			"a",
// 			"aaaaaa",
// 		},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := generateWithDict(tt.length, tt.dict)
// 			assert.Equal(t, len(got), tt.length)
// 			if tt.expected != "" {
// 				assert.Equal(t, tt.expected, got)
// 			}
// 		})
// 	}
// }
