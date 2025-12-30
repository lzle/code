package global

const (
	RequestSuccess      =   200          // 请求成功

	JsonParseError      =   4001        //json解析失败
	ParamReadError      =   4002	    //参数读取失败
	ParamMissingError   =   4003        //参数缺失
	CallNotExistError   =   4004        //call不存在
	CallerNotMatchError =   4005       //主叫不匹配
	AuthEncryptError    =   4006       //秘钥不对
	GetServerIdError    =   4006       //cti serverid 获取失败
	FileNotExists       =   4007	   // 文件不存在
	SaveFileError       =   4008	  // 保存文件失败
	FileEmpty           =   4009	  // 文件为空
)
