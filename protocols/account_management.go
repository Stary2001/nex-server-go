package protocols

import (
	NEX "github.com/Stary2001/nex-go"
	)

func Account_Management_NintendoCreateAccount(client *NEX.Client, StrPrincipalName string, StrKey string, UiGroups uint32, StrEmail string, OAuthData NEX.Data) (rmcResult uint32, Pid NEX.PID, PidHMAC string) {
    rmcResult = 0x80030066 // Already exists
    return
}