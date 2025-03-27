package main

import "fmt"

type notification interface {
    importance() int
}

type directMessage struct {
    senderUsername string
    messageContent string
    priorityLevel  int
    isUrgent       bool
}

type groupMessage struct {
    groupName      string
    messageContent string
    priorityLevel  int
}

type systemAlert struct {
    alertCode      string
    messageContent string
}

func (dm directMessage) importance() int {
    if dm.isUrgent {
        return 50
    } else {
        return dm.priorityLevel
    }
}

func (gm groupMessage) importance() int {
    return gm.priorityLevel
}

func (sm systemAlert) importance() int {
    return 100
}

func processNotification(n notification) (string, int) {
    switch v := n.(type) {
    case directMessage:
        return v.senderUsername, v.importance()
    case groupMessage:
        return v.groupName, v.importance()
    case systemAlert:
        return v.alertCode, v.importance()
    default:
        return "", 0
    }
}

func main() {
    // Create instances of each notification type
    dm := directMessage{
        senderUsername: "Alice",
        messageContent: "Hello, how are you?",
        priorityLevel:  10,
        isUrgent:       true,
    }

    gm := groupMessage{
        groupName:      "Developers Group",
        messageContent: "Meeting at 3 PM",
        priorityLevel:  20,
    }

    sa := systemAlert{
        alertCode:      "ALERT123",
        messageContent: "System maintenance scheduled.",
    }

    // Process each notification
    dmSender, dmImportance := processNotification(dm)
    fmt.Printf("Direct Message - Sender: %s, Importance: %d\n", dmSender, dmImportance)

    gmGroup, gmImportance := processNotification(gm)
    fmt.Printf("Group Message - Group: %s, Importance: %d\n", gmGroup, gmImportance)

    saCode, saImportance := processNotification(sa)
    fmt.Printf("System Alert - Code: %s, Importance: %d\n", saCode, saImportance)
}