package main

import (
	"github.com/mingz2013/chinese-chess-table-go/msg"
	robot2 "github.com/mingz2013/chinese-chess-table-go/robot"
	table2 "github.com/mingz2013/chinese-chess-table-go/table"
	"log"
	"sync"
	"time"
)

type Processor interface {
	Run()
}

func runProcessor(wg sync.WaitGroup, p Processor) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		p.Run()

	}()
}

type RobotContext struct {
	Robot  robot2.Robot
	MsgIn  chan msg.Msg
	MsgOut chan msg.Msg
}

func makeRobots() []RobotContext {
	var robots []RobotContext
	for i := 0; i < 4; i++ {
		robotMsgIn := make(chan msg.Msg)
		robotMsgOut := make(chan msg.Msg)

		robot := robot2.NewRobot(i+1000, "", robotMsgIn, robotMsgOut)

		robots = append(robots, RobotContext{robot, robotMsgIn, robotMsgOut})

	}
	return robots
}

func main() {
	//sdk := sdk2.MakerSdk("1")

	//tableManager := table.NewTableManager("1")

	var wg sync.WaitGroup // 工作goroutine个数

	//go sdk.Run()
	//go tableManager.Run()
	//RunProcessor(wg, sdk)
	//RunProcessor(wg, tableManager)

	tableMsgIn := make(chan msg.Msg)
	tableMsgOut := make(chan msg.Msg)

	table := table2.NewTable(1, tableMsgIn, tableMsgOut)

	//var robots []RobotContext

	robots := makeRobots()

	log.Println("make obj down")

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {

			select {
			case m, ok := <-tableMsgOut:
				{
					//m, ok := <-tableMsgOut
					log.Println("on msg table", m)
					if !ok {
						continue
					}
					id := m["id"].(int)

					for i := 0; i < len(robots); i++ {
						if robots[i].Robot.Id == id {
							robots[i].MsgIn <- m
						}
					}
				}

			case <-time.After(1 * time.Second):
				continue

			}

		}

	}()

	for i := 0; i < len(robots); i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			for {

				select {
				case m, ok := <-robots[index].MsgOut:
					{
						log.Println("on msg robot", index, m)
						if !ok {
							continue
						}
						tableMsgIn <- m
					}
				case <-time.After(1 * time.Second):
					continue

				}

			}

		}(i)
	}

	log.Println("bound ch down")

	runProcessor(wg, table)

	for i := 0; i < len(robots); i++ {

		runProcessor(wg, robots[i].Robot)
	}

	log.Println("run down")

	wg.Wait()
}
