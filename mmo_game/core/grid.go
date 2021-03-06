/*
AOI兴趣点  格子的模块 相关操作
*/
package core

import (
	"fmt"
	"sync"
)

type Grid struct {
	//格子ID
	GID int
	//格子的左边边界坐标
	MinX int
	//格子的右边边界坐标
	MaxX int
	//格子的上边边界坐标
	MinY int
	//格子的下面边界坐标
	MaxY int
	//当前格子内成员的ID集合  map[玩家/物体ID]
	playerIDs map[int] interface{}
	//保护当前格子内容的map的锁
	pIDLock sync.RWMutex
}
//初始化格子的方法
func NewGrid(gID,minX,maxX,minY,maxY int) *Grid  {
	return &Grid{
		GID:gID,
		MinX:minX,
		MaxX:maxX,
		MinY:minY,
		MaxY:maxY,
		playerIDs:make(map[int] interface{}),
	}
}
//给格子添加一个玩家
func (g *Grid)Add (playerID int,player interface{})  {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()

	g.playerIDs[playerID] = player
}
//从格子中删除一个玩家
func (g *Grid)Remove(playerID int)  {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()

	delete(g.playerIDs,playerID)
}
//得到当前格子所有的玩家ID
func (g *Grid)GetplayerIDs()(playerIDs []int)  {
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()
	//便利map将key封装成一个slice 返回
	for playerID,_ := range g.playerIDs{
		playerIDs = append(playerIDs,playerID)
	}
	return
}
//调试打印格子信息方法
func (g *Grid)String()string  {
	return fmt.Sprintf("Grid id = %d, minX :%d ,maxX :%d, minY :%d , maxY :%d ,playerIDs :%v\n",
		g.GID,g.MinX,g.MaxX,g.MinY,g.MaxY,g.playerIDs)
}