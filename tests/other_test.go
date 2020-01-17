package test

import (
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"
)

func Test_Constr(t *testing.T) {
	//errstr := "\\n### Error updating database. Cause: com.mysql.jdbc.MysqlDataTruncation: Data truncation: Data too long for column 'handle_desc' at row 1\\n### The error may involve com.hq.dao.base.OfflineTransferMapper.updateByPrimaryKey-Inline\\n### The error occurred while setting parameters\\n### SQL: update offline_transfer set type = ?, process_id = ?, request_date = ?, name = ?, id_card_encode = ?, mobile_encode = ?, league_for_years = ?, the_party_years = ?, development_member_number = ?, cause = ?, nation = ?, nation_code = ?, political_outlook = ?, political_outlook_code = ?, education_degree = ?, education_degree_code = ?, occupation = ?, remark = ?, email = ?, qq_num = ?, weChat_id = ?, weibo = ?, study_work_unit = ?, from_org_name = ?, from_org_code = ?, to_org_name = ?, target_org_name = ?, to_org_code = ?, handle_result = ?, handle_result_code = ?, handle_desc = ?, response_date = ?, otid = ? where id = ?\\n### Cause: com.mysql.jdbc.MysqlDataTruncation: Data truncation: Data too long for column 'handle_desc' at row 1\\n; SQL []; Data truncation: Data too long for column 'handle_desc' at row 1; nested exception is com.mysql.jdbc.MysqlDataTruncation: Data truncation: Data too long for column 'handle_desc' at row 1"
	//if strings.Contains(errstr, "Data too long for column") {
	//	errstr = fmt.Sprintf("审批原因过长,%s(三省)接收失败,请减少审批原因字数", db.Get3ShengNameByOrgCode("FJAN"))
	//}
	//fmt.Println(errstr)
	//fmt.Println(time.Now())

	///------------------------------------分割线-----------------------------------

	var sum *int
	sum = new(int)
	*sum = 99
	fmt.Println(*sum)


}




// ******************************分割线******************************

func echo(conn *net.TCPConn) {
	tick := time.Tick(5 * time.Second) // 五秒的心跳间隔
	for now := range tick {
		now.String()
		//n, err := conn.Write([]byte(now.String()))
		n, err := conn.Write([]byte("nihao你好"))
		fmt.Printf("send %d bytes to %s\n", n, conn.RemoteAddr())
		if err != nil {
			log.Println("Error:: ",err)
			conn.Close()
			return
		}
		fmt.Printf("send %d bytes to %s\n", n, conn.RemoteAddr())
	}
}
func Test_Server(t *testing.T) {
	address := net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"), // 把字符串IP地址转换为net.IP类型
		Port: 8000,
	}
	listener, err := net.ListenTCP("tcp4", &address) // 创建TCP4服务器端监听器
	if err != nil {
		log.Fatal(err) // Println + os.Exit(1)
	}
	for {
		conn, err := listener.AcceptTCP()
		fmt.Println("=========================")
		if err != nil {
			log.Fatal(err) // 错误直接退出
		}
		fmt.Println("remote address:", conn.RemoteAddr())
		go echo(conn)
	}
}
func Test_Client(t *testing.T) {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s host:port", os.Args[0])
	}
	//os.Args[0] = "127.0.0.1"
	//os.Args[1] = "127.0.0.1:8000"
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	n, err := conn.Write([]byte("HEAD / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(n)
}

// ********************************** 分割线 *****************************************
func Test_ClientUDP(t *testing.T){
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		handleClient(conn)
	}
}
func handleClient(conn * net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
