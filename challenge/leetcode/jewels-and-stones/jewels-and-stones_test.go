package jewels_and_stones

import "testing"

func Test_numJewelsInStones(t *testing.T) {
	if numJewelsInStones("aA", "aAAbbbb") != 3 {
		t.Error("invalid jewels J: aA, S: aAAbbbb")
	}

	if numJewelsInStones("z", "ZZ") != 0 {
		t.Error("invalid jewels J: z, S: ZZ")
	}
}
