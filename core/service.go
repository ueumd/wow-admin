package core

func InitService() error {
	var err error
	if err = RedisInit();err != nil {
		return err
	}
	return nil
}
