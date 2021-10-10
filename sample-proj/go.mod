module sample-proj

go 1.17

replace sample-proj/add => ./add

require sample-proj/add v0.0.0-00010101000000-000000000000 // indirect
