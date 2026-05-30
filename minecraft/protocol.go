package minecraft

import (
	"io"

	"github.com/Yeah114/gophertunnel/minecraft/protocol"
	"github.com/Yeah114/gophertunnel/minecraft/protocol/packet"
)

// Protocol 表示通过网络通信时使用的 Minecraft 协议。它由一组唯一的数据包组成，这组数据包
// 可能在任意版本中发生变化。Protocol 还负责在最新协议与 Protocol 指定协议之间转换数据包。
//
// Protocol represents the Minecraft protocol used to communicate over network. It comprises a unique set of packets
// that may be changed in any version.
// Protocol specifically handles the conversion of packets between the most recent protocol (as in the
// minecraft/protocol package) and the protocol as specified in Protocol.
type Protocol interface {
	// ID 返回协议的唯一 ID。通常每发布一个新的 Minecraft 版本，该 ID 都会上升。
	//
	// ID returns the unique ID of the Protocol. It generally goes up for every new Minecraft version released.
	ID() int32
	// Ver 返回该协议关联的 Minecraft 版本，例如 "1.18.10"。
	//
	// Ver returns the Minecraft version associated with this Protocol, such as "1.18.10".
	Ver() string
	// Packets 返回为该 Protocol 注册了所有数据包的 packet.Pool。它用于按数据包 ID 查找数据包。
	// 如果 listener 为 true，返回的池应面向 Listener 创建，也就是说只允许客户端可能发送的数据包。
	//
	// Packets returns a packet.Pool with all packets registered for this
	// Protocol. It is used to lookup packets by a packet ID. If listener is set
	// to true, the pool should be created for a Listener. This means that only
	// packets that may be sent by a client should be allowed.
	Packets(listener bool) packet.Pool
	// NewReader 返回实现该 Protocol 所需类型读取操作的 protocol.IO。
	//
	// NewReader returns a protocol.IO that implements reading operations for reading types
	// that are used for this Protocol.
	NewReader(r ByteReader, shieldID int32, enableLimits bool) protocol.IO
	// NewWriter 返回实现该 Protocol 所需类型写入操作的 protocol.IO。
	//
	// NewWriter returns a protocol.IO that implements writing operations for writing types
	// that are used for this Protocol.
	NewWriter(w ByteWriter, shieldID int32) protocol.IO
	// ConvertToLatest 将从 Conn 对端取得的 packet.Packet 转换为最新协议中的 packet.Packet 切片。
	// 通过 Packets 取得的 packet.Pool 中，凡是与最新版本载荷或 ID 不一致的数据包，都必须在这里
	// 正确转换。若数据包在该版本中相对最新版本没有变化，则返回 pk。即使只改变 ID 也必须转换。
	//
	// ConvertToLatest converts a packet.Packet obtained from the other end of a Conn to a slice of packet.Packets from
	// the latest protocol. Any packet.Packet implementation in the packet.Pool obtained through a call to Packets that
	// is not identical to the most recent version of that packet.Packet must be converted to the most recent version of
	// that packet adequately in this function. ConvertToLatest returns pk if the packet.Packet was unchanged in this
	// version compared to the latest. Note that packets must also be converted if only their ID changes.
	ConvertToLatest(pk packet.Packet, conn *Conn) []packet.Packet
	// ConvertFromLatest 将最新协议的 packet.Packet 转换为该 Protocol 的 packet.Packet 切片。
	// 它应与 ConvertToLatest 对称：如果某个数据包相对最新协议的载荷或 ID 发生变化，就应转换为
	// Packets 返回的池中对应的正确数据包。
	//
	// ConvertFromLatest converts a packet.Packet of the most recent Protocol to a slice of packet.Packets of this
	// specific Protocol. ConvertFromLatest must be synonymous to ConvertToLatest, in that it should convert any
	// packet.Packet to the correct one from the packet.Pool returned through a call to Packets if its payload or ID was
	// changed in this Protocol compared to the latest one.
	ConvertFromLatest(pk packet.Packet, conn *Conn) []packet.Packet
}

// ByteReader 组合 io.Reader 和 io.ByteReader，供协议读取器使用。
//
// ByteReader combines io.Reader and io.ByteReader for protocol readers.
type ByteReader interface {
	io.Reader
	io.ByteReader
}

// ByteWriter 组合 io.Writer 和 io.ByteWriter，供协议写入器使用。
//
// ByteWriter combines io.Writer and io.ByteWriter for protocol writers.
type ByteWriter interface {
	io.Writer
	io.ByteWriter
}

// BedrockProtocol 表示 Bedrock Edition 的 Minecraft 协议实现。
//
// BedrockProtocol represents a Minecraft protocol implementation for the Bedrock Edition.
type BedrockProtocol struct {
	protocol.Profile
}

// NewBedrockProtocol 使用给定 protocol.Profile 创建新协议。返回的协议会使用该 Profile
// 注册的 Minecraft 版本和数据包池，并且不会转换任何数据包，因为它们已经是正确类型。
//
// NewBedrockProtocol creates a new Protocol with the given protocol.Profile. The returned Protocol will use the
// Minecraft version and packet pool registered for the given protocol.Profile, and will not convert any packets,
// as they are already of the right type.
func NewBedrockProtocol(profile protocol.Profile) BedrockProtocol {
	return BedrockProtocol{Profile: profile}
}

// Packets 返回为该 Protocol 注册了所有数据包的 packet.Pool。它用于按数据包 ID 查找数据包。
// 如果 listener 为 true，池会按 Listener 使用场景创建，只允许客户端可能发送的数据包。
//
// Packets returns a packet.Pool with all packets registered for this Protocol. It is used to lookup packets by a
// packet ID. If listener is set to true, the pool should be created for a Listener. This means that only packets
// that may be sent by a client should be allowed.
func (p BedrockProtocol) Packets(listener bool) packet.Pool {
	return packet.NewPool()
}

// NewReader 返回实现读取该协议所需类型的 protocol.IO。
//
// NewReader returns a protocol.IO that implements reading operations for reading types.
func (p BedrockProtocol) NewReader(r ByteReader, shieldID int32, enableLimits bool) protocol.IO {
	return protocol.NewReaderWithProfile(r, shieldID, enableLimits, p.Profile)
}

// NewWriter 返回实现写入该协议所需类型的 protocol.IO。
//
// NewWriter returns a protocol.IO that implements writing operations for writing types.
func (p BedrockProtocol) NewWriter(w ByteWriter, shieldID int32) protocol.IO {
	return protocol.NewWriterWithProfile(w, shieldID, p.Profile)
}

// ConvertToLatest 将从 Conn 对端取得的 packet.Packet 转换为最新协议的 packet.Packet 切片。
// 当前 BedrockProtocol 使用的包已经是正确类型，因此直接返回原数据包。
//
// ConvertToLatest 将从 Conn 对端取得的 packet.Packet 转换为最新协议的 packet.Packet 切片。
// 当前 BedrockProtocol 使用的包已经是正确类型，因此直接返回原数据包。
//
// ConvertToLatest converts a packet.Packet obtained from the other end of a Conn to a slice of packet.Packets from
// the latest protocol. Any packet.Packet implementation in the packet.Pool obtained through a call to Packets that
// is not identical to the most recent version of that packet.Packet must be converted to the most recent version of
// that packet adequately in this function. ConvertToLatest returns pk if the packet.Packet was unchanged in this
// version compared to the latest. Note that packets must also be converted if only their ID changes.
func (p BedrockProtocol) ConvertToLatest(pk packet.Packet, conn *Conn) []packet.Packet {
	return []packet.Packet{pk}
}

// ConvertFromLatest 将最新协议的 packet.Packet 转换为该具体协议的 packet.Packet 切片。
// 当前 BedrockProtocol 使用的包已经是正确类型，因此直接返回原数据包。
//
// ConvertFromLatest converts a packet.Packet of the most recent Protocol to a slice of packet.Packets of this
// specific Protocol. ConvertFromLatest must be synonymous to ConvertToLatest, in that it should convert any
// packet.Packet to the correct one from the packet.Pool returned through a call to Packets if its payload or ID
// was changed in this Protocol compared to the latest one.
func (p BedrockProtocol) ConvertFromLatest(pk packet.Packet, conn *Conn) []packet.Packet {
	return []packet.Packet{pk}
}

// BedrockProtocolPool 按协议 ID 保存 BedrockProtocol。
//
// BedrockProtocolPool stores BedrockProtocol values by protocol ID.
type BedrockProtocolPool map[int32]BedrockProtocol

var bedrockProtocolPool = make(BedrockProtocolPool)

// RegisterBedrockProtocol 注册 Bedrock 协议实现，使其可按协议 ID 查找。
//
// RegisterBedrockProtocol registers a Bedrock protocol implementation so it can
// be looked up by protocol ID.
func RegisterBedrockProtocol(bedrockProtocol BedrockProtocol) {
	bedrockProtocolPool[bedrockProtocol.ID()] = bedrockProtocol
}

// NewBedrockProtocolPool 返回当前已注册 Bedrock 协议实现的副本。
//
// NewBedrockProtocolPool returns a copy of all currently registered Bedrock
// protocol implementations.
func NewBedrockProtocolPool() BedrockProtocolPool {
	pool := make(BedrockProtocolPool)
	for protocol, info := range bedrockProtocolPool {
		pool[protocol] = info
	}

	return pool
}

// proto 是默认协议实现。它返回当前协议、版本和数据包池，并且不转换数据包，
// 因为这些数据包已经是正确类型。
//
// proto is the default Protocol implementation. It returns the current protocol, version and packet pool and does not
// convert any packets, as they are already of the right type.
type proto struct{}

// ID 返回当前协议 ID。
//
// ID returns the current protocol ID.
func (proto) ID() int32 { return protocol.CurrentProtocol }

// Ver 返回当前 Minecraft 版本。
//
// Ver returns the current Minecraft version.
func (p proto) Ver() string { return protocol.CurrentVersion }

// Packets 返回当前协议的数据包池。
//
// Packets returns the packet pool for the current protocol.
func (p proto) Packets(listener bool) packet.Pool {
	if listener {
		return packet.NewClientPool()
	}
	return packet.NewServerPool()
}

// NewReader 返回当前协议的读取器。
//
// NewReader returns a reader for the current protocol.
func (p proto) NewReader(r ByteReader, shieldID int32, enableLimits bool) protocol.IO {
	return protocol.NewReader(r, shieldID, enableLimits)
}

// NewWriter 返回当前协议的写入器。
//
// NewWriter returns a writer for the current protocol.
func (p proto) NewWriter(w ByteWriter, shieldID int32) protocol.IO {
	return protocol.NewWriter(w, shieldID)
}

// ConvertToLatest 返回原数据包，因为当前协议已经是最新协议。
//
// ConvertToLatest returns the original packet because the current protocol is
// already the latest protocol.
func (p proto) ConvertToLatest(pk packet.Packet, _ *Conn) []packet.Packet { return []packet.Packet{pk} }

// ConvertFromLatest 返回原数据包，因为当前协议已经是最新协议。
//
// ConvertFromLatest returns the original packet because the current protocol is
// already the latest protocol.
func (p proto) ConvertFromLatest(pk packet.Packet, _ *Conn) []packet.Packet {
	return []packet.Packet{pk}
}

// DefaultProtocol 是默认使用的协议实现。默认情况下，它使用当前协议、版本和数据包池，
// 并且不转换数据包，因为这些数据包已经是正确类型。
//
// DefaultProtocol is the Protocol implementation used by default. It uses the current protocol, version and packet
// pool and does not convert any packets, as they are already of the right type.
var DefaultProtocol = proto{}

func init() {
	for _, profile := range protocol.Profiles() {
		RegisterBedrockProtocol(NewBedrockProtocol(profile))
	}
}
