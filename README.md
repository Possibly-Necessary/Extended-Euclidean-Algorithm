# Extended-Euclidean-Algorithm

Both Python and Go scripts are written following the Algorithm from: The MoonMath Manual to ZK-SNARKS, page 12.
Book: https://leastauthority.com/community-matters/moonmath-manual/

	Algorithm: Extended Euclidean Algorithm
	Require: a, b ∈ N with a ≥ b
		procedure EXT-EUCLID(a, b)
			r0 ← a and r1 ← b
			s0 ← 1 and s1 ← 0
			t0 ← 0 and t1 ← 1
			k ← 2
			while rk−1̸= 0 do
				qk ← rk−2 div rk−1
				rk ← rk−2 mod rk−1
				sk ← sk−2 − qk · sk−1
				tk ← tk−2 − qk · tk−1
				k ← k + 1
			end while
			return gcd(a, b) ← rk−2, s ← sk−2 and t ← tk−2
		end procedure
	Ensure: gcd(a, b) = s · a + t · b
