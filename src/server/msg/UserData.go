package msg

import (
	"image"
	"time"
)

///保存和同步用户

type UserData struct {
	Img image.Image  `json:"img"`///图片数据
	OnlineTime time.Duration `json:"online_time"`///用户在线的时间
	Friends map[string]string ///用户的好友
	///拥有的在线的车辆
}
