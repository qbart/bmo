package bmo

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

// YeeBulb data.
type YeeBulb struct {
	addr string
}

// YeeCmd json.
type YeeCmd struct {
	ID     int64      `json:"id"`
	Method string     `json:"method"`
	Params []yeeParam `json:"params"`
}

// YeeError struct.
type YeeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// YeeResponse struct.
type YeeResponse struct {
	ID     int64    `json:"id"`
	Result []string `json:"result"`
	Error  YeeError `json:"error"`
}

type yeeParam interface{}

// NewYeeBulb creates a new yee based on IP.
func NewYeeBulb(ip string) *YeeBulb {
	return &YeeBulb{addr: fmt.Sprint(ip, ":", 55443)}
}

// Power toggles bulb on/off.
func (y *YeeBulb) Power(on bool) {
	value := "on"
	if !on {
		value = "off"
	}
	y.sendMessage("set_power", []yeeParam{value, "smooth", 500})
}

// Brightness sets bright parameter.
func (y *YeeBulb) Brightness(value int) {
	value = Clamp(value, 0, 100)
	y.sendMessage("set_bright", []yeeParam{value, "smooth", 500})
}

// Color sets RGB color.
func (y *YeeBulb) Color(r, g, b uint8) {
	color := int(r) * 65536 + int(g) * 256 + int(b)
	y.sendMessage("set_rgb", []yeeParam{color, "smooth", 500})
}

func (y *YeeBulb) sendMessage(method string, params []yeeParam) YeeResponse {
	conn, err := net.Dial("tcp", y.addr)
	if err != nil {
		return YeeResponse{
			Error: YeeError{
				Message: "Failed to connect to Yee Bulb",
			},
		}
	}
	defer conn.Close()

	// prepare command
	cmd := YeeCmd{
		ID:     time.Now().Unix(),
		Method: method,
		Params: params,
	}
	jsonCmd, err := json.Marshal(cmd)
	if err != nil {
		return YeeResponse{
			Error: YeeError{
				Message: "JSON marshaling failed",
			},
		}
	}

	// send command to bulb (must end with CRLF)
	fmt.Fprintf(conn, "%s\r\n", jsonCmd)
	// fmt.Println("[DEBUG] ", string(jsonCmd))

	// read response
	var buf [4096]byte
	read, _ := conn.Read(buf[:])
	var resp YeeResponse
	if err := json.Unmarshal(buf[:read], &resp); err != nil {
		return YeeResponse{
			Error: YeeError{
				Message: "JSON unmarshaling failed",
			},
		}
	}

	return resp
}
