module goremedy

go 1.22.1

// replace github.cerner.com/OHAIFedAutoSre/gorapid => gorapid
// require (
//     gorapid v0.0.0
// )

// replace github.cerner.com/OHAIFedAutoSre/gorapid => ./gorapid
// require (
//     ./gorapid v0.0.0
// )

require github.cerner.com/OHAIFedAutoSre/gorapid v0.0.0

replace github.cerner.com/OHAIFedAutoSre/gorapid => ../gorapid
