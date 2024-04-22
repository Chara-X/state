module github.com/Chara-X/state

go 1.22.2

replace github.com/Chara-X/collections => ../collections

replace github.com/Chara-X/slices => ../slices

require (
	github.com/Chara-X/collections v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.6.0
)

require github.com/Chara-X/slices v0.0.0-00010101000000-000000000000 // indirect
