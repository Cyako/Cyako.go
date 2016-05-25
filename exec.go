package cyako

import (
	"fmt"
	// stat "github.com/Centimitr/xmessage/statistics"
	"encoding/json"
	"golang.org/x/net/websocket"
)

func (c *Cyako) do(ws *websocket.Conn, req *Req) {
	var err error

	// Phase I: AfterReceive
	// - global, req relative methods
	c.AfterReceive(req)
	// - initial context and response
	res := &Res{Id: req.Id, Method: req.Method, Temp: req.Temp}
	ctx := &Ctx{res: res, req: req, Method: req.Method, Data: req.Data, Temp: req.Temp}
	res.Init()
	ctx.Init()

	// Phase II: BeforeProcess
	// - global, ctx relative methods
	// - provide chance to manipulate context object for middlewares
	c.BeforeProcess(ctx)

	// Phase III: Process
	// - match and select processor, and then execute it on ctx
	process, err := c.matchProcessor(req.Method)
	if err != nil {
		fmt.Println(err)
		return
	}
	process(ctx)

	// Phase IV: AfterProcess
	// - global, ctx relative methods
	c.AfterProcess(ctx)
	// - mainly do response relative tasks
	// res.Data = ctx.Data
	data, err := json.Marshal(ctx.Data)
	if err != nil {
		fmt.Println(err)
	}
	res.Data = string(data)
	ctx.setResParams()

	// Phase V: BeforeSend
	c.BeforeSend(res)

	// Phase VI: Send
	// - send
	if err := websocket.JSON.Send(ws, res); err != nil {
		fmt.Println("SEND ERROR.", err)
		return
	}

	// Phase VI: AfterSend
	c.AfterSend(res)
}
