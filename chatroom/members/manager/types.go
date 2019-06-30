package manager

import "gitlab.srgow.com/warehouse/proto/chatroom"

type Member struct {
	name  MemberName
	token Token
	data  map[string]string
}

func (m Member) Token() Token {
	return m.token
}

func (m Member) ToProto() *chatroom.Member {
	return &chatroom.Member{
		Name: string(m.name),
		Data: m.data,
	}
}

type Token string

type MemberName string
