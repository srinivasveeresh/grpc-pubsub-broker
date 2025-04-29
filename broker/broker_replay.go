package broker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const maxMessagesPerTopic = 10000

func (b *Broker) trimOldMessages(topic string) {
	filePath := fmt.Sprintf("data/messages_%s.log", topic)
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) > maxMessagesPerTopic {
		lines = lines[len(lines)-maxMessagesPerTopic:]
		os.WriteFile(filePath, []byte(strings.Join(lines, "\n")+"\n"), 0644)
	}
}

func (b *Broker) replayMessages(topic string, c chan<- string) {
	filePath := fmt.Sprintf("data/messages_%s.log", topic)
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		select {
		case c <- scanner.Text():
		case <-time.After(50 * time.Millisecond):
			// non-blocking replay
		}
	}
}
