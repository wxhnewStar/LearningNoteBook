package main

// 状态的基础信息和默认实现
type StateInfo struct {
	// 状态名
	name string
}

func (s *StateInfo) Name() string {
	return s.name
}

// 提供给内部设置名字
func (s *StateInfo) setName(name string) {
	s.name = name
}

// EnableSameTransit 允许同状态转移
func (s *StateInfo) EnableSameTransit() bool {
	return false
}

// 默认将状态开启时实现
func (s *StateInfo) OnBegin() {

}

func (s *StateInfo) OnEnd() {

}

func (s *StateInfo) CanTransitTo(name string) bool {
	return true
}
