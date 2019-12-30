package ospf

type lsType int

const (
	lsaTypeRouter lsType = iota + 1
	lsaTypeNetwork
	lsaTypeSummary
	lsaTypeSummaryASBR
	lsaTypeASExternal
	lsaTypeMulticast
	lsaTypeNSSA
	lsaTypeExtAttrBGP
)
