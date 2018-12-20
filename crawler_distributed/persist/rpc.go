package persist

import (
	"jsj.golangtc/crawler/zhenai/engine"
	"gopkg.in/olivere/elastic.v5"
	"time"
	"strconv"
	"encoding/json"
	"fmt"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (ItemSaverService) Save(item engine.Item, result *string) error {
	unix := time.Now().Unix()
	bytes, _ := json.Marshal(item)

	*result = fmt.Sprintf(" %s ok：%s", bytes, strconv.Itoa(int(unix)))
	return nil
}
