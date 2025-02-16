Starting DB Update...
{{$mdb := ((dbGet 2001 "mmdb").Value) }}
{{$add := cslice
      (sdict "Name" "Crystalbeard Uragaan" "Img" "/e/e5/MHGen-Crystalbeard_Uragaan_Render_001.png" "Answer" (cslice "crystalbeard uragaan"))
      (sdict "Name" "Velocidrome" "Img" "/7/71/MH4-Velocidrome_and_Velociprey_Render_001.png" "Answer" (cslice "velocidrome"))
      (sdict "Name" "Deadeye Yian Garuga" "Img" "/b/b7/MHGen-Deadeye_Yian_Garuga_Render_001.png" "Answer" (cslice "deadeye yian garuga"))
      (sdict "Name" "Yian Kut-Ku" "Img" "/c/ce/MH4-Yian_Kut-Ku_and_Konchu_Render_001.png" "Answer" (cslice "yian kut-ku" "yian kut ku" "yian kutku"))
      (sdict "Name" "Zamtrios" "Img" "/8/89/MH4-Zamtrios_Render_001.png" "Answer" (cslice "zamtrios"))
      (sdict "Name" "Thunderlord Zinogre" "Img" "/f/fe/MHGen-Thunderlord_Zinogre_Render_001.png" "Answer" (cslice "thunderlord zinogre"))
      (sdict "Name" "Amatsu" "Img" "/c/c0/MHGen-Amatsu_Render_001.png" "Answer" (cslice "amatsu"))
      (sdict "Name" "Crimson Fatalis" "Img" "/1/16/MH15th-Crimson_Fatalis_Render_001.png" "Answer" (cslice "crimson fatalis"))
      (sdict "Name" "White Fatalis" "Img" "/f/fc/MH15th-White_Fatalis_Render_001.png" "Answer" (cslice "white fatalis"))
      (sdict "Name" "Lao-Shan Lung" "Img" "/b/bb/MHGU-Lao-Shan_Lung_Render_001.png" "Answer" (cslice "lao-shan lung" "lao shan lung" "laoshan lung"))
      (sdict "Name" "Nakarkos" "Img" "/8/81/MHGen-Nakarkos_Render_003.png" "Answer" (cslice "nakarkos"))
      (sdict "Name" "Ruby Basarios" "Img" "/f/f2/MH4-Ruby_Basarios_Render_001.png" "Answer" (cslice "ruby basarios"))
      (sdict "Name" "Emerald Congalala" "Img" "/2/21/MH4-Emerald_Congalala_Render_001.png" "Answer" (cslice "emerald congalala"))
      (sdict "Name" "Plum Daimyo Hermitaur" "Img" "/2/2f/MH4U-Plum_Daimyo_Hermitaur_Render_001.png" "Answer" (cslice "plum daimyo hermitaur"))
      (sdict "Name" "Black Gravios" "Img" "/4/44/MH4-Black_Gravios_Render_001.png" "Answer" (cslice "black gravios"))
      (sdict "Name" "Great Jaggi" "Img" "/b/b4/MH4-Great_Jaggi_and_Jaggi_Render_001.png" "Answer" (cslice "great jaggi"))
      (sdict "Name" "Purple Gypceros" "Img" "/f/f6/Purple_Gypceros_MH4_Render.png" "Answer" (cslice "purple gypceros"))
      (sdict "Name" "Ash Kecha Wacha" "Img" "/c/c6/MH4U-Ash_Kecha_Wacha_Render_001.png" "Answer" (cslice "ash kecha wacha"))
      (sdict "Name" "Red Khezu" "Img" "/9/95/MH4-Red_Khezu_Render_001.png" "Answer" (cslice "red khezu"))
      (sdict "Name" "Monoblos" "Img" "/9/9c/MH4U-Monoblos_Render_001.png" "Answer" (cslice "monoblos"))
      (sdict "Name" "White Monoblos" "Img" "/4/49/MH4U-White_Monoblos_Render_001.png" "Answer" (cslice "white monoblos"))
      (sdict "Name" "Tidal Najarala" "Img" "/9/97/MH4U-Tidal_Najarala_Render_001.png" "Answer" (cslice "tidal najarala"))
      (sdict "Name" "Shrouded Nerscylla" "Img" "/b/bc/MH4U-Shrouded_Nerscylla_Render_001.png" "Answer" (cslice "shrouded nerscylla"))
      (sdict "Name" "Desert Seltas" "Img" "/d/dd/MH4U-Desert_Seltas_Render_001.png" "Answer" (cslice "desert seltas"))
      (sdict "Name" "Desert Seltas Queen" "Img" "/7/72/MH4U-Desert_Seltas_Queen_and_Desert_Seltas_Render_001.png" "Answer" (cslice "desert seltas queen"))
      (sdict "Name" "Berserk Tetsucabra" "Img" "/d/d7/MH4U-Berserk_Tetsucabra_Render_001.png" "Answer" (cslice "berserk tetsucabra"))
      (sdict "Name" "Molten Tigrex" "Img" "/9/93/MH4-Molten_Tigrex_Render_001.png" "Answer" (cslice "molten tigrex"))
      (sdict "Name" "Blue Yian Kut-Ku" "Img" "/7/78/MH4-Blue_Yian_Kut-Ku_Render_001.png" "Answer" (cslice "blue yian kut-ku"))
      (sdict "Name" "Tigerstripe Zamtrios" "Img" "/e/e9/MH4U-Tigerstripe_Zamtrios_Render_001.png" "Answer" (cslice "tigerstripe zamtrios"))
      (sdict "Name" "Dah'ren Mohran" "Img" "/0/04/MH4-Dah%27ren_Mohran_Render_001.png" "Answer" (cslice "dah'ren mohran" "dahren mohran" "dah ren mohran"))
      (sdict "Name" "Dalamadur" "Img" "/f/fd/MH4U-Dalamadur_Render_001.png" "Answer" (cslice "dalamadur"))
      (sdict "Name" "Shah Dalamadur" "Img" "/d/dd/MH4U-Shah_Dalamadur_Render_002.png" "Answer" (cslice "shah dalamadur"))
      (sdict "Name" "Gogmazios" "Img" "/3/38/MH4U-Gogmazios_Render_001.png" "Answer" (cslice "gogmazios"))
      (sdict "Name" "Oroshi Kirin" "Img" "/8/87/MH4-Oroshi_Kirin_Render_001.png" "Answer" (cslice "oroshi kirin"))
      (sdict "Name" "Glacial Agnaktor" "Img" "/b/bb/Glacial_Agnaktor_Render.png" "Answer" (cslice "glacial agnaktor"))
      (sdict "Name" "Sand Barioth" "Img" "/f/fe/3rdGen-Sand_Barioth_Render_001.png" "Answer" (cslice "sand barioth"))
      (sdict "Name" "Jade Barroth" "Img" "/9/96/IceBarroth.png" "Answer" (cslice "jade barroth"))
      (sdict "Name" "Ceadeus" "Img" "/c/c5/3rdGen-Ceadeus_Render_001.png" "Answer" (cslice "ceadeus"))
      (sdict "Name" "Goldbeard Ceadeus" "Img" "/6/69/MH3U-Goldbeard_Ceadeus_Render_001.png" "Answer" (cslice "goldbeard ceadeus"))
      (sdict "Name" "Dire Miralis" "Img" "/0/09/MH15th-Dire_Miralis_Render_001.png" "Answer" (cslice "dire miralis"))
      (sdict "Name" "Rust Duramboros" "Img" "/6/67/Dobosubspecies.png" "Answer" (cslice "rust duramboros"))
      (sdict "Name" "Gigginox" "Img" "/8/8f/3rdGen-Gigginox_Render_001.png" "Answer" (cslice "gigginox"))
      (sdict "Name" "Baleful Gigginox" "Img" "/7/7d/3rdGen-Baleful_Gigginox_Render_001.png" "Answer" (cslice "baleful gigginox"))
      (sdict "Name" "Gobul" "Img" "/7/73/3rdGen-Gobul_Render_001.png" "Answer" (cslice "gobul"))
      (sdict "Name" "Jhen Mohran" "Img" "/a/a4/3rdGen-Jhen_Mohran_Render_001.png" "Answer" (cslice "jhen mohran"))
      (sdict "Name" "Hallowed Jhen Mohran" "Img" "/7/76/MH3U-Hallowed_Jhen_Mohran_Render_001.png" "Answer" (cslice "hallowed jhen mohran"))
      (sdict "Name" "Ivory Lagiacrus" "Img" "/9/9a/MH3G-Lagiacrus_Subspecies.png" "Answer" (cslice "ivory lagiacrus"))
      (sdict "Name" "Abyssal Lagiacrus" "Img" "/c/c0/MH3U-Abyssal_Lagiacrus_Render_001.png" "Answer" (cslice "abyssal lagiacrus"))
      (sdict "Name" "Purple Ludroth" "Img" "/7/72/3rdGen-Purple_Ludroth_Render_001.png" "Answer" (cslice "purple ludroth"))
      (sdict "Name" "Green Nargacuga" "Img" "/4/4c/3rdGen-Green_Nargacuga_Render_001.png" "Answer" (cslice "green nargacuga"))
      (sdict "Name" "Green Plesioth" "Img" "/0/00/MH3U-Green_Plesioth_Render_001.png" "Answer" (cslice "green plesioth"))
      (sdict "Name" "Qurupeco" "Img" "/2/2c/3rdGen-Qurupeco_Render_001.png" "Answer" (cslice "qurupeco"))
      (sdict "Name" "Crimson Qurupeco" "Img" "/c/c0/MHP3-Crimson_Qurupeco_Render_001.png" "Answer" (cslice "crimson qurupeco"))
      (sdict "Name" "Steel Uragaan" "Img" "/3/3e/Uragaan_Subspecies_render.png" "Answer" (cslice "steel uragaan"))
      (sdict "Name" "Violet Mizutsune" "Img" "/7/78/MHRS-Violet_Mizutsune_Render_001.png" "Answer" (cslice "violet mizutsune"))
      (sdict "Name" "Flaming Espinas" "Img" "/3/3f/MHRS-Flaming_Espinas_Render_001.png" "Answer" (cslice "flaming espinas"))
      (sdict "Name" "Risen Chameleos" "Img" "/5/5c/MHRS-Risen_Chameleos_Render_001.png" "Answer" (cslice "risen chameleos"))
 }}
{{$load := $mdb.AppendSlice $add}}
{{ if not $mdb }}
	{{$load = $add}}
{{end}}
{{if $load}}
{{dbSet 2001 "mmdb" $load }}
Loaded {{len $add}} entries.
Total DB Entries: {{len $load}}
Last Entry: {{index $load (sub (len $load) 1)}}
{{else}}
Dataset is nil. No data loaded.
{{end}}
