{{/* 
	Trigger Type: Command
	Trigger Pattern: -mm
*/}}

{{ $mm := (dbGet 2000 (randInt (sub (dbCount 2000) 1))).Value }}

{{/* Output Random Monster */}}
{{execAdmin "se"
	"-title" "***A New MYSTERY MONSTER Appeared!!!***" 
	"-image" (joinStr "" "https://static.wikia.nocookie.net/monsterhunter/images" $mm.Img) 
	"-desc" "Type the monster's name in this chat to earn points for identifying it!"
	"-footer" "*Click image to enlarge*"
	"-color" "0000ff"
}}
{{/* Update DB with chosen monster */}}
{{dbSet .Channel.ID "MM" (sdict 
	"MonsterName" $mm.Name 
	"Answer" $mm.Answer
	"Img" $mm.Img
)}}
