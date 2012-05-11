// Package xvmc is the X client API for the XVideo-MotionCompensation extension.
package xvmc

/*
	This file was generated by xvmc.xml on May 10 2012 11:56:20pm EDT.
	This file is automatically generated. Edit at your peril!
*/

import (
	"github.com/BurntSushi/xgb"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgb/xv"
)

// Init must be called before using the XVideo-MotionCompensation extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 25, "XVideo-MotionCompensation").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named XVideo-MotionCompensation could be found on on the server.")
	}

	xgb.ExtLock.Lock()
	c.Extensions["XVideo-MotionCompensation"] = reply.MajorOpcode
	for evNum, fun := range xgb.NewExtEventFuncs["XVideo-MotionCompensation"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["XVideo-MotionCompensation"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	xgb.ExtLock.Unlock()

	return nil
}

func init() {
	xgb.NewExtEventFuncs["XVideo-MotionCompensation"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["XVideo-MotionCompensation"] = make(map[int]xgb.NewErrorFun)
}

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Card32'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Int8'

type Context uint32

func NewContextId(c *xgb.Conn) (Context, error) {
	id, err := c.NewId()
	if err != nil {
		return 0, err
	}
	return Context(id), nil
}

type Surface uint32

func NewSurfaceId(c *xgb.Conn) (Surface, error) {
	id, err := c.NewId()
	if err != nil {
		return 0, err
	}
	return Surface(id), nil
}

type Subpicture uint32

func NewSubpictureId(c *xgb.Conn) (Subpicture, error) {
	id, err := c.NewId()
	if err != nil {
		return 0, err
	}
	return Subpicture(id), nil
}

type SurfaceInfo struct {
	Id                  Surface
	ChromaFormat        uint16
	Pad0                uint16
	MaxWidth            uint16
	MaxHeight           uint16
	SubpictureMaxWidth  uint16
	SubpictureMaxHeight uint16
	McType              uint32
	Flags               uint32
}

// SurfaceInfoRead reads a byte slice into a SurfaceInfo value.
func SurfaceInfoRead(buf []byte, v *SurfaceInfo) int {
	b := 0

	v.Id = Surface(xgb.Get32(buf[b:]))
	b += 4

	v.ChromaFormat = xgb.Get16(buf[b:])
	b += 2

	v.Pad0 = xgb.Get16(buf[b:])
	b += 2

	v.MaxWidth = xgb.Get16(buf[b:])
	b += 2

	v.MaxHeight = xgb.Get16(buf[b:])
	b += 2

	v.SubpictureMaxWidth = xgb.Get16(buf[b:])
	b += 2

	v.SubpictureMaxHeight = xgb.Get16(buf[b:])
	b += 2

	v.McType = xgb.Get32(buf[b:])
	b += 4

	v.Flags = xgb.Get32(buf[b:])
	b += 4

	return b
}

// SurfaceInfoReadList reads a byte slice into a list of SurfaceInfo values.
func SurfaceInfoReadList(buf []byte, dest []SurfaceInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = SurfaceInfo{}
		b += SurfaceInfoRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a SurfaceInfo value to a byte slice.
func (v SurfaceInfo) Bytes() []byte {
	buf := make([]byte, 24)
	b := 0

	xgb.Put32(buf[b:], uint32(v.Id))
	b += 4

	xgb.Put16(buf[b:], v.ChromaFormat)
	b += 2

	xgb.Put16(buf[b:], v.Pad0)
	b += 2

	xgb.Put16(buf[b:], v.MaxWidth)
	b += 2

	xgb.Put16(buf[b:], v.MaxHeight)
	b += 2

	xgb.Put16(buf[b:], v.SubpictureMaxWidth)
	b += 2

	xgb.Put16(buf[b:], v.SubpictureMaxHeight)
	b += 2

	xgb.Put32(buf[b:], v.McType)
	b += 4

	xgb.Put32(buf[b:], v.Flags)
	b += 4

	return buf
}

// SurfaceInfoListBytes writes a list of %s(MISSING) values to a byte slice.
func SurfaceInfoListBytes(buf []byte, list []SurfaceInfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}
	return b
}

// QueryVersionCookie is a cookie used only for QueryVersion requests.
type QueryVersionCookie struct {
	*xgb.Cookie
}

// QueryVersion sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryVersionCookie.Reply()
func QueryVersion(c *xgb.Conn) QueryVersionCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryVersionRequest(c), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryVersionUnchecked(c *xgb.Conn) QueryVersionCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryVersionRequest(c), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Major uint32
	Minor uint32
}

// Reply blocks and returns the reply data for a QueryVersion request.
func (cook QueryVersionCookie) Reply() (*QueryVersionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryVersionReply(buf), nil
}

// queryVersionReply reads a byte slice into a QueryVersionReply value.
func queryVersionReply(buf []byte) *QueryVersionReply {
	v := new(QueryVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.Major = xgb.Get32(buf[b:])
	b += 4

	v.Minor = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// ListSurfaceTypesCookie is a cookie used only for ListSurfaceTypes requests.
type ListSurfaceTypesCookie struct {
	*xgb.Cookie
}

// ListSurfaceTypes sends a checked request.
// If an error occurs, it will be returned with the reply by calling ListSurfaceTypesCookie.Reply()
func ListSurfaceTypes(c *xgb.Conn, PortId xv.Port) ListSurfaceTypesCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(listSurfaceTypesRequest(c, PortId), cookie)
	return ListSurfaceTypesCookie{cookie}
}

// ListSurfaceTypesUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func ListSurfaceTypesUnchecked(c *xgb.Conn, PortId xv.Port) ListSurfaceTypesCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(listSurfaceTypesRequest(c, PortId), cookie)
	return ListSurfaceTypesCookie{cookie}
}

// ListSurfaceTypesReply represents the data returned from a ListSurfaceTypes request.
type ListSurfaceTypesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Num uint32
	// padding: 20 bytes
	Surfaces []SurfaceInfo // size: xgb.Pad((int(Num) * 24))
}

// Reply blocks and returns the reply data for a ListSurfaceTypes request.
func (cook ListSurfaceTypesCookie) Reply() (*ListSurfaceTypesReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return listSurfaceTypesReply(buf), nil
}

// listSurfaceTypesReply reads a byte slice into a ListSurfaceTypesReply value.
func listSurfaceTypesReply(buf []byte) *ListSurfaceTypesReply {
	v := new(ListSurfaceTypesReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.Num = xgb.Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.Surfaces = make([]SurfaceInfo, v.Num)
	b += SurfaceInfoReadList(buf[b:], v.Surfaces)

	return v
}

// Write request to wire for ListSurfaceTypes
// listSurfaceTypesRequest writes a ListSurfaceTypes request to a byte slice.
func listSurfaceTypesRequest(c *xgb.Conn, PortId xv.Port) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(PortId))
	b += 4

	return buf
}

// CreateContextCookie is a cookie used only for CreateContext requests.
type CreateContextCookie struct {
	*xgb.Cookie
}

// CreateContext sends a checked request.
// If an error occurs, it will be returned with the reply by calling CreateContextCookie.Reply()
func CreateContext(c *xgb.Conn, ContextId Context, PortId xv.Port, SurfaceId Surface, Width uint16, Height uint16, Flags uint32) CreateContextCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(createContextRequest(c, ContextId, PortId, SurfaceId, Width, Height, Flags), cookie)
	return CreateContextCookie{cookie}
}

// CreateContextUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func CreateContextUnchecked(c *xgb.Conn, ContextId Context, PortId xv.Port, SurfaceId Surface, Width uint16, Height uint16, Flags uint32) CreateContextCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(createContextRequest(c, ContextId, PortId, SurfaceId, Width, Height, Flags), cookie)
	return CreateContextCookie{cookie}
}

// CreateContextReply represents the data returned from a CreateContext request.
type CreateContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	WidthActual  uint16
	HeightActual uint16
	FlagsReturn  uint32
	// padding: 20 bytes
	PrivData []uint32 // size: xgb.Pad((int(Length) * 4))
}

// Reply blocks and returns the reply data for a CreateContext request.
func (cook CreateContextCookie) Reply() (*CreateContextReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return createContextReply(buf), nil
}

// createContextReply reads a byte slice into a CreateContextReply value.
func createContextReply(buf []byte) *CreateContextReply {
	v := new(CreateContextReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.WidthActual = xgb.Get16(buf[b:])
	b += 2

	v.HeightActual = xgb.Get16(buf[b:])
	b += 2

	v.FlagsReturn = xgb.Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.PrivData = make([]uint32, v.Length)
	for i := 0; i < int(v.Length); i++ {
		v.PrivData[i] = xgb.Get32(buf[b:])
		b += 4
	}
	b = xgb.Pad(b)

	return v
}

// Write request to wire for CreateContext
// createContextRequest writes a CreateContext request to a byte slice.
func createContextRequest(c *xgb.Conn, ContextId Context, PortId xv.Port, SurfaceId Surface, Width uint16, Height uint16, Flags uint32) []byte {
	size := 24
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(ContextId))
	b += 4

	xgb.Put32(buf[b:], uint32(PortId))
	b += 4

	xgb.Put32(buf[b:], uint32(SurfaceId))
	b += 4

	xgb.Put16(buf[b:], Width)
	b += 2

	xgb.Put16(buf[b:], Height)
	b += 2

	xgb.Put32(buf[b:], Flags)
	b += 4

	return buf
}

// DestroyContextCookie is a cookie used only for DestroyContext requests.
type DestroyContextCookie struct {
	*xgb.Cookie
}

// DestroyContext sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func DestroyContext(c *xgb.Conn, ContextId Context) DestroyContextCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(destroyContextRequest(c, ContextId), cookie)
	return DestroyContextCookie{cookie}
}

// DestroyContextChecked sends a checked request.
// If an error occurs, it can be retrieved using DestroyContextCookie.Check()
func DestroyContextChecked(c *xgb.Conn, ContextId Context) DestroyContextCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(destroyContextRequest(c, ContextId), cookie)
	return DestroyContextCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook DestroyContextCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for DestroyContext
// destroyContextRequest writes a DestroyContext request to a byte slice.
func destroyContextRequest(c *xgb.Conn, ContextId Context) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(ContextId))
	b += 4

	return buf
}

// CreateSurfaceCookie is a cookie used only for CreateSurface requests.
type CreateSurfaceCookie struct {
	*xgb.Cookie
}

// CreateSurface sends a checked request.
// If an error occurs, it will be returned with the reply by calling CreateSurfaceCookie.Reply()
func CreateSurface(c *xgb.Conn, SurfaceId Surface, ContextId Context) CreateSurfaceCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(createSurfaceRequest(c, SurfaceId, ContextId), cookie)
	return CreateSurfaceCookie{cookie}
}

// CreateSurfaceUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func CreateSurfaceUnchecked(c *xgb.Conn, SurfaceId Surface, ContextId Context) CreateSurfaceCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(createSurfaceRequest(c, SurfaceId, ContextId), cookie)
	return CreateSurfaceCookie{cookie}
}

// CreateSurfaceReply represents the data returned from a CreateSurface request.
type CreateSurfaceReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	// padding: 24 bytes
	PrivData []uint32 // size: xgb.Pad((int(Length) * 4))
}

// Reply blocks and returns the reply data for a CreateSurface request.
func (cook CreateSurfaceCookie) Reply() (*CreateSurfaceReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return createSurfaceReply(buf), nil
}

// createSurfaceReply reads a byte slice into a CreateSurfaceReply value.
func createSurfaceReply(buf []byte) *CreateSurfaceReply {
	v := new(CreateSurfaceReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	b += 24 // padding

	v.PrivData = make([]uint32, v.Length)
	for i := 0; i < int(v.Length); i++ {
		v.PrivData[i] = xgb.Get32(buf[b:])
		b += 4
	}
	b = xgb.Pad(b)

	return v
}

// Write request to wire for CreateSurface
// createSurfaceRequest writes a CreateSurface request to a byte slice.
func createSurfaceRequest(c *xgb.Conn, SurfaceId Surface, ContextId Context) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(SurfaceId))
	b += 4

	xgb.Put32(buf[b:], uint32(ContextId))
	b += 4

	return buf
}

// DestroySurfaceCookie is a cookie used only for DestroySurface requests.
type DestroySurfaceCookie struct {
	*xgb.Cookie
}

// DestroySurface sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func DestroySurface(c *xgb.Conn, SurfaceId Surface) DestroySurfaceCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(destroySurfaceRequest(c, SurfaceId), cookie)
	return DestroySurfaceCookie{cookie}
}

// DestroySurfaceChecked sends a checked request.
// If an error occurs, it can be retrieved using DestroySurfaceCookie.Check()
func DestroySurfaceChecked(c *xgb.Conn, SurfaceId Surface) DestroySurfaceCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(destroySurfaceRequest(c, SurfaceId), cookie)
	return DestroySurfaceCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook DestroySurfaceCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for DestroySurface
// destroySurfaceRequest writes a DestroySurface request to a byte slice.
func destroySurfaceRequest(c *xgb.Conn, SurfaceId Surface) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(SurfaceId))
	b += 4

	return buf
}

// CreateSubpictureCookie is a cookie used only for CreateSubpicture requests.
type CreateSubpictureCookie struct {
	*xgb.Cookie
}

// CreateSubpicture sends a checked request.
// If an error occurs, it will be returned with the reply by calling CreateSubpictureCookie.Reply()
func CreateSubpicture(c *xgb.Conn, SubpictureId Subpicture, Context Context, XvimageId uint32, Width uint16, Height uint16) CreateSubpictureCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(createSubpictureRequest(c, SubpictureId, Context, XvimageId, Width, Height), cookie)
	return CreateSubpictureCookie{cookie}
}

// CreateSubpictureUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func CreateSubpictureUnchecked(c *xgb.Conn, SubpictureId Subpicture, Context Context, XvimageId uint32, Width uint16, Height uint16) CreateSubpictureCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(createSubpictureRequest(c, SubpictureId, Context, XvimageId, Width, Height), cookie)
	return CreateSubpictureCookie{cookie}
}

// CreateSubpictureReply represents the data returned from a CreateSubpicture request.
type CreateSubpictureReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	WidthActual       uint16
	HeightActual      uint16
	NumPaletteEntries uint16
	EntryBytes        uint16
	ComponentOrder    []byte // size: 4
	// padding: 12 bytes
	PrivData []uint32 // size: xgb.Pad((int(Length) * 4))
}

// Reply blocks and returns the reply data for a CreateSubpicture request.
func (cook CreateSubpictureCookie) Reply() (*CreateSubpictureReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return createSubpictureReply(buf), nil
}

// createSubpictureReply reads a byte slice into a CreateSubpictureReply value.
func createSubpictureReply(buf []byte) *CreateSubpictureReply {
	v := new(CreateSubpictureReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.WidthActual = xgb.Get16(buf[b:])
	b += 2

	v.HeightActual = xgb.Get16(buf[b:])
	b += 2

	v.NumPaletteEntries = xgb.Get16(buf[b:])
	b += 2

	v.EntryBytes = xgb.Get16(buf[b:])
	b += 2

	v.ComponentOrder = make([]byte, 4)
	copy(v.ComponentOrder[:4], buf[b:])
	b += xgb.Pad(int(4))

	b += 12 // padding

	v.PrivData = make([]uint32, v.Length)
	for i := 0; i < int(v.Length); i++ {
		v.PrivData[i] = xgb.Get32(buf[b:])
		b += 4
	}
	b = xgb.Pad(b)

	return v
}

// Write request to wire for CreateSubpicture
// createSubpictureRequest writes a CreateSubpicture request to a byte slice.
func createSubpictureRequest(c *xgb.Conn, SubpictureId Subpicture, Context Context, XvimageId uint32, Width uint16, Height uint16) []byte {
	size := 20
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(SubpictureId))
	b += 4

	xgb.Put32(buf[b:], uint32(Context))
	b += 4

	xgb.Put32(buf[b:], XvimageId)
	b += 4

	xgb.Put16(buf[b:], Width)
	b += 2

	xgb.Put16(buf[b:], Height)
	b += 2

	return buf
}

// DestroySubpictureCookie is a cookie used only for DestroySubpicture requests.
type DestroySubpictureCookie struct {
	*xgb.Cookie
}

// DestroySubpicture sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func DestroySubpicture(c *xgb.Conn, SubpictureId Subpicture) DestroySubpictureCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(destroySubpictureRequest(c, SubpictureId), cookie)
	return DestroySubpictureCookie{cookie}
}

// DestroySubpictureChecked sends a checked request.
// If an error occurs, it can be retrieved using DestroySubpictureCookie.Check()
func DestroySubpictureChecked(c *xgb.Conn, SubpictureId Subpicture) DestroySubpictureCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(destroySubpictureRequest(c, SubpictureId), cookie)
	return DestroySubpictureCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook DestroySubpictureCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for DestroySubpicture
// destroySubpictureRequest writes a DestroySubpicture request to a byte slice.
func destroySubpictureRequest(c *xgb.Conn, SubpictureId Subpicture) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(SubpictureId))
	b += 4

	return buf
}

// ListSubpictureTypesCookie is a cookie used only for ListSubpictureTypes requests.
type ListSubpictureTypesCookie struct {
	*xgb.Cookie
}

// ListSubpictureTypes sends a checked request.
// If an error occurs, it will be returned with the reply by calling ListSubpictureTypesCookie.Reply()
func ListSubpictureTypes(c *xgb.Conn, PortId xv.Port, SurfaceId Surface) ListSubpictureTypesCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(listSubpictureTypesRequest(c, PortId, SurfaceId), cookie)
	return ListSubpictureTypesCookie{cookie}
}

// ListSubpictureTypesUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func ListSubpictureTypesUnchecked(c *xgb.Conn, PortId xv.Port, SurfaceId Surface) ListSubpictureTypesCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(listSubpictureTypesRequest(c, PortId, SurfaceId), cookie)
	return ListSubpictureTypesCookie{cookie}
}

// ListSubpictureTypesReply represents the data returned from a ListSubpictureTypes request.
type ListSubpictureTypesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Num uint32
	// padding: 20 bytes
	Types []xv.ImageFormatInfo // size: xv.ImageFormatInfoListSize(Types)
}

// Reply blocks and returns the reply data for a ListSubpictureTypes request.
func (cook ListSubpictureTypesCookie) Reply() (*ListSubpictureTypesReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return listSubpictureTypesReply(buf), nil
}

// listSubpictureTypesReply reads a byte slice into a ListSubpictureTypesReply value.
func listSubpictureTypesReply(buf []byte) *ListSubpictureTypesReply {
	v := new(ListSubpictureTypesReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.Num = xgb.Get32(buf[b:])
	b += 4

	b += 20 // padding

	v.Types = make([]xv.ImageFormatInfo, v.Num)
	b += xv.ImageFormatInfoReadList(buf[b:], v.Types)

	return v
}

// Write request to wire for ListSubpictureTypes
// listSubpictureTypesRequest writes a ListSubpictureTypes request to a byte slice.
func listSubpictureTypesRequest(c *xgb.Conn, PortId xv.Port, SurfaceId Surface) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["XVIDEO-MOTIONCOMPENSATION"]
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(PortId))
	b += 4

	xgb.Put32(buf[b:], uint32(SurfaceId))
	b += 4

	return buf
}
