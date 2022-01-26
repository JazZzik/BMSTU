package main
 
import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net"
    "os/exec"
    "sync"
 
    "github.com/kr/pty"
    "golang.org/x/crypto/ssh"
)
 
func passwordCallback(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
    if (c.User() == "alex" && string(pass) == "123") {
        return nil, nil
    }
    return nil, fmt.Errorf("Password rejected for %q", c.User())
}
 
func handleChannels(chans <-chan ssh.NewChannel) {
    for newChannel := range chans {
        go handleChannel(newChannel)
    }
}
 
func handleChannel(newChannel ssh.NewChannel) {
    if t := newChannel.ChannelType(); t != "session" {
        newChannel.Reject(ssh.UnknownChannelType, fmt.Sprintf("unknown channel type: %s", t))
        return
    }
    connection, _, err := newChannel.Accept()
    if err != nil {
        log.Printf("Could not accept channel (%s)", err)
        return
    }
 
    bash := exec.Command("bash")
    close := func() {
        connection.Close()
        _, err := bash.Process.Wait()
        if err != nil {
            log.Printf("Failed to exit bash (%s)", err)
        }
        log.Printf("Session closed")
    }
 
    log.Print("Creating pty...")
    bashf, err := pty.Start(bash)
    log.Print("Created...")
    if err != nil {
        log.Printf("Could not start pty (%s)", err)
        close()
        return
    }
 
    var once sync.Once
    go func() {
        io.Copy(connection, bashf)
        once.Do(close)
    }()
    go func() {
        io.Copy(bashf, connection)
        once.Do(close)
    }()
}
 
 
func main() {
    config := &ssh.ServerConfig{
        PasswordCallback: passwordCallback,
    }
 
    privateBytes, err := ioutil.ReadFile("/home/alex/.ssh/id_rsa")
    if err != nil {
        log.Fatal("Failed to load private key (./id_rsa)")
    }
 
    private, err := ssh.ParsePrivateKey(privateBytes)
    if err != nil {
        log.Fatal("Failed to parse private key")
    }
 
    config.AddHostKey(private)
 
 
    listener, err := net.Listen("tcp", "0.0.0.0:2200")
    if err != nil {
        log.Fatalf("Failed to listen on 2200 (%s)", err)
    }
 
    log.Print("Listening on 2200...")
    for {
        tcpConn, err := listener.Accept()
        if err != nil {
            log.Printf("Failed to accept incoming connection (%s)", err)
            continue
        }
        sshConn, chans, reqs, err := ssh.NewServerConn(tcpConn, config)
        if err != nil {
            log.Printf("Failed to handshake (%s)", err)
            continue
        }
 
        log.Printf("New SSH connection from %s (%s)", sshConn.RemoteAddr(), sshConn.ClientVersion())
        go ssh.DiscardRequests(reqs)
        go handleChannels(chans)
    }
}
