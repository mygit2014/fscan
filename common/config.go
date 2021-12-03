package common

var Userdict = map[string][]string{
	"ftp":        {"ftp", "admin", "www", "web", "root", "db", "wwwroot", "data"},
	"mysql":      {"root"},
	"mssql":      {"sa"},
	"smb":        {"administrator", "admin"},
	"postgresql": {"postgres"},
	"ssh":        {"root"},
	"mongodb":    {"root", "admin"},
}

var Passwords = []string{"123456", "admin", "admin123", "root", "", "pass123", "pass@123", "password", "123123", "654321", "111111", "123", "1", "admin@123", "Admin@123", "admin123!@#", "{user}", "{user}1", "{user}111", "{user}123", "{user}@123", "{user}_123", "{user}#123", "{user}@111", "{user}@2019", "P@ssw0rd!", "P@ssword", "p@ssword", "P@ssw0rd", "Passw0rd", "qwe123", "12345678", "test", "test123", "123qwe!@#", "123456789", "123321", "666666", "a123456.", "123456~a", "000000", "1234567890", "8888888", "!QAZ2wsx", "1qaz2wsx", "abc123", "abc123456", "1qaz@WSX", "a11111", "a12345", "Aa1234", "Aa1234.", "Aa12345", "a123456", "a123123", "Aa123123", "Aa123456", "Aa12345.","qwer1234"}
var PORTList = map[string]int{
	"ftp":         21,
	"ssh":         22,
	"mem":         11211,
	"mgo":         27017,
	"mssql":       1433,
	"psql":        5432,
	"redis":       6379,
	"mysql":       3306,
	"smb":         445,
	"ms17010":     1000001,
	"cve20200796": 1000002,
	"web":         1000003,
	"findnet":     135,
	"netbios":     139,
	"all":         0,
	"portscan":    0,
	"icmp":        0,
	"main":        0,
}

var Outputfile = getpath() + "result.txt"
var IsSave = true
var Webport = "80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,443,800,801,808,61000,888,889,1000,1010,1080,1081,1082,1118,1888,2008,2020,2100,3000,3008,3505,5555,6080,6648,6868,7000,7001,7002,7003,7004,7005,7007,7008,7070,7071,7074,7078,7080,7088,7200,7680,7687,7688,7777,7890,8000,8001,8002,8003,8004,8006,8008,8009,8010,8070,8080,8081,8082,8083,8084,8085,8086,8087,8088,8089,8090,8091,8092,8093,8094,8095,8096,8097,8098,8099,8100,8101,8102,8118,8103,8180,8181,8200,8300,8443,8448,8800,8848,8880,8881,8888,8899,9777,8890,8989,9000,9001,9002,9008,9010,9080,9081,9082,9087,9088,9089,9090,9091,9092,9093,9094,9095,9096,9097,9098,9099,9100,9200,9443,9800,9988,9999,10000,10001,10002,10004,10008,10010,11000,18001,18002,18080,19000,19999,14777,6033,33060,8500,813"
var DefaultPorts = "22,80,81,135,443,445,1433,3306,5432,6379,7001,8000,8080,8089,3389,2222"

type HostInfo struct {
	Host      string
	Ports     string
	Domain    string
	Url       string
	Timeout   int64
	Scantype  string
	Command   string
	Username  string
	Password  string
	Usernames []string
	Passwords []string
}

type PocInfo struct {
	Num        int
	Rate       int
	Timeout    int64
	Proxy      string
	PocName    string
	PocDir     string
	Target     string
	TargetFile string
	RawFile    string
	Cookie     string
	ForceSSL   bool
	ApiKey     string
	CeyeDomain string
}

var (
	TmpOutputfile string
	TmpSave       bool
	IsPing        bool
	Ping          bool
	Pocinfo       PocInfo
	IsWebCan      bool
	RedisFile     string
	RedisShell    string
	Userfile      string
	Passfile      string
	HostFile      string
	Threads       int
	URL           string
	UrlFile       string
	Urls          []string
	NoPorts       string
)
