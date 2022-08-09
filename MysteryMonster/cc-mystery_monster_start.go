{{/* 
	Trigger Type: Command
	Trigger Pattern: -mm
*/}}

{{/* Create dictionary */}}
{{ $mDict := sdict 
    "Aknosom" (sdict "Img" "/2/23/MHRise-Aknosom_Render_001.png" "Answer" (cslice "aknosom"))
    "Almudron" (sdict "Img" "/7/71/MHRise-Almudron_Render_001.png" "Answer" (cslice "almudron"))
    "Anjanath" (sdict "Img" "/6/6b/MHRise-Anjanath_Render_001.png" "Answer" (cslice "anjanath"))
    "Apex Arzuros" (sdict "Img" "/b/b0/MHRise-Apex_Arzuros_Render_001.png" "Answer" (cslice "apex arzuros"))
    "Apex Diablos" (sdict "Img" "/7/7e/MHRise-Apex_Diablos_Render_001.png" "Answer" (cslice "apex diablos"))
    "Apex Mizutsune" (sdict "Img" "/b/bf/MHRise-Apex_Mizutsune_Render_001.png" "Answer" (cslice "apex mizutsune"))
    "Apex Rathalos" (sdict "Img" "/5/55/MHRise-Apex_Rathalos_Render_001.png" "Answer" (cslice "apex rathalos"))
    "Apex Rathian" (sdict "Img" "/7/70/MHRise-Apex_Rathian_Render_001.png" "Answer" (cslice "apex rathian"))
    "Apex Zinogre" (sdict "Img" "/b/bf/MHRise-Apex_Zinogre_Render_001.png" "Answer" (cslice "apex zinogre"))
    "Arzuros" (sdict "Img" "/e/ed/MHRise-Arzuros_Render_001.png" "Answer" (cslice "arzuros"))
    "Astalos" (sdict "Img" "/b/be/MHRS-Astalos_Render_001.png" "Answer" (cslice "astalos"))
    "Aurora Somnacanth" (sdict "Img" "/b/b6/MHRS-Aurora_Somnacanth_Render_001.png" "Answer" (cslice "aurora somnacanth"))
    "Barioth" (sdict "Img" "/2/2f/MHRise-Barioth_Render_001.png" "Answer" (cslice "barioth"))
    "Barroth" (sdict "Img" "/3/3a/MHRise-Barroth_Render_001.png" "Answer" (cslice "barroth"))
    "Basarios" (sdict "Img" "/3/36/MHRise-Basarios_Render_001.png" "Answer" (cslice "basarios"))
    "Bazelgeuse" (sdict "Img" "/e/e2/MHW-Bazelgeuse_Render_001.png" "Answer" (cslice "bazelgeuse"))
    "Bishaten" (sdict "Img" "/e/eb/MHRise-Bishaten_Render_001.png" "Answer" (cslice "bishaten"))
    "Blood Orange Bishaten" (sdict "Img" "/5/57/MHRS-Blood_Orange_Bishaten_Render_001.png" "Answer" (cslice "blood orange bishaten"))
    "Chameleos" (sdict "Img" "/0/0d/MHRise-Chameleos_Render_001.png" "Answer" (cslice "chameleos"))
    "Crimson Glow Valstrax" (sdict "Img" "/7/78/MHRise-Crimson_Glow_Valstrax_Render_001.png" "Answer" (cslice "crimson glow valstrax" "valstrax"))
    "Daimyo Hermitaur" (sdict "Img" "/c/cd/MHRS-Daimyo_Hermitaur_Render_001.png" "Answer" (cslice "daimyo hermitaur"))
    "Diablos" (sdict "Img" "/3/39/MHRise-Diablos_Render_001.png" "Answer" (cslice "diablos"))
    "Espinas" (sdict "Img" "/8/8e/MHRS-Espinas_Render_001.png" "Answer" (cslice "espinas"))
    "Furious Rajang" (sdict "Img" "/0/0e/MHRS-Furious_Rajang_Render_001.png" "Answer" (cslice "furious rajang"))
    "Gaismagorm" (sdict "Img" "/b/b0/MHRS-Gaismagorm_Icon.png" "Answer" (cslice "gaismagorm"))
    "Garangolm" (sdict "Img" "/2/2d/MHRS-Garangolm_Render_001.png" "Answer" (cslice "garangolm"))
    "Gore Magala" (sdict "Img" "/2/24/MHRS-Gore_Magala_Render_001.png" "Answer" (cslice "gore magala" "gore-magala" "goremagala"))
    "Goss Harag" (sdict "Img" "/7/77/MHRise-Goss_Harag_Render_001.png" "Answer" (cslice "goss harag"))
    "Great Baggi" (sdict "Img" "/a/a7/MHRise-Great_Baggi_Render_001.png" "Answer" (cslice "great baggi"))
    "Great Izuchi" (sdict "Img" "/8/8e/MHRise-Great_Izuchi_Render_001.png" "Answer" (cslice "great izuchi"))
    "Great Wroggi" (sdict "Img" "/9/91/MHRise-Great_Wroggi_Render_001.png" "Answer" (cslice "great wroggi"))
    "Jyuratodus" (sdict "Img" "/7/70/MHRise-Jyuratodus_Render_001.png" "Answer" (cslice "jyuratodus"))
    "Khezu" (sdict "Img" "/8/85/MHRise-Khezu_Render_001.png" "Answer" (cslice "khezu"))
    "Kulu-Ya-Ku" (sdict "Img" "/0/0e/MHRise-Kulu-Ya-Ku_Render_001.png" "Answer" (cslice "kulu-ya-ku" "kulu ya ku" "kuluyaku" "kulu yaku"))
    "Kushala Daora" (sdict "Img" "/0/07/MHRise-Kushala_Daora_Render_001.png" "Answer" (cslice "kushala daora"))
    "Lagombi" (sdict "Img" "/d/db/MHRise-Lagombi_Render_001.png" "Answer" (cslice "lagombi"))
    "Lucent Nargacuga" (sdict "Img" "/4/4f/MHRS-Lucent_Nargacuga_Render_001.png" "Answer" (cslice "lucent nargacuga"))
    "Lunagaron" (sdict "Img" "/b/b0/MHRS-Lunagaron_Render_001.png" "Answer" (cslice "lunagaron"))
    "Magma Almudron" (sdict "Img" "/9/9d/MHRS-Magma_Almudron_Render_001.png" "Answer" (cslice "magma almudron"))
    "Magnamalo" (sdict "Img" "/7/72/MHRise-Magnamalo_Render_001.png" "Answer" (cslice "magnamalo"))
    "Malzeno" (sdict "Img" "/2/20/MHRS-Malzeno_Render_001.png" "Answer" (cslice "malzeno"))
    "Mizutsune" (sdict "Img" "/e/ed/MHRise-Mizutsune_Render_001.png" "Answer" (cslice "mizutsune"))
    "Nargacuga" (sdict "Img" "/2/23/MHRise-Nargacuga_Render_001.png" "Answer" (cslice "nargacuga"))
    "Narwa the Allmother" (sdict "Img" "/0/05/MHRise-Narwa_the_Allmother_Render_001.png" "Answer" (cslice "narwa the allmother" "allmother narwa"))
    "Pukei-Pukei" (sdict "Img" "/1/13/MHRise-Pukei-Pukei_Render_001.png" "Answer" (cslice "pukei-pukei" "pukeipukei"))
    "Pyre Rakna-Kadaki" (sdict "Img" "/3/3b/MHRS-Pyre_Rakna-Kadaki_Render_001.png" "Answer" (cslice "pyre rakna-kadaki" "pyre rakna kadaki" "pyre-rakna-kadaki" "pyre raknakadaki" "pyreraknakadaki"))
    "Rajang" (sdict "Img" "/6/6d/MHRise-Rajang_Render_001.png" "Answer" (cslice "rajang"))
    "Rakna-Kadaki" (sdict "Img" "/5/5e/MHRise-Rakna-Kadaki_Render_001.png" "Answer" (cslice "rakna-kadaki" "rakna kadaki" "raknakadaki"))
    "Rathalos" (sdict "Img" "/0/00/MHRise-Rathalos_Render_001.png" "Answer" (cslice "rathalos"))
    "Rathian" (sdict "Img" "/b/be/MHRise-Rathian_Render_001.png" "Answer" (cslice "rathian"))
    "Royal Ludroth" (sdict "Img" "/c/c9/MHRise-Royal_Ludroth_Render_001.png" "Answer" (cslice "royal ludroth"))
    "Scorned Magnamalo" (sdict "Img" "/7/75/MHRS-Scorned_Magnamalo_Render_001.png" "Answer" (cslice "scorned magnamal"))
    "Seething Bazelgeuse" (sdict "Img" "/5/58/MHRS-Seething_Bazelgeuse_Render_001.png" "Answer" (cslice "seething bazelgeuse"))
    "Seregios" (sdict "Img" "/e/e9/MHRS-Seregios_Render_001.png" "Answer" (cslice "seregios"))
    "Shagaru Magala" (sdict "Img" "/c/c6/MHRS-Shagaru_Magala_Render_001.png" "Answer" (cslice "shagaru magala"))
    "Shogun Ceanatau" (sdict "Img" "/7/7b/MHRS-Shogun_Ceanataur_Render_001.png" "Answer" (cslice "shogun ceanataur"))
    "Somnacanth" (sdict "Img" "/e/e5/MHRise-Somnacanth_Render_001.png" "Answer" (cslice "somnacanth"))
    "Teostra" (sdict "Img" "/2/29/MHRise-Teostra_Render_001.png" "Answer" (cslice "teostra"))
    "Tetranadon" (sdict "Img" "/9/9e/MHRise-Tetranadon_Render_001.png" "Answer" (cslice "tetranadon"))
    "Thunder Serpent Narwa" (sdict "Img" "/e/e9/MHRise-Thunder_Serpent_Narwa_Render_001.png" "Answer" (cslice "thunder serpent narwa" "narwa"))
    "Tigrex" (sdict "Img" "/2/21/MHRise-Tigrex_Render_001.png" "Answer" (cslice "tigrex"))
    "Tobi-Kadachi" (sdict "Img" "/2/22/MHRise-Tobi-Kadachi_Render_001.png" "Answer" (cslice "tobi-kadachi" "tobi kadachi" "tobikadachi"))
    "Volvidon" (sdict "Img" "/d/d7/MHRise-Volvidon_Render_001.png" "Answer" (cslice "volvidon"))
    "Wind Serpent Ibushi" (sdict "Img" "/b/b0/MHRise-Wind_Serpent_Ibushi_Render_001.png" "Answer" (cslice "wind serpent ibushi" "ibushi"))
    "Zinogre" (sdict "Img" "/3/37/MHRise-Zinogre_Render_001.png" "Answer" (cslice "zinogre"))
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
	"-desc" "Type the name in chat"
	"-footer" "*Click image to enlarge*"
	"-color" "0000ff"
}}
{{/* Update DB with chosen monster */}}
{{dbSet .Channel.ID "MM" (sdict 
	"MonsterName" $randMon 
	"Answer" (($mDict.Get $randMon).Answer)
	"Img" $randMonImg
)}}
