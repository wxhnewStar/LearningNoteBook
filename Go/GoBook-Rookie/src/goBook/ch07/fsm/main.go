package main

import (
	"fmt"
)

// IdleState 闲置状态
type IdleState struct {
	StateInfo // 内嵌 StateInfo 实现基础接口
}

// OnBegin 重新实现状态开始
func (i *IdleState) OnBegin() {
	fmt.Println("IdleState begin")
}

// OnEnd 重新实现状态结束
func (i *IdleState) OnEnd() {
	fmt.Println("IdleState end")
}

// MoveState 移动状态
type MoveState struct {
	StateInfo
}

// OnBegin 重新实现状态开始
func (m *MoveState) OnBegin() {
	fmt.Println("MoveState begin")
}

// OnEnd 重新实现状态结束
func (m *MoveState) OnEnd() {
	fmt.Println("MoveSate end")
}

// EnableSameTransit 允许移动状态相互转换
func (m *MoveState) EnableSameTransit() bool {
	return true
}

// 跳跃状态
type JumpState struct {
	StateInfo
}

func (j *JumpState) OnBegin() {
	fmt.Println("JumpState begin")
}

func (j *JumpState) OnEnd() {
	fmt.Println("JumpState end")
}

/* -------- 整合代码,将自定义状态添加到状态管理器中,同时在状态改变时,打印状态转移的详细日志 --------- */

// 封装转移状态和输出日志
func transitAndReport(sm *StateManager, target string) {
	if err := sm.Transit(target); err != nil {
		fmt.Printf("FAILED! %s --> %s, %s\n\n", sm.CurState().Name(), target, err.Error())
	}
}

func main() {
	// 实例化一个状态管理器
	sm := NewStateManager()

	// 响应状态转移的通知
	sm.OnChange = func(from, to State) {
		// 打印状态转移的流向
		fmt.Printf("%s ---> %s\n\n", StateName(from), StateName(to))
	}

	// 添加 3 个状态
	sm.Add(new(IdleState))
	sm.Add(new(JumpState))
	sm.Add(new(MoveState))

	var ids string = "IdleState"
	var mvs string = "MoveState"
	var jms string = "JumpState"

	// 在不同状态间转移
	transitAndReport(sm, ids)

	transitAndReport(sm, mvs)

	transitAndReport(sm, mvs)

	transitAndReport(sm, jms)

	transitAndReport(sm, jms)

	transitAndReport(sm, ids)
	
}
