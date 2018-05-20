package uartwrap

***REMOVED***
	"io"
	"time"

	"github.com/tarm/serial"
***REMOVED***

type UART interface {
	io.ReadWriter
	io.Closer
	Open(***REMOVED*** error
	IsOpen(***REMOVED*** bool
***REMOVED***

type Config struct {
	Path        string        `mapstructure:"path"`
	Baudrate    uint32        `mapstructure:"baudrate"`
	ReadTimeout time.Duration `mapstructure:"read_timeout"`
***REMOVED***

type UARTDevice struct {
	Conf Config
	uart *serial.Port
***REMOVED***

func (u *UARTDevice***REMOVED*** Open(***REMOVED*** error {
	var err error
	u.uart, err = serial.OpenPort(&serial.Config{
		Name:        u.Conf.Path,
		Baud:        int(u.Conf.Baudrate***REMOVED***,
		ReadTimeout: time.Second * u.Conf.ReadTimeout,
***REMOVED******REMOVED***

	return err
***REMOVED***

func (u *UARTDevice***REMOVED*** IsOpen(***REMOVED*** bool {
	return u.uart != nil
***REMOVED***

func (u *UARTDevice***REMOVED*** Read(p []byte***REMOVED*** (int, error***REMOVED*** {
	return u.uart.Read(p***REMOVED***
***REMOVED***

func (u *UARTDevice***REMOVED*** Write(p []byte***REMOVED*** (int, error***REMOVED*** {
	return u.uart.Write(p***REMOVED***
***REMOVED***

func (u *UARTDevice***REMOVED*** Close(***REMOVED*** error {
	if err := u.uart.Close(***REMOVED***; err != nil {
		return err
***REMOVED***
	u.uart = nil
	return nil
***REMOVED***