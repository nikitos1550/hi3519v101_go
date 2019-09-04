package hisi

import (
    "sync"
)

var (
    chns channels
)

type channel struct {
	lock    sync.Mutex
	id      uint
	enabled uint
}

type channels struct {
	maxNum    uint
	minWidth  uint
	maxWidth  uint
	minHeight uint
	maxHeight uint
	maxFps    uint
	chn       []channel
}

/*
func ChannelsMaxNum() int {
	return int(chns.maxNum)
}

func ChannelGetInfo(id uint) (ChannelInfo, int) {
	var out ChannelInfo

	if id >= chns.maxNum {
		return out, ERR_OBJ_NOT_FOUND
	}
	var c *channel = &chns.chn[id]

	c.lock.Lock()
	defer c.lock.Unlock()

	//log.Println("id ", c.id, " enabled ", c.enabled)

	if c.enabled == 1 {
		var params C.struct_channel_params
		if C.hisi_channel_info(C.uint(c.id), &params) != 0 {
			panic("1")
		}
		//log.Println("\twidth ", tmp.width, " height ", tmp.height, " fps ", tmp.fps)
		out.Enabled = 1
		out.Width = int(params.width)
		out.Height = int(params.height)
		out.Fps = int(params.fps)
	}

	return out, ERR_NONE
}

func ChannelEnable(id uint, w, h, fps int) int {
	if id >= chns.maxNum {
		return ERR_OBJ_NOT_FOUND
	}
	var c *channel = &chns.chn[id]

	c.lock.Lock()
	defer c.lock.Unlock()

	if c.enabled == 1 {
		return ERR_NOT_ALLOWED
	}

	var params C.struct_channel_params
	params.width = C.int(w)
	params.height = C.int(h)
	params.fps = C.int(fps)
	if C.hisi_channel_enable(C.uint(c.id), &params) != 0 {
		panic("1")
	}

	c.enabled = 1

	return ERR_NONE
}

func ChannelDisable(id uint) int {
	if id >= chns.maxNum {
		return ERR_OBJ_NOT_FOUND
	}
	var c *channel = &chns.chn[id]

	c.lock.Lock()
	defer c.lock.Unlock()

	if c.enabled == 0 {
		return ERR_NOT_ALLOWED
	}

	if C.hisi_channel_disable(C.uint(c.id)) != 0 {
		panic("1")
	}

	c.enabled = 0

	return ERR_NONE
}
*/
