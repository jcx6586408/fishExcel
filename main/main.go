package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Level 关卡
type Level struct {
	ID       int    `json:"id"`
	Exp      int    `json:"exp"`
	Prefab   string `json:"prefab"`
	SelfID   int    `json:"self"`
	Soliders []int  `json:"soliders"`
	BOSS     []int  `json:"boss"`
	Foods    []int  `json:"foods"`
	Bombs    []int  `json:"bombs"`
}

var levels = []*Level{}

// GameObject 游戏物体
type GameObject struct {
	ID          int    `json:"id"`
	Player      bool   `json:"player"`
	Avatar      string `json:"avatar"`
	Num         int    `json:"num"`
	OriginLevel int    `json:"level"`
	Speed       int    `json:"speed"`
	RebornDur   int    `json:"rebornDur"`
	Exp         int    `json:"exp"`
	Type        int    `json:"type"`
	SubType     int    `json:"subType"`
}

var gameObjects = []*GameObject{}

func debug(err error) {
	if err != nil {
		panic("AI条件表导出出错")
	}
}

func main() {
	println("吃鱼表格解析工具")
	f, err := excelize.OpenFile("./关卡.xlsx")
	if err != nil {
		println(err.Error())
		return
	}

	for _, sheetName := range f.GetSheetMap() {

		if sheetName == "level" {
			rows := f.GetRows(sheetName)
			for i, row := range rows {
				if i == 0 || i == 1 {
					continue
				}

				level := &Level{}

				id, ok := strconv.Atoi(row[0])
				debug(ok)

				level.ID = id

				exp, ok := strconv.Atoi(row[1])
				debug(ok)

				level.Exp = exp

				prefab := row[2]

				level.Prefab = prefab

				selfID, ok := strconv.Atoi(row[3])
				debug(ok)

				level.SelfID = selfID

				soliders := strings.Split(row[4], "#")
				for _, sid := range soliders {
					nid, ok := strconv.Atoi(sid)
					if ok == nil {
						level.Soliders = append(level.Soliders, nid)
					}
				}

				boss := strings.Split(row[5], "#")
				for _, sid := range boss {
					nid, ok := strconv.Atoi(sid)
					if ok == nil {
						level.BOSS = append(level.BOSS, nid)
					}
				}

				foods := strings.Split(row[6], "#")
				for _, sid := range foods {
					nid, ok := strconv.Atoi(sid)
					if ok == nil {
						level.Foods = append(level.Foods, nid)
					}
				}

				bombs := strings.Split(row[7], "#")
				for _, sid := range bombs {
					nid, ok := strconv.Atoi(sid)
					if ok == nil {
						level.Bombs = append(level.Bombs, nid)
					}
				}

				levels = append(levels, level)
			}

		}

		if sheetName == "fish" {
			for i, row := range f.GetRows(sheetName) {
				if i == 0 || i == 1 {
					continue
				}

				gameObject := &GameObject{}

				id, ok := strconv.Atoi(row[0])
				debug(ok)

				gameObject.ID = id

				player, ok := strconv.Atoi(row[1])
				debug(ok)

				if player == 1 {
					gameObject.Player = true
				} else {
					gameObject.Player = false
				}

				gameObject.Avatar = row[2]

				num, ok := strconv.Atoi(row[4])
				debug(ok)

				gameObject.Num = num

				level, ok := strconv.Atoi(row[5])
				debug(ok)

				gameObject.OriginLevel = level

				speed, ok := strconv.Atoi(row[6])
				debug(ok)

				gameObject.Speed = speed

				reborn, ok := strconv.Atoi(row[7])
				debug(ok)

				gameObject.RebornDur = reborn

				exp, ok := strconv.Atoi(row[8])
				debug(ok)

				gameObject.Exp = exp

				t, ok := strconv.Atoi(row[9])
				debug(ok)

				gameObject.Type = t

				st, ok := strconv.Atoi(row[10])
				debug(ok)

				gameObject.SubType = st

				gameObjects = append(gameObjects, gameObject)
			}
		}
	}

	// 创建Json编码器
	fileLevel, err := os.Create("./assets/config/" + "level" + ".json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}

	defer fileLevel.Close()

	dataLevel, err := json.MarshalIndent(levels, "", "\t")
	fileLevel.Write(dataLevel)
	println("关卡表：", "文件导出成功")

	// 创建Json编码器
	filePtr, err := os.Create("./assets/config/" + "gameObject" + ".json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}

	defer filePtr.Close()

	data, err := json.MarshalIndent(gameObjects, "", "\t")
	filePtr.Write(data)
	println("怪物表：", "文件导出成功")
}
