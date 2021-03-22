package services

const TemplateIdMeetingParticipants = "h1-e7A4a5gR0gzn1VF42RnpKtFEm5oZ4S2Nx2f6BM9s"

type BindMeetingParticipants struct {
	First    string //标题
	Keyword1 string //会议室名称
	Keyword2 string //会议室地点
	Keyword3 string //预约时间
	Remark   string //备注
}
