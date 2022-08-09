{{/* 
	Trigger Type: REGEX
	Trigger Pattern: .* 
*/}}

{{/* Get Message & DB Values */}}
{{ $author := "" }} {{ if .Member.Nick }}{{ $author = .Member.Nick }}{{else}}{{ $author := .User.UserName }}
{{ $Guess := .Message.Content }}
{{ $MonName := ((dbGet .Channel.ID "MM").Value).MonsterName }}
{{ $MonAnswer := ((dbGet .Channel.ID "MM").Value).Answer }}
{{ $MonImg := ((dbGet .Channel.ID "MM").Value).Img }}
{{ $IsCorrect := false }}
{{ range $i, $v := $MonAnswer }} {{/* Compare the guess with slice of possible answers */}}
	{{- if (reFind $v (lower $Guess)) -}}
		{{- $IsCorrect = true -}} {{/* Avoids multiple execution of code below */}}
	{{- end -}}
{{ end }}
{{if $IsCorrect }}
	{{execAdmin "se"
		"-author" .Member.Nick
		"-authoricon" (.User.AvatarURL "256")
		"-title" (joinStr "" "***Correct Answer! - " $MonName "***")
		"-desc" "Next round will start in `10` seconds!" 
		"-color" "00ff00"
	}}
	{{(joinStr "" "https://monsterhunter.fandom.com/wiki/" (reReplace " " $MonName "_"))}}
	{{/* Clear DB Entry to prevent spam from subsequent correct answers */}}
	{{dbSet .Channel.ID "MM" (sdict 
		"MonsterName" "" 
		"Answer" (cslice "")
		"Img" ""
	)}}
	{{execCC 19 nil 10 0}} {{/* Restarts the game */}}
{{ end }}
