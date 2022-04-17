package store


//存储层抽象接口

type CmdFactory interface {
	Check() CheckCmd
}
