{{/* Create dictionary */}}
{{ $mDict := sdict 
    "Aknosom" (sdict "Img" "/2/23/MHRise-Aknosom_Render_001.png" "Answer" (cslice "aknosom"))
    "Almudron" (sdict "Img" "/7/71/MHRise-Almudron_Render_001.png" "Answer" (cslice "almudron"))
    "Anjanath" (sdict "Img" "/6/6b/MHRise-Anjanath_Render_001.png" "Answer" (cslice "anjanath"))
    "Arzuros" (sdict "Img" "/e/ed/MHRise-Arzuros_Render_001.png" "Answer" (cslice "arzuros"))
    "Apex Arzuros" (sdict "Img" "/b/b0/MHRise-Apex_Arzuros_Render_001.png" "Answer" (cslice "apex arzuros"))
    "Barioth" (sdict "Img" "/2/2f/MHRise-Barioth_Render_001.png" "Answer" (cslice "barioth"))
    "Barroth" (sdict "Img" "/3/3a/MHRise-Barroth_Render_001.png" "Answer" (cslice "barroth"))
    "Bazelgeuse" (sdict "Img" "/e/e2/MHW-Bazelgeuse_Render_001.png" "Answer" (cslice "bazelgeuse"))
    "Basarios" (sdict "Img" "/3/36/MHRise-Basarios_Render_001.png" "Answer" (cslice "basarios"))
    "Bishaten" (sdict "Img" "/e/eb/MHRise-Bishaten_Render_001.png" "Answer" (cslice "bishaten"))
    "Chameleos" (sdict "Img" "/0/0d/MHRise-Chameleos_Render_001.png" "Answer" (cslice "chameleos"))
    "Diablos" (sdict "Img" "/3/39/MHRise-Diablos_Render_001.png" "Answer" (cslice "diablos"))
    "Apex Diablos" (sdict "Img" "/7/7e/MHRise-Apex_Diablos_Render_001.png" "Answer" (cslice "apex diablos"))
    "Goss Harag" (sdict "Img" "/7/77/MHRise-Goss_Harag_Render_001.png" "Answer" (cslice "goss harag"))
    "Great Baggi" (sdict "Img" "/a/a7/MHRise-Great_Baggi_Render_001.png" "Answer" (cslice "great baggi"))
    "Great Izuchi" (sdict "Img" "/8/8e/MHRise-Great_Izuchi_Render_001.png" "Answer" (cslice "great izuchi"))
    "Great Wroggi" (sdict "Img" "/9/91/MHRise-Great_Wroggi_Render_001.png" "Answer" (cslice "great wroggi"))
    "Jyuratodus" (sdict "Img" "/7/70/MHRise-Jyuratodus_Render_001.png" "Answer" (cslice "jyuratodus"))
    "Khezu" (sdict "Img" "/8/85/MHRise-Khezu_Render_001.png" "Answer" (cslice "khezu"))
    "Kulu-Ya-Ku" (sdict "Img" "/0/0e/MHRise-Kulu-Ya-Ku_Render_001.png" "Answer" (cslice "kulu-ya-ku" "kulu ya ku" "kuluyaku" "kulu yaku"))
    "Kushala Daora" (sdict "Img" "/0/07/MHRise-Kushala_Daora_Render_001.png" "Answer" (cslice "kushala daora"))
    "Lagombi" (sdict "Img" "/d/db/MHRise-Lagombi_Render_001.png" "Answer" (cslice "lagombi"))
    "Magnamalo" (sdict "Img" "/7/72/MHRise-Magnamalo_Render_001.png" "Answer" (cslice "magnamalo"))
    "Mizutsune" (sdict "Img" "/e/ed/MHRise-Mizutsune_Render_001.png" "Answer" (cslice "mizutsune"))
    "Apex Mizutsune" (sdict "Img" "/b/bf/MHRise-Apex_Mizutsune_Render_001.png" "Answer" (cslice "apex mizutsune"))
    "Nargacuga" (sdict "Img" "/2/23/MHRise-Nargacuga_Render_001.png" "Answer" (cslice "nargacuga"))
    "Pukei-Pukei" (sdict "Img" "/1/13/MHRise-Pukei-Pukei_Render_001.png" "Answer" (cslice "pukei-pukei" "pukeipukei"))
    "Rajang" (sdict "Img" "/6/6d/MHRise-Rajang_Render_001.png" "Answer" (cslice "rajang"))
    "Rakna-Kadaki" (sdict "Img" "/5/5e/MHRise-Rakna-Kadaki_Render_001.png" "Answer" (cslice "rakna-kadaki" "rakna kadaki" "raknakadaki"))
    "Rathalos" (sdict "Img" "/0/00/MHRise-Rathalos_Render_001.png" "Answer" (cslice "rathalos"))
    "Apex Rathalos" (sdict "Img" "/5/55/MHRise-Apex_Rathalos_Render_001.png" "Answer" (cslice "apex rathalos"))
    "Rathian" (sdict "Img" "/b/be/MHRise-Rathian_Render_001.png" "Answer" (cslice "rathian"))
    "Apex Rathian" (sdict "Img" "/7/70/MHRise-Apex_Rathian_Render_001.png" "Answer" (cslice "apex rathian"))
    "Royal Ludroth" (sdict "Img" "/c/c9/MHRise-Royal_Ludroth_Render_001.png" "Answer" (cslice "royal ludroth"))
    "Somnacanth" (sdict "Img" "/e/e5/MHRise-Somnacanth_Render_001.png" "Answer" (cslice "somnacanth"))
    "Teostra" (sdict "Img" "/2/29/MHRise-Teostra_Render_001.png" "Answer" (cslice "teostra"))
    "Tetranadon" (sdict "Img" "/9/9e/MHRise-Tetranadon_Render_001.png" "Answer" (cslice "tetranadon"))
    "Thunder Serpent Narwa" (sdict "Img" "/e/e9/MHRise-Thunder_Serpent_Narwa_Render_001.png" "Answer" (cslice "thunder serpent narwa" "narwa"))
    "Narwa the Allmother" (sdict "Img" "/0/05/MHRise-Narwa_the_Allmother_Render_001.png" "Answer" (cslice "narwa the allmother" "allmother narwa"))
    "Tigrex" (sdict "Img" "/2/21/MHRise-Tigrex_Render_001.png" "Answer" (cslice "tigrex"))
    "Tobi-Kadachi" (sdict "Img" "/2/22/MHRise-Tobi-Kadachi_Render_001.png" "Answer" (cslice "tobi-kadachi" "tobi kadachi" "tobikadachi"))
    "Crimson Glow Valstrax" (sdict "Img" "/7/78/MHRise-Crimson_Glow_Valstrax_Render_001.png" "Answer" (cslice "crimson glow valstrax" "valstrax"))
    "Volvidon" (sdict "Img" "/d/d7/MHRise-Volvidon_Render_001.png" "Answer" (cslice "volvidon"))
    "Wind Serpent Ibushi" (sdict "Img" "/b/b0/MHRise-Wind_Serpent_Ibushi_Render_001.png" "Answer" (cslice "wind serpent ibushi" "ibushi"))
    "Zinogre" (sdict "Img" "/3/37/MHRise-Zinogre_Render_001.png" "Answer" (cslice "zinogre"))
    "Apex Zinogre" (sdict "Img" "/b/bf/MHRise-Apex_Zinogre_Render_001.png" "Answer" (cslice "apex zinogre"))
}}
{{ $mSlice := cslice "" }} {{/* Declare slice outside of range */}}
{{ range $key, $value := $mDict }}
  {{- if eq (index $mSlice 0) "" }}
    {{- $mSlice.Set 0 $key -}} {{/* Replace first entry in slice */}}
  {{- else -}}
    {{- $mSlice = $mSlice.Append $key -}} {{/* Append subsequent keys */}}
  {{- end -}}
{{- end -}}
{{/* Generate random key/value from dictionary */}}
{{ $randMon := index $mSlice (randInt (len $mSlice)) }}
{{ $randMonImg := ($mDict.Get $randMon).Img }}
{{/* Output Random Monster */}}
{{execAdmin "se"
	"-title" "***MYSTERY MONSTER***" 
	"-image" (joinStr "" "https://static.wikia.nocookie.net/monsterhunter/images" $randMonImg) 
	"-desc" "Can you identify the monster?"
	"-footer" "*Click image to enlarge*"
	"-color" "0000ff"
}}
{{/* Update DB with chosen monster */}}
{{dbSet .Channel.ID "MM" (sdict 
	"MonsterName" $randMon 
	"Answer" (($mDict.Get $randMon).Answer)
	"Img" $randMonImg
)}}
