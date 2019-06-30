package manager

import (
	"fmt"
	"gitlab.srgow.com/warehouse/common/utils"
	"sync"
)

func Create() *Manager {
	return &Manager{
		lock:        sync.RWMutex{},
		idRooms:     make(map[RoomID]*Room),
		memberRooms: make(map[MemberName]map[RoomID]bool),
	}
}

type Manager struct {
	lock        sync.RWMutex
	idRooms     map[RoomID]*Room
	memberRooms map[MemberName]map[RoomID]bool
}

func (m *Manager) Create(memberNames []MemberName, data map[string]string) (*Room, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if utils.EmptyArray(memberNames) {
		return nil, fmt.Errorf("聊天室成员为空")
	}

	id, err := utils.RandomID()
	if nil != err {
		return nil, fmt.Errorf("生成房间ID失败: %s", err.Error())
	}
	room := &Room{
		id:      RoomID(id),
		members: make(map[MemberName]bool),
		data:    data,
	}
	for _, name := range memberNames {
		room.members[name] = true
		memberRoom, ok := m.memberRooms[name]
		if !ok {
			m.memberRooms[name] = map[RoomID]bool{room.id: true}
		} else {
			memberRoom[room.id] = true
		}
		m.idRooms[room.id] = room
	}

	return room, nil
}

func (m *Manager) Available(memberName MemberName) ([]*Room, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	available, ok := m.memberRooms[memberName]
	if !ok {
		return EmptyRooms, nil
	}
	rooms := make([]*Room, 0, len(available))
	for roomID := range available {
		rooms = append(rooms, m.idRooms[roomID])
	}
	return rooms, nil
}

func (m *Manager) GetRoomByID(id RoomID) (*Room, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.idRooms[id], nil
}

func (m *Manager) DeleteRoom(id RoomID) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	room, ok := m.idRooms[id]
	if ok {
		for memberName := range room.members {
			delete(m.memberRooms[memberName], id)
		}
		delete(m.idRooms, id)
	}

	return nil
}
