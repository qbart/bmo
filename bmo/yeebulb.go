package bmo

type Any interface{}

type YeeBulb struct {
	host string
	port int
}

// func NewYeeBulb(host string) *YeeBulb {
// 	return &YeeBulb{
// 		host: host,
// 		port: 55443,
// 	}
// }

// func (y *YeeBulb) Power(on bool) {
// 	value := "on"
// 	if !on {
// 		value = "off"
// 	}
// 	y.sendMessage("power", []Any{value, "smooth", 500})
// }

// func (y *YeeBulb) Brightness(int value) {
// 	if value < 0 {
// 		value = 0
// 	} else if value > 100 {
// 		value = 100
// 	}
// 	y.sendMessage("set_bright", []Any{value, "smooth", 500})
// }

// func (y *YeeBulb) RGB(r, g, b uint8) {
// 	color := ((r & 0xFF) << 16) | ((g & 0xFF) << 8) | (b & 0xFF)
// 	y.sendMessage("set_rgb", []Any{color, "smooth", "500"})
// }

// func (y *YeeBulb) sendMessage(method string, params []Any) {
// 	//   def build_msg(id:, method:, params: [])
// 	//     msg = JSON.dump({
// 	//       id: id,
// 	//       method: method,
// 	//       params: params
// 	//     })
// 	//     "#{msg}\r\n"
// 	//   end

// 	//     TCPSocket.open(host, port) do |tcp|
// 	//       tcp.puts(build_msg(id: generate_id(), method: method, params: params))
// 	//       json = JSON.parse(tcp.gets)
// 	//       Response.new(json["id"], json["result"]&.first == "ok")
// 	//     end
// }

// func (y *YeeBulb) generateId() uint64 {
// 	// return Time.now.to_i
// }

// // a = Bulb.new("192.168.1.50")
// // b = Bulb.new("192.168.1.51")

// // puts a.rgb(100, 255, 255)
// // puts a.brightness(1)
