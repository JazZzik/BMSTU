package main
 
import (
    "github.com/gliderlabs/ssh"
    "log"
    "bytes"
    "strings"
    "os/exec"
    "golang.org/x/crypto/ssh/terminal"
)
 
// получаю аргументы из командной строки клиента
// args[0] - команда, args[1:] - оставшиеся аргументы (ключи, путь)
func getArgs(line string) (string, []string) {
    args := strings.Split(line, " ")
    if len(args) < 1 {
        return line, nil
    }
    return args[0], args[1:]
}
 
// выполняю команду, переданную клиентом, на сервере
func runCommand(command string, args []string) (error, string, string) {
    var cmd *exec.Cmd // тип "команда"
    if args == nil {
        cmd = exec.Command(command) // возвращаю структуру для выполнения, "команда"
    } else {
        cmd = exec.Command(command, args...) // возвращаю структуру для выполнения, "команда + аргументы"
    }
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &stdout // перенаправляю вывод
    cmd.Stderr = &stderr
    err := cmd.Run() // выполняю команду
    return err, stdout.String(), stderr.String()
}
 
func main() {
    ssh.Handle(func(sess ssh.Session) { // задаю дефолтный обработчик запросов
        term := terminal.NewTerminal(sess, "> ")
        for {
            line, err := term.ReadLine() //считываю запрос клиента
            if err != nil {
                break
            }
            command, args := getArgs(line) //разбиваю запрос клиента на команду и слайс аргументов
            log.Println(command)
            if command != "dir" && command != "mkdir" && command != "rmdir" { //неизвестная команда
                log.Println("error: invalid command")
                term.Write(append([]byte("invalid command"), '\n'))
            } else {
                err, out, errout := runCommand(command, args) // выполняю запрос
                if err != nil {
                    log.Println("error: ", err)
                }
                if out != "" { // вывожу поток вывода - результат выполнения команды
                    log.Println(out)
                    term.Write(append([]byte(out), '\n'))
                }
                if errout != "" { // вывожу поток ошибок
                    log.Println(errout)
                    term.Write(append([]byte(errout), '\n'))
                }
            }
        }
        log.Println("terminal closed")
    })
 
    // принимаю запросы на 2222 порту дефолтным обработчиком, указываю путь к публичному ключу
    log.Fatal(ssh.ListenAndServe(":2222", nil, ssh.HostKeyFile("/home/alex/.ssh/id_rsa")))
}
