package protocols
import (
    NEX "github.com/Stary2001/nex-go"
    "fmt"
    )

func Secure_Connection_Register(VecMyURLs []NEX.StationURL) (rmcResult uint32, returnValue NEX.Result, PidConnectionID uint32, UrlPublic NEX.StationURL) {
    rmcResult = 0x00010001
    returnValue = 0x00010001
    PidConnectionID = 2737106758
    UrlPublic = VecMyURLs[0]
    fmt.Println(VecMyURLs)
    return
}

func Secure_Connection_RequestConnectionData(CidTarget uint32, PidTarget uint32) (rmcResult uint32, returnValue bool, PvecConnectionsData []NEX.ConnectionData) {
    rmcResult = 0x80010002
    return
}

func Secure_Connection_RequestUrls(CidTarget uint32, PidTarget uint32) (rmcResult uint32, returnValue bool, PlstURLs []NEX.StationURL) {
    rmcResult = 0x80010002
    return
}

func Secure_Connection_RegisterEx(VecMyURLs []NEX.StationURL, HCustomData NEX.Data) (rmcResult uint32, returnValue NEX.Result, PidConnectionID uint32, UrlPublic NEX.StationURL) {
    rmcResult = 0x00010001
    returnValue = 0x00010001
    PidConnectionID = 2750392109
    UrlPublic = "prudp:/address=69.69.69.69;port=61679;natf=0;natm=0;pmp=0;sid=15;type=3;upnp=0"
    fmt.Println(VecMyURLs)
    return
}


func Secure_Connection_TestConnectivity() (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}

func Secure_Connection_UpdateURLs(VecMyURLs []NEX.StationURL) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}

func Secure_Connection_ReplaceURL(Target NEX.StationURL, Url NEX.StationURL) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}

func Secure_Connection_SendReport(ReportId uint32, ReportData NEX.QBuffer) (rmcResult uint32, ) {
    rmcResult = 0x80010002
    return
}
