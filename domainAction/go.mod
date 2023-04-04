module github.com/irniclab/nicaction/domainAction

require (
	github.com/irniclab/nicaction/xmlRequest v0.0.0-20230403143947-efe843c00fb5
	github.com/irniclab/nicaction/types v0.0.0-20230403143947-efe843c00fb5
)

replace github.com/irniclab/nicaction/config => ./config

replace github.com/irniclab/nicaction/xmlRequest => ./xmlRequest


go 1.19
