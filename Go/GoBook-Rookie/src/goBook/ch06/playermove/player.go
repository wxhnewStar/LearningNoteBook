package main

type Player struct {
	curpos    Vec2
	targetPos Vec2
	speed     float32
}

// MoveTo 设置玩家移动的目标位置
func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}

// Pos 获取当前的位置
func (p *Player) Pos() Vec2 {
	return p.curpos
}

// IsArrived 判断是否到达目的地
func (p *Player) IsArrived() bool {
	// 计算当前玩家位置与目标位置的距离不超过移动的步长,判断已经到达目标点
	return p.curpos.DistanceTo(p.targetPos) < p.speed
}

// Update 更新玩家的位置
func (p *Player) Update() {
	if !p.IsArrived() {
		// 计算出当前位置指向目标的朝向
		dir := p.targetPos.Sub(p.curpos).Normalize()

		// 添加速度矢量生成新的位置
		newPos := p.curpos.Add(dir.Scale(p.speed))

		// 移动完成后,更新当前位置
		p.curpos = newPos
	}
}

func NewPlayer(speed float32) *Player {
	return &Player{
		speed: speed,
	}
}
