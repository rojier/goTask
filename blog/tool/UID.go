package tool

/*
1、如果需要标准的、跨系统的唯一ID，使用UUID。
2、 如果需要数字形式的、有时间顺序的ID，并且是分布式系统，使用雪花算法。

*/
import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

/*
这种方法生成的UID长度较长，但碰撞概率极低

如果对ID格式没有特殊要求，且不需要分布式，可以使用基于时间戳和随机数的方法
*/
func GenerateUID() string {
	// 获取当前时间戳（纳秒）
	timestamp := time.Now().UnixNano()
	// 生成随机数
	randomBytes := make([]byte, 10)
	rand.Read(randomBytes)
	randomStr := hex.EncodeToString(randomBytes)
	// 组合时间戳和随机数
	return fmt.Sprintf("%d%s", timestamp, randomStr)
}

// Snowflake 雪花算法结构
type Snowflake struct {
	mutex     sync.Mutex
	lastTime  int64
	machineID int64
	sequence  int64
}

// NewSnowflake 创建雪花算法实例
func NewSnowflake(machineID int64) *Snowflake {
	return &Snowflake{
		machineID: machineID,
	}
}

// Generate 生成唯一ID
func (s *Snowflake) Generate() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	currentTime := time.Now().UnixNano() / 1e6 // 转换为毫秒

	if currentTime < s.lastTime {
		panic("时钟回拨异常")
	}

	if currentTime == s.lastTime {
		s.sequence = (s.sequence + 1) & 4095 // 4095是序列号最大值
		if s.sequence == 0 {
			// 当前毫秒序列号用完，等待下一毫秒
			for currentTime <= s.lastTime {
				currentTime = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastTime = currentTime

	// 组成64位ID: 时间戳(41位) + 机器ID(10位) + 序列号(12位)
	return (currentTime << 22) | (s.machineID << 12) | s.sequence
}

// func main() {
// 	// 创建雪花算法实例，机器ID为1
// 	snowflake := NewSnowflake(1)

// 	// 生成多个ID
// 	for i := 0; i < 10; i++ {
// 		id := snowflake.Generate()
// 		fmt.Printf("生成的ID: %d\n", id)
// 	}
// }
