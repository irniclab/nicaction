module github.com/irniclab/nicaction/domainAction

require (
	github.com/irniclab/nicaction/types v0.0.0-20230403143947-efe843c00fb5
	github.com/irniclab/nicaction/xmlRequest v0.0.0-20230403143947-efe843c00fb5
)

require (
	github.com/irniclab/nicaction/nicResponse v0.0.0-20230405120616-2ec074230def // indirect
	github.com/yaa110/go-persian-calendar v1.1.3 // indirect
)

replace github.com/irniclab/nicaction/types => ../types

replace github.com/irniclab/nicaction/xmlRequest => ../xmlRequest

go 1.19
