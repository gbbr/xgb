// Package record is the X client API for the RECORD extension.
package record

/*
	This file was generated by record.xml on May 10 2012 11:56:19pm EDT.
	This file is automatically generated. Edit at your peril!
*/

import (
	"github.com/BurntSushi/xgb"

	"github.com/BurntSushi/xgb/xproto"
)

// Init must be called before using the RECORD extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 6, "RECORD").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named RECORD could be found on on the server.")
	}

	xgb.ExtLock.Lock()
	c.Extensions["RECORD"] = reply.MajorOpcode
	for evNum, fun := range xgb.NewExtEventFuncs["RECORD"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["RECORD"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	xgb.ExtLock.Unlock()

	return nil
}

func init() {
	xgb.NewExtEventFuncs["RECORD"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["RECORD"] = make(map[int]xgb.NewErrorFun)
}

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Card32'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Bool'

const (
	HTypeFromServerTime     = 1
	HTypeFromClientTime     = 2
	HTypeFromClientSequence = 4
)

const (
	CsCurrentClients = 1
	CsFutureClients  = 2
	CsAllClients     = 3
)

type Context uint32

func NewContextId(c *xgb.Conn) (Context, error) {
	id, err := c.NewId()
	if err != nil {
		return 0, err
	}
	return Context(id), nil
}

type ElementHeader byte

type ClientSpec uint32

type Range8 struct {
	First byte
	Last  byte
}

// Range8Read reads a byte slice into a Range8 value.
func Range8Read(buf []byte, v *Range8) int {
	b := 0

	v.First = buf[b]
	b += 1

	v.Last = buf[b]
	b += 1

	return b
}

// Range8ReadList reads a byte slice into a list of Range8 values.
func Range8ReadList(buf []byte, dest []Range8) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Range8{}
		b += Range8Read(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a Range8 value to a byte slice.
func (v Range8) Bytes() []byte {
	buf := make([]byte, 2)
	b := 0

	buf[b] = v.First
	b += 1

	buf[b] = v.Last
	b += 1

	return buf
}

// Range8ListBytes writes a list of %s(MISSING) values to a byte slice.
func Range8ListBytes(buf []byte, list []Range8) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}
	return b
}

type Range16 struct {
	First uint16
	Last  uint16
}

// Range16Read reads a byte slice into a Range16 value.
func Range16Read(buf []byte, v *Range16) int {
	b := 0

	v.First = xgb.Get16(buf[b:])
	b += 2

	v.Last = xgb.Get16(buf[b:])
	b += 2

	return b
}

// Range16ReadList reads a byte slice into a list of Range16 values.
func Range16ReadList(buf []byte, dest []Range16) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Range16{}
		b += Range16Read(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a Range16 value to a byte slice.
func (v Range16) Bytes() []byte {
	buf := make([]byte, 4)
	b := 0

	xgb.Put16(buf[b:], v.First)
	b += 2

	xgb.Put16(buf[b:], v.Last)
	b += 2

	return buf
}

// Range16ListBytes writes a list of %s(MISSING) values to a byte slice.
func Range16ListBytes(buf []byte, list []Range16) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}
	return b
}

type ExtRange struct {
	Major Range8
	Minor Range16
}

// ExtRangeRead reads a byte slice into a ExtRange value.
func ExtRangeRead(buf []byte, v *ExtRange) int {
	b := 0

	v.Major = Range8{}
	b += Range8Read(buf[b:], &v.Major)

	v.Minor = Range16{}
	b += Range16Read(buf[b:], &v.Minor)

	return b
}

// ExtRangeReadList reads a byte slice into a list of ExtRange values.
func ExtRangeReadList(buf []byte, dest []ExtRange) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ExtRange{}
		b += ExtRangeRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a ExtRange value to a byte slice.
func (v ExtRange) Bytes() []byte {
	buf := make([]byte, 6)
	b := 0

	{
		structBytes := v.Major.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}

	{
		structBytes := v.Minor.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}

	return buf
}

// ExtRangeListBytes writes a list of %s(MISSING) values to a byte slice.
func ExtRangeListBytes(buf []byte, list []ExtRange) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}
	return b
}

type Range struct {
	CoreRequests    Range8
	CoreReplies     Range8
	ExtRequests     ExtRange
	ExtReplies      ExtRange
	DeliveredEvents Range8
	DeviceEvents    Range8
	Errors          Range8
	ClientStarted   bool
	ClientDied      bool
}

// RangeRead reads a byte slice into a Range value.
func RangeRead(buf []byte, v *Range) int {
	b := 0

	v.CoreRequests = Range8{}
	b += Range8Read(buf[b:], &v.CoreRequests)

	v.CoreReplies = Range8{}
	b += Range8Read(buf[b:], &v.CoreReplies)

	v.ExtRequests = ExtRange{}
	b += ExtRangeRead(buf[b:], &v.ExtRequests)

	v.ExtReplies = ExtRange{}
	b += ExtRangeRead(buf[b:], &v.ExtReplies)

	v.DeliveredEvents = Range8{}
	b += Range8Read(buf[b:], &v.DeliveredEvents)

	v.DeviceEvents = Range8{}
	b += Range8Read(buf[b:], &v.DeviceEvents)

	v.Errors = Range8{}
	b += Range8Read(buf[b:], &v.Errors)

	if buf[b] == 1 {
		v.ClientStarted = true
	} else {
		v.ClientStarted = false
	}
	b += 1

	if buf[b] == 1 {
		v.ClientDied = true
	} else {
		v.ClientDied = false
	}
	b += 1

	return b
}

// RangeReadList reads a byte slice into a list of Range values.
func RangeReadList(buf []byte, dest []Range) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Range{}
		b += RangeRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a Range value to a byte slice.
func (v Range) Bytes() []byte {
	buf := make([]byte, 24)
	b := 0

	{
		structBytes := v.CoreRequests.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}

	{
		structBytes := v.CoreReplies.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}

	{
		structBytes := v.ExtRequests.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}

	{
		structBytes := v.ExtReplies.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}

	{
		structBytes := v.DeliveredEvents.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}

	{
		structBytes := v.DeviceEvents.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}

	{
		structBytes := v.Errors.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}

	if v.ClientStarted {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	if v.ClientDied {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	return buf
}

// RangeListBytes writes a list of %s(MISSING) values to a byte slice.
func RangeListBytes(buf []byte, list []Range) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}
	return b
}

type ClientInfo struct {
	ClientResource ClientSpec
	NumRanges      uint32
	Ranges         []Range // size: xgb.Pad((int(NumRanges) * 24))
}

// ClientInfoRead reads a byte slice into a ClientInfo value.
func ClientInfoRead(buf []byte, v *ClientInfo) int {
	b := 0

	v.ClientResource = ClientSpec(xgb.Get32(buf[b:]))
	b += 4

	v.NumRanges = xgb.Get32(buf[b:])
	b += 4

	v.Ranges = make([]Range, v.NumRanges)
	b += RangeReadList(buf[b:], v.Ranges)

	return b
}

// ClientInfoReadList reads a byte slice into a list of ClientInfo values.
func ClientInfoReadList(buf []byte, dest []ClientInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ClientInfo{}
		b += ClientInfoRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Bytes writes a ClientInfo value to a byte slice.
func (v ClientInfo) Bytes() []byte {
	buf := make([]byte, (8 + xgb.Pad((int(v.NumRanges) * 24))))
	b := 0

	xgb.Put32(buf[b:], uint32(v.ClientResource))
	b += 4

	xgb.Put32(buf[b:], v.NumRanges)
	b += 4

	b += RangeListBytes(buf[b:], v.Ranges)

	return buf
}

// ClientInfoListBytes writes a list of %s(MISSING) values to a byte slice.
func ClientInfoListBytes(buf []byte, list []ClientInfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}
	return b
}

// ClientInfoListSize computes the size (bytes) of a list of ClientInfo values.
func ClientInfoListSize(list []ClientInfo) int {
	size := 0
	for _, item := range list {
		size += (8 + xgb.Pad((int(item.NumRanges) * 24)))
	}
	return size
}

// BadBadContext is the error number for a BadBadContext.
const BadBadContext = 0

type BadContextError struct {
	Sequence      uint16
	NiceName      string
	InvalidRecord uint32
}

// BadContextErrorNew constructs a BadContextError value that implements xgb.Error from a byte slice.
func BadContextErrorNew(buf []byte) xgb.Error {
	v := BadContextError{}
	v.NiceName = "BadContext"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.InvalidRecord = xgb.Get32(buf[b:])
	b += 4

	return v
}

// SequenceId returns the sequence id attached to the BadBadContext error.
// This is mostly used internally.
func (err BadContextError) SequenceId() uint16 {
	return err.Sequence
}

func (err BadContextError) BadId() uint32 {
	return 0
}

func (err BadContextError) Error() string {
	fieldVals := make([]string, 0, 1)
	fieldVals = append(fieldVals, "NiceName: "+err.NiceName)
	fieldVals = append(fieldVals, xgb.Sprintf("Sequence: %d", err.Sequence))
	fieldVals = append(fieldVals, xgb.Sprintf("InvalidRecord: %d", err.InvalidRecord))
	return "BadBadContext {" + xgb.StringsJoin(fieldVals, ", ") + "}"
}

func init() {
	xgb.NewExtErrorFuncs["RECORD"][0] = BadContextErrorNew
}

// QueryVersionCookie is a cookie used only for QueryVersion requests.
type QueryVersionCookie struct {
	*xgb.Cookie
}

// QueryVersion sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryVersionCookie.Reply()
func QueryVersion(c *xgb.Conn, MajorVersion uint16, MinorVersion uint16) QueryVersionCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryVersionRequest(c, MajorVersion, MinorVersion), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryVersionUnchecked(c *xgb.Conn, MajorVersion uint16, MinorVersion uint16) QueryVersionCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryVersionRequest(c, MajorVersion, MinorVersion), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	MajorVersion uint16
	MinorVersion uint16
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

	v.MajorVersion = xgb.Get16(buf[b:])
	b += 2

	v.MinorVersion = xgb.Get16(buf[b:])
	b += 2

	return v
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(c *xgb.Conn, MajorVersion uint16, MinorVersion uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["RECORD"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put16(buf[b:], MajorVersion)
	b += 2

	xgb.Put16(buf[b:], MinorVersion)
	b += 2

	return buf
}

// CreateContextCookie is a cookie used only for CreateContext requests.
type CreateContextCookie struct {
	*xgb.Cookie
}

// CreateContext sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func CreateContext(c *xgb.Conn, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) CreateContextCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(createContextRequest(c, Context, ElementHeader, NumClientSpecs, NumRanges, ClientSpecs, Ranges), cookie)
	return CreateContextCookie{cookie}
}

// CreateContextChecked sends a checked request.
// If an error occurs, it can be retrieved using CreateContextCookie.Check()
func CreateContextChecked(c *xgb.Conn, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) CreateContextCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(createContextRequest(c, Context, ElementHeader, NumClientSpecs, NumRanges, ClientSpecs, Ranges), cookie)
	return CreateContextCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook CreateContextCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for CreateContext
// createContextRequest writes a CreateContext request to a byte slice.
func createContextRequest(c *xgb.Conn, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) []byte {
	size := xgb.Pad(((20 + xgb.Pad((int(NumClientSpecs) * 4))) + xgb.Pad((int(NumRanges) * 24))))
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["RECORD"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Context))
	b += 4

	buf[b] = byte(ElementHeader)
	b += 1

	b += 3 // padding

	xgb.Put32(buf[b:], NumClientSpecs)
	b += 4

	xgb.Put32(buf[b:], NumRanges)
	b += 4

	for i := 0; i < int(NumClientSpecs); i++ {
		xgb.Put32(buf[b:], uint32(ClientSpecs[i]))
		b += 4
	}
	b = xgb.Pad(b)

	b += RangeListBytes(buf[b:], Ranges)

	return buf
}

// RegisterClientsCookie is a cookie used only for RegisterClients requests.
type RegisterClientsCookie struct {
	*xgb.Cookie
}

// RegisterClients sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func RegisterClients(c *xgb.Conn, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) RegisterClientsCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(registerClientsRequest(c, Context, ElementHeader, NumClientSpecs, NumRanges, ClientSpecs, Ranges), cookie)
	return RegisterClientsCookie{cookie}
}

// RegisterClientsChecked sends a checked request.
// If an error occurs, it can be retrieved using RegisterClientsCookie.Check()
func RegisterClientsChecked(c *xgb.Conn, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) RegisterClientsCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(registerClientsRequest(c, Context, ElementHeader, NumClientSpecs, NumRanges, ClientSpecs, Ranges), cookie)
	return RegisterClientsCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook RegisterClientsCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for RegisterClients
// registerClientsRequest writes a RegisterClients request to a byte slice.
func registerClientsRequest(c *xgb.Conn, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) []byte {
	size := xgb.Pad(((20 + xgb.Pad((int(NumClientSpecs) * 4))) + xgb.Pad((int(NumRanges) * 24))))
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["RECORD"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Context))
	b += 4

	buf[b] = byte(ElementHeader)
	b += 1

	b += 3 // padding

	xgb.Put32(buf[b:], NumClientSpecs)
	b += 4

	xgb.Put32(buf[b:], NumRanges)
	b += 4

	for i := 0; i < int(NumClientSpecs); i++ {
		xgb.Put32(buf[b:], uint32(ClientSpecs[i]))
		b += 4
	}
	b = xgb.Pad(b)

	b += RangeListBytes(buf[b:], Ranges)

	return buf
}

// UnregisterClientsCookie is a cookie used only for UnregisterClients requests.
type UnregisterClientsCookie struct {
	*xgb.Cookie
}

// UnregisterClients sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func UnregisterClients(c *xgb.Conn, Context Context, NumClientSpecs uint32, ClientSpecs []ClientSpec) UnregisterClientsCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(unregisterClientsRequest(c, Context, NumClientSpecs, ClientSpecs), cookie)
	return UnregisterClientsCookie{cookie}
}

// UnregisterClientsChecked sends a checked request.
// If an error occurs, it can be retrieved using UnregisterClientsCookie.Check()
func UnregisterClientsChecked(c *xgb.Conn, Context Context, NumClientSpecs uint32, ClientSpecs []ClientSpec) UnregisterClientsCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(unregisterClientsRequest(c, Context, NumClientSpecs, ClientSpecs), cookie)
	return UnregisterClientsCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook UnregisterClientsCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for UnregisterClients
// unregisterClientsRequest writes a UnregisterClients request to a byte slice.
func unregisterClientsRequest(c *xgb.Conn, Context Context, NumClientSpecs uint32, ClientSpecs []ClientSpec) []byte {
	size := xgb.Pad((12 + xgb.Pad((int(NumClientSpecs) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["RECORD"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Context))
	b += 4

	xgb.Put32(buf[b:], NumClientSpecs)
	b += 4

	for i := 0; i < int(NumClientSpecs); i++ {
		xgb.Put32(buf[b:], uint32(ClientSpecs[i]))
		b += 4
	}
	b = xgb.Pad(b)

	return buf
}

// GetContextCookie is a cookie used only for GetContext requests.
type GetContextCookie struct {
	*xgb.Cookie
}

// GetContext sends a checked request.
// If an error occurs, it will be returned with the reply by calling GetContextCookie.Reply()
func GetContext(c *xgb.Conn, Context Context) GetContextCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(getContextRequest(c, Context), cookie)
	return GetContextCookie{cookie}
}

// GetContextUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func GetContextUnchecked(c *xgb.Conn, Context Context) GetContextCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(getContextRequest(c, Context), cookie)
	return GetContextCookie{cookie}
}

// GetContextReply represents the data returned from a GetContext request.
type GetContextReply struct {
	Sequence      uint16 // sequence number of the request for this reply
	Length        uint32 // number of bytes in this reply
	Enabled       bool
	ElementHeader ElementHeader
	// padding: 3 bytes
	NumInterceptedClients uint32
	// padding: 16 bytes
	InterceptedClients []ClientInfo // size: ClientInfoListSize(InterceptedClients)
}

// Reply blocks and returns the reply data for a GetContext request.
func (cook GetContextCookie) Reply() (*GetContextReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getContextReply(buf), nil
}

// getContextReply reads a byte slice into a GetContextReply value.
func getContextReply(buf []byte) *GetContextReply {
	v := new(GetContextReply)
	b := 1 // skip reply determinant

	if buf[b] == 1 {
		v.Enabled = true
	} else {
		v.Enabled = false
	}
	b += 1

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.ElementHeader = ElementHeader(buf[b])
	b += 1

	b += 3 // padding

	v.NumInterceptedClients = xgb.Get32(buf[b:])
	b += 4

	b += 16 // padding

	v.InterceptedClients = make([]ClientInfo, v.NumInterceptedClients)
	b += ClientInfoReadList(buf[b:], v.InterceptedClients)

	return v
}

// Write request to wire for GetContext
// getContextRequest writes a GetContext request to a byte slice.
func getContextRequest(c *xgb.Conn, Context Context) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["RECORD"]
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Context))
	b += 4

	return buf
}

// EnableContextCookie is a cookie used only for EnableContext requests.
type EnableContextCookie struct {
	*xgb.Cookie
}

// EnableContext sends a checked request.
// If an error occurs, it will be returned with the reply by calling EnableContextCookie.Reply()
func EnableContext(c *xgb.Conn, Context Context) EnableContextCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(enableContextRequest(c, Context), cookie)
	return EnableContextCookie{cookie}
}

// EnableContextUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func EnableContextUnchecked(c *xgb.Conn, Context Context) EnableContextCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(enableContextRequest(c, Context), cookie)
	return EnableContextCookie{cookie}
}

// EnableContextReply represents the data returned from a EnableContext request.
type EnableContextReply struct {
	Sequence      uint16 // sequence number of the request for this reply
	Length        uint32 // number of bytes in this reply
	Category      byte
	ElementHeader ElementHeader
	ClientSwapped bool
	// padding: 2 bytes
	XidBase        uint32
	ServerTime     uint32
	RecSequenceNum uint32
	// padding: 8 bytes
	Data []byte // size: xgb.Pad(((int(Length) * 4) * 1))
}

// Reply blocks and returns the reply data for a EnableContext request.
func (cook EnableContextCookie) Reply() (*EnableContextReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return enableContextReply(buf), nil
}

// enableContextReply reads a byte slice into a EnableContextReply value.
func enableContextReply(buf []byte) *EnableContextReply {
	v := new(EnableContextReply)
	b := 1 // skip reply determinant

	v.Category = buf[b]
	b += 1

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.ElementHeader = ElementHeader(buf[b])
	b += 1

	if buf[b] == 1 {
		v.ClientSwapped = true
	} else {
		v.ClientSwapped = false
	}
	b += 1

	b += 2 // padding

	v.XidBase = xgb.Get32(buf[b:])
	b += 4

	v.ServerTime = xgb.Get32(buf[b:])
	b += 4

	v.RecSequenceNum = xgb.Get32(buf[b:])
	b += 4

	b += 8 // padding

	v.Data = make([]byte, (int(v.Length) * 4))
	copy(v.Data[:(int(v.Length)*4)], buf[b:])
	b += xgb.Pad(int((int(v.Length) * 4)))

	return v
}

// Write request to wire for EnableContext
// enableContextRequest writes a EnableContext request to a byte slice.
func enableContextRequest(c *xgb.Conn, Context Context) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["RECORD"]
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Context))
	b += 4

	return buf
}

// DisableContextCookie is a cookie used only for DisableContext requests.
type DisableContextCookie struct {
	*xgb.Cookie
}

// DisableContext sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func DisableContext(c *xgb.Conn, Context Context) DisableContextCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(disableContextRequest(c, Context), cookie)
	return DisableContextCookie{cookie}
}

// DisableContextChecked sends a checked request.
// If an error occurs, it can be retrieved using DisableContextCookie.Check()
func DisableContextChecked(c *xgb.Conn, Context Context) DisableContextCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(disableContextRequest(c, Context), cookie)
	return DisableContextCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook DisableContextCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for DisableContext
// disableContextRequest writes a DisableContext request to a byte slice.
func disableContextRequest(c *xgb.Conn, Context Context) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["RECORD"]
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Context))
	b += 4

	return buf
}

// FreeContextCookie is a cookie used only for FreeContext requests.
type FreeContextCookie struct {
	*xgb.Cookie
}

// FreeContext sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func FreeContext(c *xgb.Conn, Context Context) FreeContextCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(freeContextRequest(c, Context), cookie)
	return FreeContextCookie{cookie}
}

// FreeContextChecked sends a checked request.
// If an error occurs, it can be retrieved using FreeContextCookie.Check()
func FreeContextChecked(c *xgb.Conn, Context Context) FreeContextCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(freeContextRequest(c, Context), cookie)
	return FreeContextCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook FreeContextCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for FreeContext
// freeContextRequest writes a FreeContext request to a byte slice.
func freeContextRequest(c *xgb.Conn, Context Context) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["RECORD"]
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Context))
	b += 4

	return buf
}
