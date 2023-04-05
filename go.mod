module NicActions

go 1.19

require (
	github.com/irniclab/nicaction/config v0.0.0-20230403143947-efe843c00fb5
	github.com/irniclab/nicaction/domainAction v0.0.0-20230403143947-efe843c00fb5
	github.com/irniclab/nicaction/types v0.0.0-20230403143947-efe843c00fb5
	github.com/irniclab/nicaction/xmlRequest v0.0.0-20230403143947-efe843c00fb5 // indirect
)
 

replace github.com/irniclab/nicaction/config => ./config

replace github.com/irniclab/nicaction/domainAction => ./domainAction

replace github.com/irniclab/nicaction/types => ./types

replace github.com/irniclab/nicaction/xmlRequest => ./xmlRequest
