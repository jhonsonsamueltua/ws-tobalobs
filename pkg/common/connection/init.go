package connection

import (
	"database/sql"
	"fmt"
	"net"
	"os"

	// _ "github.com/lib/pq"
	"github.com/go-redis/redis"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type ViaSSHDialer struct {
	client *ssh.Client
}

func (self *ViaSSHDialer) Dial(addr string) (net.Conn, error) {
	return self.client.Dial("tcp", addr)
}

func ConnectSSH() {
	sshHost := "66.70.190.240" // SSH Server Hostname/IP
	sshPort := 22              // SSH Port
	sshUser := "ubuntu"        // SSH Username
	sshPass := "WKDW3eQr"      // Empty string for no password
	// dbUser := "jhonson-sth"    // DB username
	// dbPass := "hutagaol2020"   // DB Password
	// dbHost := "localhost:3306" // DB Hostname/IP
	// dbName := "tobalobs"       // Database name

	var agentClient agent.Agent
	// Establish a connection to the local ssh-agent
	if conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		defer conn.Close()

		// Create a new instance of the ssh agent
		agentClient = agent.NewClient(conn)
	}

	// The client configuration with configuration option to use the ssh-agent
	sshConfig := &ssh.ClientConfig{
		User:            sshUser,
		Auth:            []ssh.AuthMethod{},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// When the agentClient connection succeeded, add them as AuthMethod
	if agentClient != nil {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeysCallback(agentClient.Signers))
	}
	// When there's a non empty password add the password AuthMethod
	if sshPass != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PasswordCallback(func() (string, error) {
			return sshPass, nil
		}))
	}

	// Connect to the SSH Server
	sshcon, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", sshHost, sshPort), sshConfig)
	if err == nil {
		// defer sshcon.Close()
		// Now we register the ViaSSHDialer with the ssh connection as a parameter
		mysql.RegisterDial("mysql+tcp", (&ViaSSHDialer{sshcon}).Dial)

		// And now we can use our new driver with the regular mysql connection string tunneled through the SSH connection

		// if db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@mysql+tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)); err == nil {

		// 	fmt.Println("Successfully connected DB!")

		// 	if rows, err := db.Query("SELECT user_id, username FROM user ORDER BY user_id"); err == nil {
		// 		for rows.Next() {
		// 			var id int64
		// 			var name string
		// 			rows.Scan(&id, &name)
		// 			fmt.Printf("ID: %d  Name: %s\n", id, name)
		// 		}
		// 		rows.Close()
		// 	} else {
		// 		fmt.Printf("Failure: %s", err.Error())
		// 	}

		// 	db.Close()

		// } else {

		// 	fmt.Printf("Failed to connect to the db: %s\n", err.Error())
		// }

	} else {
		fmt.Printf("Failed to connect to the ssh server: %s\n", err.Error())
	}
}

func InitDB(conn string) *sql.DB {
	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected DB!")
	return db
}

func InitRedis(addr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
		Network:  "tcp",
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected Redis!")
	return client
}
