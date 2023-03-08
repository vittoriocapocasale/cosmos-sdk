package orm

const (
	// Group Table
	GroupTablePrefix        byte = 0x0
	GroupTableSeqPrefix     byte = 0x1
	GroupByAdminIndexPrefix byte = 0x2

	// Group Member Table
	GroupMemberTablePrefix         byte = 0x10
	GroupMemberByGroupIndexPrefix  byte = 0x11
	GroupMemberByMemberIndexPrefix byte = 0x12

	// Group Policy Table
	GroupPolicyTablePrefix        byte = 0x20
	GroupPolicyTableSeqPrefix     byte = 0x21
	GroupPolicyByGroupIndexPrefix byte = 0x22
	GroupPolicyByAdminIndexPrefix byte = 0x23

	// Proposal Table
	ProposalTablePrefix              byte = 0x30
	ProposalTableSeqPrefix           byte = 0x31
	ProposalByGroupPolicyIndexPrefix byte = 0x32
	ProposalsByVotingPeriodEndPrefix byte = 0x33

	// Vote Table
	VoteTablePrefix           byte = 0x40
	VoteByProposalIndexPrefix byte = 0x41
	VoteByVoterIndexPrefix    byte = 0x42
)
