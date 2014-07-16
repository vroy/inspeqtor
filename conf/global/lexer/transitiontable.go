
package lexer



/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates] func(rune) int

var TransTab = TransitionTable{
	
		// S0
		func(r rune) int {
			switch {
			case r == 9 : // ['\t','\t']
				return 1
			case r == 10 : // ['\n','\n']
				return 1
			case r == 13 : // ['\r','\r']
				return 1
			case r == 32 : // [' ',' ']
				return 1
			case r == 33 : // ['!','!']
				return 2
			case r == 34 : // ['"','"']
				return 3
			case r == 35 : // ['#','#']
				return 4
			case r == 36 : // ['$','$']
				return 2
			case r == 37 : // ['%','%']
				return 2
			case r == 38 : // ['&','&']
				return 2
			case r == 39 : // [''',''']
				return 2
			case r == 42 : // ['*','*']
				return 2
			case r == 43 : // ['+','+']
				return 2
			case r == 45 : // ['-','-']
				return 2
			case r == 47 : // ['/','/']
				return 2
			case 48 <= r && r <= 57 : // ['0','9']
				return 5
			case r == 58 : // [':',':']
				return 6
			case r == 59 : // [';',';']
				return 2
			case r == 61 : // ['=','=']
				return 2
			case r == 63 : // ['?','?']
				return 2
			case r == 64 : // ['@','@']
				return 2
			case 65 <= r && r <= 90 : // ['A','Z']
				return 7
			case r == 91 : // ['[','[']
				return 8
			case r == 93 : // [']',']']
				return 9
			case r == 94 : // ['^','^']
				return 2
			case r == 95 : // ['_','_']
				return 10
			case 97 <= r && r <= 122 : // ['a','z']
				return 7
			case r == 123 : // ['{','{']
				return 2
			case r == 124 : // ['|','|']
				return 2
			case r == 125 : // ['}','}']
				return 2
			case r == 126 : // ['~','~']
				return 2
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 2
			case r == 8216 : // [\u2018,\u2018]
				return 2
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 2
			
			
			
			}
			return NoState
			
		},
	
		// S1
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S2
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 2
			case r == 35 : // ['#','#']
				return 2
			case r == 36 : // ['$','$']
				return 2
			case r == 37 : // ['%','%']
				return 2
			case r == 38 : // ['&','&']
				return 2
			case r == 39 : // [''',''']
				return 2
			case r == 42 : // ['*','*']
				return 2
			case r == 43 : // ['+','+']
				return 2
			case r == 45 : // ['-','-']
				return 2
			case r == 47 : // ['/','/']
				return 2
			case 48 <= r && r <= 57 : // ['0','9']
				return 5
			case r == 59 : // [';',';']
				return 2
			case r == 61 : // ['=','=']
				return 2
			case r == 63 : // ['?','?']
				return 2
			case r == 64 : // ['@','@']
				return 2
			case 65 <= r && r <= 90 : // ['A','Z']
				return 7
			case r == 94 : // ['^','^']
				return 2
			case r == 95 : // ['_','_']
				return 10
			case 97 <= r && r <= 122 : // ['a','z']
				return 7
			case r == 123 : // ['{','{']
				return 2
			case r == 124 : // ['|','|']
				return 2
			case r == 125 : // ['}','}']
				return 2
			case r == 126 : // ['~','~']
				return 2
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 2
			case r == 8216 : // [\u2018,\u2018]
				return 2
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 2
			
			
			
			}
			return NoState
			
		},
	
		// S3
		func(r rune) int {
			switch {
			case r == 92 : // ['\','\']
				return 11
			
			
			default:
				return 12
			}
			
		},
	
		// S4
		func(r rune) int {
			switch {
			case r == 10 : // ['\n','\n']
				return 13
			case r == 33 : // ['!','!']
				return 2
			case r == 35 : // ['#','#']
				return 2
			case r == 36 : // ['$','$']
				return 2
			case r == 37 : // ['%','%']
				return 2
			case r == 38 : // ['&','&']
				return 2
			case r == 39 : // [''',''']
				return 2
			case r == 42 : // ['*','*']
				return 2
			case r == 43 : // ['+','+']
				return 2
			case r == 45 : // ['-','-']
				return 2
			case r == 47 : // ['/','/']
				return 2
			case 48 <= r && r <= 57 : // ['0','9']
				return 5
			case r == 59 : // [';',';']
				return 2
			case r == 61 : // ['=','=']
				return 2
			case r == 63 : // ['?','?']
				return 2
			case r == 64 : // ['@','@']
				return 2
			case 65 <= r && r <= 90 : // ['A','Z']
				return 7
			case r == 94 : // ['^','^']
				return 2
			case r == 95 : // ['_','_']
				return 10
			case 97 <= r && r <= 122 : // ['a','z']
				return 7
			case r == 123 : // ['{','{']
				return 2
			case r == 124 : // ['|','|']
				return 2
			case r == 125 : // ['}','}']
				return 2
			case r == 126 : // ['~','~']
				return 2
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 2
			case r == 8216 : // [\u2018,\u2018]
				return 2
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 2
			
			
			default:
				return 14
			}
			
		},
	
		// S5
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 2
			case r == 35 : // ['#','#']
				return 2
			case r == 36 : // ['$','$']
				return 2
			case r == 37 : // ['%','%']
				return 2
			case r == 38 : // ['&','&']
				return 2
			case r == 39 : // [''',''']
				return 2
			case r == 42 : // ['*','*']
				return 2
			case r == 43 : // ['+','+']
				return 2
			case r == 45 : // ['-','-']
				return 2
			case r == 47 : // ['/','/']
				return 2
			case 48 <= r && r <= 57 : // ['0','9']
				return 5
			case r == 59 : // [';',';']
				return 2
			case r == 61 : // ['=','=']
				return 2
			case r == 63 : // ['?','?']
				return 2
			case r == 64 : // ['@','@']
				return 2
			case 65 <= r && r <= 90 : // ['A','Z']
				return 7
			case r == 94 : // ['^','^']
				return 2
			case r == 95 : // ['_','_']
				return 10
			case 97 <= r && r <= 122 : // ['a','z']
				return 7
			case r == 123 : // ['{','{']
				return 2
			case r == 124 : // ['|','|']
				return 2
			case r == 125 : // ['}','}']
				return 2
			case r == 126 : // ['~','~']
				return 2
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 2
			case r == 8216 : // [\u2018,\u2018]
				return 2
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 2
			
			
			
			}
			return NoState
			
		},
	
		// S6
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S7
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 2
			case r == 35 : // ['#','#']
				return 2
			case r == 36 : // ['$','$']
				return 2
			case r == 37 : // ['%','%']
				return 2
			case r == 38 : // ['&','&']
				return 2
			case r == 39 : // [''',''']
				return 2
			case r == 42 : // ['*','*']
				return 2
			case r == 43 : // ['+','+']
				return 2
			case r == 45 : // ['-','-']
				return 2
			case r == 47 : // ['/','/']
				return 2
			case 48 <= r && r <= 57 : // ['0','9']
				return 5
			case r == 59 : // [';',';']
				return 2
			case r == 61 : // ['=','=']
				return 2
			case r == 63 : // ['?','?']
				return 2
			case r == 64 : // ['@','@']
				return 2
			case 65 <= r && r <= 90 : // ['A','Z']
				return 7
			case r == 94 : // ['^','^']
				return 2
			case r == 95 : // ['_','_']
				return 10
			case 97 <= r && r <= 122 : // ['a','z']
				return 7
			case r == 123 : // ['{','{']
				return 2
			case r == 124 : // ['|','|']
				return 2
			case r == 125 : // ['}','}']
				return 2
			case r == 126 : // ['~','~']
				return 2
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 2
			case r == 8216 : // [\u2018,\u2018]
				return 2
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 2
			
			
			
			}
			return NoState
			
		},
	
		// S8
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S9
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			case r == 33 : // ['!','!']
				return 2
			case r == 35 : // ['#','#']
				return 2
			case r == 36 : // ['$','$']
				return 2
			case r == 37 : // ['%','%']
				return 2
			case r == 38 : // ['&','&']
				return 2
			case r == 39 : // [''',''']
				return 2
			case r == 42 : // ['*','*']
				return 2
			case r == 43 : // ['+','+']
				return 2
			case r == 45 : // ['-','-']
				return 2
			case r == 47 : // ['/','/']
				return 2
			case 48 <= r && r <= 57 : // ['0','9']
				return 5
			case r == 59 : // [';',';']
				return 2
			case r == 61 : // ['=','=']
				return 2
			case r == 63 : // ['?','?']
				return 2
			case r == 64 : // ['@','@']
				return 2
			case 65 <= r && r <= 90 : // ['A','Z']
				return 7
			case r == 94 : // ['^','^']
				return 2
			case r == 95 : // ['_','_']
				return 10
			case 97 <= r && r <= 122 : // ['a','z']
				return 7
			case r == 123 : // ['{','{']
				return 2
			case r == 124 : // ['|','|']
				return 2
			case r == 125 : // ['}','}']
				return 2
			case r == 126 : // ['~','~']
				return 2
			case 256 <= r && r <= 8215 : // [\u0100,\u2017]
				return 2
			case r == 8216 : // [\u2018,\u2018]
				return 2
			case 8217 <= r && r <= 1114111 : // [\u2019,\U0010ffff]
				return 2
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			
			
			default:
				return 15
			}
			
		},
	
		// S12
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 16
			case r == 92 : // ['\','\']
				return 17
			
			
			default:
				return 12
			}
			
		},
	
		// S13
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			case r == 10 : // ['\n','\n']
				return 13
			
			
			default:
				return 14
			}
			
		},
	
		// S15
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 16
			case r == 92 : // ['\','\']
				return 17
			
			
			default:
				return 12
			}
			
		},
	
		// S16
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			
			
			default:
				return 15
			}
			
		},
	
}
