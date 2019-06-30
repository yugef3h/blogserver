package manager

import (
	"fmt"
	"gitlab.srgow.com/warehouse/common/utils"
	"sync"
)

func Create() *Manager {
	return &Manager{
		lock:         sync.RWMutex{},
		tokenMembers: make(map[Token]*Member),
		nameMembers:  make(map[MemberName]*Member),
	}
}

type Manager struct {
	lock         sync.RWMutex
	tokenMembers map[Token]*Member
	nameMembers  map[MemberName]*Member
}

func (m *Manager) Login(name MemberName, data map[string]string) (*Member, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	token := Token(utils.EmptyString)
	for utils.EmptyOrWhiteSpace(string(token)) {
		t, err := utils.RandomID()
		if nil != err {
			return nil, fmt.Errorf("生成Token失败: %s", err.Error())
		}
		if _, ok := m.tokenMembers[Token(t)]; !ok {
			token = Token(t)
		}
	}

	member, exists := m.nameMembers[name]
	oldToken := Token(utils.EmptyString)
	if exists {
		oldToken = member.token
		member.token = token
		member.data = data
	} else {
		member = &Member{
			name:  name,
			token: token,
			data:  data,
		}
		m.nameMembers[name] = member
	}

	m.tokenMembers[token] = member
	if !utils.EmptyOrWhiteSpace(oldToken) {
		delete(m.tokenMembers, oldToken)
	}

	return member, nil
}

func (m *Manager) Logout(token Token) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	member, ok := m.tokenMembers[token]
	if ok {
		delete(m.tokenMembers, token)
		delete(m.nameMembers, member.name)
	}

	return nil
}

func (m *Manager) Validate(token Token) (*Member, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	member, ok := m.tokenMembers[token]
	if !ok {
		return nil, fmt.Errorf("非法Token")
	}

	return member, nil
}

func (m *Manager) GetMember(name MemberName) (*Member, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	member := m.nameMembers[name]
	if nil == member {
		return nil, fmt.Errorf("未找到用户: %s", name)
	}
	return member, nil
}

func (m *Manager) SetData(token Token, data map[string]string) error {
	m.lock.RLock()
	defer m.lock.RUnlock()

	member, ok := m.tokenMembers[token]
	if !ok {
		return fmt.Errorf("非法的Token")
	}

	member.data = data
	return nil
}
