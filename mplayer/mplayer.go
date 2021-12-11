package mplayer

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type MPlayer struct {
	cmd                  *exec.Cmd
	stdin                io.WriteCloser
	stdout               io.ReadCloser
	scanner              *bufio.Scanner
	commandResultChannel chan string
}

func New() *MPlayer {
	cmd := exec.Command("mplayer", "-softvol", "-quiet", "-slave", "-idle")
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(stdout)
	mplayer := MPlayer{cmd, stdin, stdout, scanner, make(chan string)}
	go func() {
		for mplayer.scanner.Scan() {
			s := mplayer.scanner.Text()
			if strings.Contains(s, "ANS_PERCENT_POSITION=") {
				mplayer.commandResultChannel <- strings.Split(s, "ANS_PERCENT_POSITION=")[1]
			}
			if strings.Contains(s, "ANS_TIME_POSITION=") {
				mplayer.commandResultChannel <- strings.Split(s, "ANS_TIME_POSITION=")[1]
			}
			if strings.Contains(s, "ANS_LENGTH=") {
				mplayer.commandResultChannel <- strings.Split(s, "ANS_LENGTH=")[1]
			}
		}
	}()
	if err := mplayer.cmd.Start(); err != nil {
		panic(err)
	}
	return &mplayer
}

func (mplayer *MPlayer) LoadFile(path string) {
	mplayer.stdin.Write([]byte(fmt.Sprintf("loadfile %s\n", path)))
}

func (mplayer *MPlayer) CurrentPercentPosition() string {
	mplayer.stdin.Write([]byte("get_percent_pos\n"))
	result := <-mplayer.commandResultChannel
	return result
}

func (mplayer *MPlayer) CurrentTimePosition() string {
	mplayer.stdin.Write([]byte("get_time_pos\n"))
	result := <-mplayer.commandResultChannel
	return result
}

func (mplayer *MPlayer) TotalLength() string {
	mplayer.stdin.Write([]byte("get_time_length\n"))
	result := <-mplayer.commandResultChannel
	return result
}

func (mplayer *MPlayer) Pause() {
	mplayer.stdin.Write([]byte("pause\n"))
}

func (mplayer *MPlayer) Volume(volume string) {
	mplayer.stdin.Write([]byte(fmt.Sprintf("volume %s 1\n", volume)))
}

func (mplayer *MPlayer) Quit() {
	mplayer.stdin.Write([]byte("quit\n"))
}
