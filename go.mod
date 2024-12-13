module github.com/Cprakhar/Pokedex

go 1.23.4

replace github.com/Cprakhar/pokeapi => ./internal/pokeapi
replace github.com/Cprakhar/pokecache => ./internal/pokecache

require (
    github.com/Cprakhar/pokeapi v0.0.0
    github.com/Cprakhar/pokecache v0.0.0
)

