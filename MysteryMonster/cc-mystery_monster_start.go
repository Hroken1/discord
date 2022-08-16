{{/* 
	Trigger Type: Command
	Trigger Pattern: -mm
*/}}

{{/* Create dictionary */}}
{{ $mDict := sdict 
    "Acidic Glavenus" (sdict "Img" "/9/94/MHWI-Acidic_Glavenus_Render_001.png" "Answer" (cslice "acidic glavenus"))
    "Aknosom" (sdict "Img" "/2/23/MHRise-Aknosom_Render_001.png" "Answer" (cslice "aknosom"))
    "Alatreon" (sdict "Img" "/2/27/MHWI-Alatreon_Render_002.png" "Answer" (cslice "alatreon"))
    "Almudron" (sdict "Img" "/7/71/MHRise-Almudron_Render_001.png" "Answer" (cslice "almudron"))
    "Ancient Leshen" (sdict "Img" "/9/9f/MHW-Ancient_Leshen_Render_001.png" "Answer" (cslice "ancient leshen"))
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
    "Azure Rathalos" (sdict "Img" "/5/50/MHW-Azure_Rathalos_Render_001.png" "Answer" (cslice "azure rathalos"))
    "Banbaro" (sdict "Img" "/0/01/MHWI-Banbaro_Render_001.png" "Answer" (cslice "banbaro"))
    "Barioth" (sdict "Img" "/2/2f/MHRise-Barioth_Render_001.png" "Answer" (cslice "barioth"))
    "Barroth" (sdict "Img" "/3/3a/MHRise-Barroth_Render_001.png" "Answer" (cslice "barroth"))
    "Basarios" (sdict "Img" "/3/36/MHRise-Basarios_Render_001.png" "Answer" (cslice "basarios"))
    "Bazelgeuse" (sdict "Img" "/e/e2/MHW-Bazelgeuse_Render_001.png" "Answer" (cslice "bazelgeuse"))
    "Behemoth" (sdict "Img" "/c/cc/MHW-Behemoth_Render_001.png" "Answer" (cslice "behemoth"))
    "Beotodus" (sdict "Img" "/5/53/MHWI-Beotodus_Render_001.png" "Answer" (cslice "beotodus"))
    "Bishaten" (sdict "Img" "/e/eb/MHRise-Bishaten_Render_001.png" "Answer" (cslice "bishaten"))
    "Black Diablos" (sdict "Img" "/d/dc/MHW-Black_Diablos_Render_001.png" "Answer" (cslice "black diablos"))
    "Blackveil Vaal Hazak" (sdict "Img" "/e/e5/MHWI-Blackveil_Vaal_Hazak_Render_001.png" "Answer" (cslice "blackveil vaal hazak"))
    "Blood Orange Bishaten" (sdict "Img" "/5/57/MHRS-Blood_Orange_Bishaten_Render_001.png" "Answer" (cslice "blood orange bishaten"))
    "Brachydios" (sdict "Img" "/b/bf/MHWI-Brachydios_Render_001.png" "Answer" (cslice "brachydios"))
    "Brute Tigrex" (sdict "Img" "/2/2f/MHWI-Brute_Tigrex_Render_001.png" "Answer" (cslice "brute tigrex"))
    "Chameleos" (sdict "Img" "/0/0d/MHRise-Chameleos_Render_001.png" "Answer" (cslice "chameleos"))
    "Coral Pukei-Pukei" (sdict "Img" "/b/b4/MHWI-Coral_Pukei-Pukei_Render_001.png" "Answer" (cslice "coral pukei-pukei" "coral pukeipukei" "coral pukei pukei"))
    "Crimson Glow Valstrax" (sdict "Img" "/7/78/MHRise-Crimson_Glow_Valstrax_Render_001.png" "Answer" (cslice "crimson glow valstrax" "valstrax"))
    "Daimyo Hermitaur" (sdict "Img" "/c/cd/MHRS-Daimyo_Hermitaur_Render_001.png" "Answer" (cslice "daimyo hermitaur"))
    "Deviljho" (sdict "Img" "/c/cb/MHW-Deviljho_Render_001.png" "Answer" (cslice "deviljho"))
    "Diablos" (sdict "Img" "/3/39/MHRise-Diablos_Render_001.png" "Answer" (cslice "diablos"))
    "Dodogama" (sdict "Img" "/d/d3/MHW-Dodogama_Render_001.png" "Answer" (cslice "dodogama"))
    "Ebony Odogaron" (sdict "Img" "/8/87/MHWI-Ebony_Odogaron_Render_001.png" "Answer" (cslice "ebony odogaron"))
    "Espinas" (sdict "Img" "/8/8e/MHRS-Espinas_Render_001.png" "Answer" (cslice "espinas"))
    "Fatalis" (sdict "Img" "/1/1c/MHWI-Fatalis_Render_001.png" "Answer" (cslice "fatalis"))
    "Frostfang Barioth" (sdict "Img" "/2/23/MHWI-Frostfang_Barioth_Render_001.png" "Answer" (cslice "frostfang barioth"))
    "Fulgur Anjanath" (sdict "Img" "/d/d7/MHWI-Fulgur_Anjanath_Render_001.png" "Answer" (cslice "fulgur anjanath"))
    "Furious Rajang" (sdict "Img" "/0/0e/MHRS-Furious_Rajang_Render_001.png" "Answer" (cslice "furious rajang"))
    "Gaismagorm" (sdict "Img" "/b/b0/MHRS-Gaismagorm_Icon.png" "Answer" (cslice "gaismagorm"))
    "Garangolm" (sdict "Img" "/2/2d/MHRS-Garangolm_Render_001.png" "Answer" (cslice "garangolm"))
    "Glavenus" (sdict "Img" "/5/5c/MHWI-Glavenus_Render_001.png" "Answer" (cslice "glavenus"))
    "Gold Rathian" (sdict "Img" "/c/c8/MHRS-Gold_Rathian_Render_001.png" "Answer" (cslice "gold rathian"))
    "Gore Magala" (sdict "Img" "/2/24/MHRS-Gore_Magala_Render_001.png" "Answer" (cslice "gore magala" "gore-magala" "goremagala"))
    "Goss Harag" (sdict "Img" "/7/77/MHRise-Goss_Harag_Render_001.png" "Answer" (cslice "goss harag"))
    "Great Baggi" (sdict "Img" "/a/a7/MHRise-Great_Baggi_Render_001.png" "Answer" (cslice "great baggi"))
    "Great Girros" (sdict "Img" "/4/4c/MHW-Great_Girros_Render_001.png" "Answer" (cslice "great girros"))
    "Great Izuchi" (sdict "Img" "/8/8e/MHRise-Great_Izuchi_Render_001.png" "Answer" (cslice "great izuchi"))
    "Great Jagras" (sdict "Img" "/f/f5/MHW-Great_Jagras_Render_001.png" "Answer" (cslice "great jagras"))
    "Great Wroggi" (sdict "Img" "/9/91/MHRise-Great_Wroggi_Render_001.png" "Answer" (cslice "great wroggi"))
    "Jyuratodus" (sdict "Img" "/7/70/MHRise-Jyuratodus_Render_001.png" "Answer" (cslice "jyuratodus"))
    "Khezu" (sdict "Img" "/8/85/MHRise-Khezu_Render_001.png" "Answer" (cslice "khezu"))
    "Kirin" (sdict "Img" "/f/f5/MHW-Kirin_Render_001.png" "Answer" (cslice "kirin"))
    "Kulu-Ya-Ku" (sdict "Img" "/0/0e/MHRise-Kulu-Ya-Ku_Render_001.png" "Answer" (cslice "kulu-ya-ku" "kulu ya ku" "kuluyaku" "kulu yaku"))
    "Kulve Taroth" (sdict "Img" "/0/00/MHW-Kulve_Taroth_Render_002.png" "Answer" (cslice "kulve taroth"))
    "Kushala Daora" (sdict "Img" "/0/07/MHRise-Kushala_Daora_Render_001.png" "Answer" (cslice "kushala daora"))
    "Lagombi" (sdict "Img" "/d/db/MHRise-Lagombi_Render_001.png" "Answer" (cslice "lagombi"))
    "Lavasioth" (sdict "Img" "/8/84/MHW-Lavasioth_Render_001.png" "Answer" (cslice "lavasioth"))
    "Legiana" (sdict "Img" "/3/34/MHW-Legiana_Render_001.png" "Answer" (cslice "legiana"))
    "Leshen" (sdict "Img" "/0/07/MHW-Leshen_Render_001.png" "Answer" (cslice "leshen"))
    "Lucent Nargacuga" (sdict "Img" "/4/4f/MHRS-Lucent_Nargacuga_Render_001.png" "Answer" (cslice "lucent nargacuga"))
    "Lunagaron" (sdict "Img" "/b/b0/MHRS-Lunagaron_Render_001.png" "Answer" (cslice "lunagaron"))
    "Lunastra" (sdict "Img" "/f/f8/MHW-Lunastra_Render_001.png" "Answer" (cslice "lunastra"))
    "Magma Almudron" (sdict "Img" "/9/9d/MHRS-Magma_Almudron_Render_001.png" "Answer" (cslice "magma almudron"))
    "Magnamalo" (sdict "Img" "/7/72/MHRise-Magnamalo_Render_001.png" "Answer" (cslice "magnamalo"))
    "Malzeno" (sdict "Img" "/2/20/MHRS-Malzeno_Render_001.png" "Answer" (cslice "malzeno"))
    "Mizutsune" (sdict "Img" "/e/ed/MHRise-Mizutsune_Render_001.png" "Answer" (cslice "mizutsune"))
    "Namielle" (sdict "Img" "/d/d6/MHWI-Namielle_Render_001.png" "Answer" (cslice "namielle"))
    "Nargacuga" (sdict "Img" "/2/23/MHRise-Nargacuga_Render_001.png" "Answer" (cslice "nargacuga"))
    "Narwa the Allmother" (sdict "Img" "/0/05/MHRise-Narwa_the_Allmother_Render_001.png" "Answer" (cslice "narwa the allmother" "allmother narwa"))
    "Nergigante" (sdict "Img" "/8/89/MHW-Nergigante_Render_001.png" "Answer" (cslice "nergigante"))
    "Nightshade Paolumu" (sdict "Img" "/2/27/MHWI-Nightshade_Paolumu_Render_001.png" "Answer" (cslice "nightshade paolumu"))
    "Odogaron" (sdict "Img" "/9/9f/MHW-Odogaron_Render_001.png" "Answer" (cslice "odogaron"))
    "Paolumu" (sdict "Img" "/b/b2/MHW-Paolumu_Render_001.png" "Answer" (cslice "paolumu"))
    "Pink Rathian" (sdict "Img" "/1/14/MHW-Pink_Rathian_Render_001.png" "Answer" (cslice "pink rathian"))
    "Pukei-Pukei" (sdict "Img" "/1/13/MHRise-Pukei-Pukei_Render_001.png" "Answer" (cslice "pukei-pukei" "pukeipukei" "pukei pukei"))
    "Pyre Rakna-Kadaki" (sdict "Img" "/3/3b/MHRS-Pyre_Rakna-Kadaki_Render_001.png" "Answer" (cslice "pyre rakna-kadaki" "pyre rakna kadaki" "pyre-rakna-kadaki" "pyre raknakadaki" "pyreraknakadaki"))
    "Radobaan" (sdict "Img" "/d/d3/MHW-Radobaan_Render_001.png" "Answer" (cslice "radobaan"))
    "Raging Brachydios" (sdict "Img" "/4/41/MHWI-Raging_Brachydios_Render_001.png" "Answer" (cslice "raging brachydios"))
    "Rajang" (sdict "Img" "/6/6d/MHRise-Rajang_Render_001.png" "Answer" (cslice "rajang"))
    "Rakna-Kadaki" (sdict "Img" "/5/5e/MHRise-Rakna-Kadaki_Render_001.png" "Answer" (cslice "rakna-kadaki" "rakna kadaki" "raknakadaki"))
    "Rathalos" (sdict "Img" "/0/00/MHRise-Rathalos_Render_001.png" "Answer" (cslice "rathalos"))
    "Rathian" (sdict "Img" "/b/be/MHRise-Rathian_Render_001.png" "Answer" (cslice "rathian"))
    "Royal Ludroth" (sdict "Img" "/c/c9/MHRise-Royal_Ludroth_Render_001.png" "Answer" (cslice "royal ludroth"))
    "Ruiner Nergigante" (sdict "Img" "/4/47/MHWI-Ruiner_Nergigante_Render_001.png" "Answer" (cslice "ruiner nergigante"))
    "Safi'jiiva" (sdict "Img" "/9/92/MHWI-Safi%27jiiva_Render_001.png" "Answer" (cslice "safi'jiiva" "safijiiva"))
    "Savage Deviljho" (sdict "Img" "/f/ff/MHWI-Savage_Deviljho_Render_001.png" "Answer" (cslice "savage deviljho"))
    "Scarred Yian Garuga" (sdict "Img" "/f/f0/MHWI-Scarred_Yian_Garuga_Render_001.png" "Answer" (cslice "scarred yian garuga"))
    "Scorned Magnamalo" (sdict "Img" "/7/75/MHRS-Scorned_Magnamalo_Render_001.png" "Answer" (cslice "scorned magnamal"))
    "Seething Bazelgeuse" (sdict "Img" "/5/58/MHRS-Seething_Bazelgeuse_Render_001.png" "Answer" (cslice "seething bazelgeuse"))
    "Seregios" (sdict "Img" "/e/e9/MHRS-Seregios_Render_001.png" "Answer" (cslice "seregios"))
    "Shagaru Magala" (sdict "Img" "/c/c6/MHRS-Shagaru_Magala_Render_001.png" "Answer" (cslice "shagaru magala"))
    "Shara Ishvalda" (sdict "Img" "/0/0f/MHWI-Shara_Ishvalda_Render_002.png" "Answer" (cslice "shara ishvalda"))
    "Shogun Ceanatau" (sdict "Img" "/7/7b/MHRS-Shogun_Ceanataur_Render_001.png" "Answer" (cslice "shogun ceanataur"))
    "Shrieking Legiana" (sdict "Img" "/2/2d/MHWI-Shrieking_Legiana_Render_001.png" "Answer" (cslice "shrieking legiana"))
    "Silver Rathalos" (sdict "Img" "/7/75/MHRS-Silver_Rathalos_Render_001.png" "Answer" (cslice "silver rathalos"))
    "Somnacanth" (sdict "Img" "/e/e5/MHRise-Somnacanth_Render_001.png" "Answer" (cslice "somnacanth"))
    "Stygian Zinogre" (sdict "Img" "/e/ee/MHWI-Stygian_Zinogre_Render_001.png" "Answer" (cslice "stygian zinogre"))
    "Teostra" (sdict "Img" "/2/29/MHRise-Teostra_Render_001.png" "Answer" (cslice "teostra"))
    "Tetranadon" (sdict "Img" "/9/9e/MHRise-Tetranadon_Render_001.png" "Answer" (cslice "tetranadon"))
    "Thunder Serpent Narwa" (sdict "Img" "/e/e9/MHRise-Thunder_Serpent_Narwa_Render_001.png" "Answer" (cslice "thunder serpent narwa" "narwa"))
    "Tigrex" (sdict "Img" "/2/21/MHRise-Tigrex_Render_001.png" "Answer" (cslice "tigrex"))
    "Tobi-Kadachi" (sdict "Img" "/2/22/MHRise-Tobi-Kadachi_Render_001.png" "Answer" (cslice "tobi-kadachi" "tobi kadachi" "tobikadachi"))
    "Tzitzi-Ya-Ku" (sdict "Img" "/0/04/MHW-Tzitzi-Ya-Ku_Render_001.png" "Answer" (cslice "tzitzi-ya-ku" "tzitziyaku" "tzitzi ya ku" "tzitzi-yaku" "tzitzi yaku"))
    "Uragaan" (sdict "Img" "/5/52/MHW-Uragaan_Render_001.png" "Answer" (cslice "uragaan"))
    "Vaal Hazak" (sdict "Img" "/5/54/MHW-Vaal_Hazak_Render_001.png" "Answer" (cslice "vaal hazak"))
    "Velkhana" (sdict "Img" "/d/d5/MHWI-Velkhana_Render_001.png" "Answer" (cslice "velkhana"))
    "Viper Tobi-Kadachi" (sdict "Img" "/0/0f/MHWI-Viper_Tobi-Kadachi_Render_001.png" "Answer" (cslice "viper tobi-kadachi" "viper tobikadachi" "viper tobi kadachi"))
    "Volvidon" (sdict "Img" "/d/d7/MHRise-Volvidon_Render_001.png" "Answer" (cslice "volvidon"))
    "Wind Serpent Ibushi" (sdict "Img" "/b/b0/MHRise-Wind_Serpent_Ibushi_Render_001.png" "Answer" (cslice "wind serpent ibushi" "ibushi"))
    "Xeno'jiiva" (sdict "Img" "/2/2f/MHW-Xeno%27jiiva_Render_001.png" "Answer" (cslice "xeno'jiiva" "xenojiiva"))
    "Yian Garuga" (sdict "Img" "/7/73/MHWI-Yian_Garuga_Render_001.png" "Answer" (cslice "yian garuga"))
    "Zinogre" (sdict "Img" "/3/37/MHRise-Zinogre_Render_001.png" "Answer" (cslice "zinogre"))
    "Zorah Magdaros" (sdict "Img" "/7/73/MHW-Zorah_Magdaros_Render_001.png" "Answer" (cslice "zorah magdaros"))
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
