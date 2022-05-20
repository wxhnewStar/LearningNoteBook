package main

import "errors"

// StateManager 状态管理器
type StateManager struct {
	// 已经添加的状态
	stateByName map[string]State

	// 状态改变时的回调
	OnChange func(from, to State)

	// 当前状态
	cur State
}

// Add 添加一个状态到管理器中
func (sm *StateManager) Add(s State) {
	// 获取状态的名称
	name := StateName(s)

	// 将 s 转换为能设置名字的接口,然后调用该接口
	s.(interface {
		setName(name string)
	}).setName(name) // ?? 名字不是已经设置好了吗,为什么要弄这个啊 不理解,而且这样子是不是只有 StateInfo 可以了,因为只有它实现了 setName

	// 根据状态名称获取已经添加的状态,检查该状态是否存在
	if sm.Get(name) != nil {
		panic("duplicate state:" + name)
	}

	// 根据名字保存到 map 中
	sm.stateByName[name] = s
}

// Get 根据名字获取指定状态
func (sm *StateManager) Get(name string) State {
	if v, ok := sm.stateByName[name]; ok {
		return v
	}
	return nil
}

// NewStateManager 初始化状态管理器
func NewStateManager() *StateManager {
	return &StateManager{
		stateByName: make(map[string]State),
		OnChange:    nil,
		cur:         nil,
	}
}

/*------------- 状态转移实现 ------------------*/

// ErrStateNotFound 状态没有找到的错误
var ErrStateNotFound = errors.New("state not found")

// ErrForbidSameStateTransit 禁止在同状态间转移的错误
var ErrForbidSameStateTransit = errors.New("forbid same state transmit")

// ErrCannotTransitToState 不能转移到指定状态的错误
var ErrCannotTransitToState = errors.New("cannot transit to state")

// CurState 获取当前状态
func (sm *StateManager) CurState() State {
	return sm.cur
}

// CanCurTransitTo 当前状态能够转移到目标状态
func (sm *StateManager) CanCurTransitTo(name string) bool {
	if sm.cur == nil {
		return true
	}

	// 相同的状态,且不允许同状态转换,返回 false
	if sm.cur.Name() == name && !sm.cur.EnableSameTransit() {
		return false
	}

	// 使用当前状态,检查能否转移到指定名字的状态
	return sm.cur.CanTransitTo(name)
}

// Transit 转移到指定状态
func (sm *StateManager) Transit(name string) error {
	// 获取目标状态
	next := sm.Get(name)

	// 目标不存在
	if next == nil {
		return ErrStateNotFound
	}

	// 调用旧状态的结束
	if StateName(sm.cur) != "None" {
		sm.cur.OnEnd()
	}
	// 记录转移前的状态
	pre := sm.cur

	// 当前有状态
	if sm.cur != nil {

		// 相同的状态不用转移
		if sm.cur.Name() == name && !sm.cur.EnableSameTransit() {
			return ErrForbidSameStateTransit
		}

		// 不能转移到目标状态
		if !sm.cur.CanTransitTo(name) {
			return ErrCannotTransitToState
		}
	}

	//将当前状态切换为要转移到的目标状态
	sm.cur = next

	// 调用新状态的开始
	sm.cur.OnBegin() // 不调用前状态的 onEnd() 吗?

	// 通知回调
	if sm.OnChange != nil {
		sm.OnChange(pre, sm.cur)
	}

	return nil
}
