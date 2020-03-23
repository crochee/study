package arithmetic

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

/*
* @
* @Author:
* @Date: 2020/3/22 14:14
 */
// ---------------------------------------定义F=G+H-----------------------------------------
// 定义地图点
type Point struct {
	x    int
	y    int
	view string // 表示障碍物等信息
}

// A星点
type AstarPoint struct {
	Point // 当前节点
	g int
	h int
	f int
}

func NewAstarPoint(point, end *Point, father *AstarPoint) *AstarPoint {
	asp := &AstarPoint{
		Point: *point,
		g:     0,
		h:     0,
		f:     0,
	}
	if end != nil && father != nil {
		asp.F(father, end) // 计算g,h,f
	}
	return asp
}

// 综合评估G,H,如从起点到当前节点，再从当前节点到目标节点的总步数
func (asp *AstarPoint) F(father *AstarPoint, end *Point) int {
	asp.f = asp.G(father) + asp.H(end)
	return asp.f
}

// 计算从起点（即父节点）到当前的成本，即步数
func (asp *AstarPoint) G(father *AstarPoint) int {
	// 这里是没有斜着走的情况，实际需要考虑产品规则
	asp.g = father.g + int(math.Abs(float64(father.x-asp.x))+math.Abs(float64(father.y-asp.y)))
	return asp.g
}

// 不考虑障碍的情况下，从当前到目标的距离
func (asp *AstarPoint) H(end *Point) int {
	asp.h = int(math.Abs(float64(end.x-asp.x)) + math.Abs(float64(end.y-asp.y)))
	return asp.h
}

// ---------------------------------------openlist-----------------------------------------
// 可到达的格子
type OpenList struct {
	listMap map[string]*AstarPoint
}

func NewOpenList() *OpenList {
	return &OpenList{listMap: make(map[string]*AstarPoint)}
}

func (ol *OpenList) Has(p *AstarPoint) bool {
	_, ok := ol.listMap[pointAsKey(p.x, p.y)]
	return ok
}

func (ol *OpenList) Push(p *AstarPoint) {
	ol.listMap[pointAsKey(p.x, p.y)] = p
}

func (ol *OpenList) Pop(p *AstarPoint) {
	delete(ol.listMap, pointAsKey(p.x, p.y))
}

func (ol *OpenList) Len() int {
	return len(ol.listMap)
}

// ---------------------------------------closelist-----------------------------------------
// 已到达的格子
type CloseList struct {
	list []*AstarPoint
}

func NewCloseList() *CloseList {
	return &CloseList{list: make([]*AstarPoint, 0)}
}

func (cl *CloseList) Has(p *AstarPoint) bool {
	for _, v := range cl.list {
		if pointAsKey(v.x, v.y) == pointAsKey(p.x, p.y) {
			return true
		}
	}
	return false
}

func (cl *CloseList) Push(p *AstarPoint) {
	cl.list = append(cl.list, p)
}

func (cl *CloseList) Pop() {
	cl.list = cl.list[:len(cl.list)-1]
}

func (cl *CloseList) Tail() *AstarPoint {
	if len(cl.list) < 1 {
		return nil
	}
	return cl.list[len(cl.list)-1]
}

// --------------------------------地图信息----------------------
type ThisMap struct {
	points [][]string
	blocks map[string]uint8 // (x,y)->物体    用于确定x,y处的物体是啥  需要定义规则，如1是障碍物等
	maxX   int
	maxY   int
}

// 根据字符地图生成thisMap的结构
func NewThisMap(charMap []string) *ThisMap {
	m := &ThisMap{}
	m.points = make([][]string, len(charMap))
	m.blocks = make(map[string]uint8, len(charMap)*2)
	for x, row := range charMap {
		cols := strings.Split(row, " ")
		m.points[x] = make([]string, len(cols))
		for y, view := range cols { // 规则定义，明确各符号表示的物体
			m.points[x][y] = view
			if view == "X" { // X->表示障碍物
				m.blocks[pointAsKey(x, y)] = 1
			}
		} // end of cols
	} // end of row
	m.maxX = len(m.points)
	m.maxY = len(m.points[0])
	return m
}

func pointAsKey(x, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}

// --------------------------------A星算法对象----------------------
type SearchRoad struct {
	theMap  *ThisMap
	start   *Point
	end     *Point
	closeLi *CloseList
	openLi  *OpenList
}

func NewSearchRoad(startx, starty, endx, endy int, m *ThisMap) *SearchRoad {
	sr := &SearchRoad{}
	sr.theMap = m
	sr.start = &Point{
		x:    startx,
		y:    starty,
		view: "S",
	}
	sr.end = &Point{
		x:    endx,
		y:    endy,
		view: "E",
	}
	sr.closeLi = NewCloseList()
	sr.openLi = NewOpenList()
	return sr
}

func (sr *SearchRoad) FindoutRoad() error {
	// 1.判断起点合法性
	if !(sr.start.x <= sr.theMap.maxX && sr.start.x >= 0 && sr.start.y <= sr.theMap.maxY && sr.start.y >= 0) {
		return errors.New("起点位于地图之外！")
	} else if sr.theMap.blocks[pointAsKey(sr.start.x, sr.start.y)] == 1 {
		return errors.New("起点位于障碍物！")
	}
	// 1.直接将起点放入close
	sartAsp := NewAstarPoint(sr.start, nil, nil)
	sr.closeLi.Push(sartAsp)
	var (
		startx = sr.start.x
		starty = sr.start.y
	)
	// 2.遍历地图
	for {
		if startx == sr.end.x && starty == sr.end.y {
			break
		}
		var (
			min  = math.MaxInt64 // 记录f的最小值
			temp *AstarPoint     // 记录取得最小值的point
		)
		// 下走
		if x, y := startx+1, starty; x < sr.theMap.maxX && x >= 0 { // 不越界的情况
			if v, ok := sr.theMap.blocks[pointAsKey(x, y)]; !ok || v != 1 { // 不碰到障碍物
				asp := NewAstarPoint(&Point{
					x:    x,
					y:    y,
					view: "",
				}, sr.end, sr.closeLi.Tail())
				if !sr.openLi.Has(asp) && !sr.closeLi.Has(asp) { // 考虑的点不能是之前考虑过的
					sr.openLi.Push(asp) // 加入可达的集合
					if asp.f < min { // 记录最小的可达点
						min = asp.f
						temp = asp
					}
				}
			}
		}
		// 上走
		if x, y := startx-1, starty; x < sr.theMap.maxX && x >= 0 {
			if v, ok := sr.theMap.blocks[pointAsKey(x, y)]; !ok || v != 1 {
				asp := NewAstarPoint(&Point{
					x:    x,
					y:    y,
					view: "",
				}, sr.end, sr.closeLi.Tail())
				if !sr.openLi.Has(asp) && !sr.closeLi.Has(asp) {
					sr.openLi.Push(asp)
					if asp.f < min {
						min = asp.f
						temp = asp
					}
				}
			}
		}
		// 左走
		if x, y := startx, starty-1; y < sr.theMap.maxY && y >= 0 {
			if v, ok := sr.theMap.blocks[pointAsKey(x, y)]; !ok || v != 1 {
				asp := NewAstarPoint(&Point{
					x:    x,
					y:    y,
					view: "",
				}, sr.end, sr.closeLi.Tail())
				if !sr.openLi.Has(asp) && !sr.closeLi.Has(asp) {
					sr.openLi.Push(asp)
					if asp.f < min {
						min = asp.f
						temp = asp
					}
				}
			}
		}
		// 右走
		if x, y := startx, starty+1; y < sr.theMap.maxY && y >= 0 {
			if v, ok := sr.theMap.blocks[pointAsKey(x, y)]; !ok || v != 1 {
				asp := NewAstarPoint(&Point{
					x:    x,
					y:    y,
					view: "",
				}, sr.end, sr.closeLi.Tail())
				if !sr.openLi.Has(asp) && !sr.closeLi.Has(asp) {
					sr.openLi.Push(asp)
					if asp.f < min {
						min = asp.f
						temp = asp
					}
				}
			}
		}
		// 计算最小值的点从open出，加入close
		if temp == nil {
			return errors.New("此处死路")
		}
		sr.openLi.Pop(temp)
		sr.closeLi.Push(temp)
		startx = temp.x
		starty = temp.y
	}
	// TODO 存在的问题是，当发现是死路的时候无法回溯，还需要回溯节点
	return nil
}

func (sr *SearchRoad) Print() {
	for _, v := range sr.closeLi.list {
		sr.theMap.points[v.x][v.y] = "*"
	}
	sr.theMap.points[sr.start.x][sr.start.y] = "S"
	sr.theMap.points[sr.end.x][sr.end.y] = "E"
	for _, v := range sr.theMap.points {
		for _, view := range v {
			fmt.Printf("%s ", view)
		}
		fmt.Println()
	}
}
