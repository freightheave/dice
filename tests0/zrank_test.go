// Copyright (c) 2022-present, DiceDB contributors
// All rights reserved. Licensed under the BSD 3-Clause License. See LICENSE file in the project root for full license information.

package ironhawk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZRANK(t *testing.T) {
	client := getLocalConnection()
	defer client.Close()

	client.FireString("DEL key")
	defer client.FireString("DEL key")

	// Initialize the sorted set with members and their scores
	client.FireString("ZADD key 1 member1 2 member2 3 member3 4 member4 5 member5")

	testCases := []TestCase{
		{
			name:     "ZRANK of existing member",
			commands: []string{"ZRANK key member1"},
			expected: []interface{}{int64(0)},
		},
		{
			name:     "ZRANK of non-existing member",
			commands: []string{"ZRANK key member6"},
			expected: []interface{}{"(nil)"},
		},
		{
			name:     "ZRANK with WITHSCORE option for existing member",
			commands: []string{"ZRANK key member3 WITHSCORE"},
			expected: []interface{}{[]interface{}{int64(2), "3"}},
		},
		{
			name:     "ZRANK with WITHSCORE option for non-existing member",
			commands: []string{"ZRANK key member6 WITHSCORE"},
			expected: []interface{}{"(nil)"},
		},
		{
			name:     "ZRANK on non-existing key",
			commands: []string{"ZRANK nonexisting member1"},
			expected: []interface{}{"(nil)"},
		},
		{
			name:     "ZRANK with wrong number of arguments",
			commands: []string{"ZRANK key"},
			expected: []interface{}{"ERR wrong number of arguments for 'zrank' command"},
		},
		{
			name:     "ZRANK with invalid option",
			commands: []string{"ZRANK key member1 INVALID_OPTION"},
			expected: []interface{}{"ERR syntax error"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for i, cmd := range tc.commands {
				result := client.FireString(cmd)
				assert.Equal(t, tc.expected[i], result)
			}
		})
	}
}
