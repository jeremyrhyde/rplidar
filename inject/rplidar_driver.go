//nolint:golint
package inject

import (
	"go.viam.com/rplidar/gen"
)

type RPlidarDriver struct {
	gen.RPlidarDriver

	SwigcptrFunc            func() uintptr
	SwigIsRPlidarDriverFunc func()

	ConnectFunc     func(a ...interface{}) uint
	DisconnectFunc  func()
	IsConnectedFunc func() (_swig_ret bool)

	ResetFunc                    func(a ...interface{}) uint
	ClearNetSerialRxCacheFunc    func() (_swig_ret uint)
	GetAllSupportedScanModesFunc func(a ...interface{}) uint
	GetTypicalScanModeFunc       func(a ...interface{}) uint
	StartScanFunc                func(a ...interface{}) uint
	StartScanExpressFunc         func(a ...interface{}) uint
	GetHealthFunc                func(a ...interface{}) uint
	GetDeviceInfoFunc            func(a ...interface{}) uint
	GetSampleDuration_uSFunc     func(a ...interface{}) uint

	SetMotorPWMFunc       func(arg2 uint16) (_swig_ret uint)
	SetLidarSpinSpeedFunc func(a ...interface{}) uint
	StartMotorFunc        func() (_swig_ret uint)
	StopMotorFunc         func() (_swig_ret uint)

	CheckMotorCtrlSupportFunc     func(a ...interface{}) uint
	CheckIfTofLidarFunc           func(a ...interface{}) uint
	GetFrequencyFunc              func(a ...interface{}) uint
	StartScanNormalFunc           func(a ...interface{}) uint
	CheckExpressScanSupportedFunc func(a ...interface{}) uint
	StopFunc                      func(a ...interface{}) uint

	GrabScanDataFunc              func(a ...interface{}) uint
	GrabScanDataHqFunc            func(a ...interface{}) uint
	AscendScanDataFunc            func(a ...interface{}) uint
	GetScanDataWithIntervalFunc   func(arg2 gen.Rplidar_response_measurement_node_t, arg3 *int64) (_swig_ret uint)
	GetScanDataWithIntervalHqFunc func(arg2 gen.Rplidar_response_measurement_node_hq_t, arg3 *int64) (_swig_ret uint)

	SetX_chanDevFunc func(arg2 gen.ChannelDevice)
	GetX_chanDevFunc func() (_swig_ret gen.ChannelDevice)
}

func (driver *RPlidarDriver) Swigcptr() uintptr {
	if driver.SwigcptrFunc == nil {
		return driver.RPlidarDriver.Swigcptr()
	}
	return driver.SwigcptrFunc()
}

func (driver *RPlidarDriver) SwigIsRPlidarDriver() {
	if driver.SwigIsRPlidarDriverFunc == nil {
		driver.RPlidarDriver.SwigIsRPlidarDriver()
	}
	driver.SwigIsRPlidarDriverFunc()
}

func (driver *RPlidarDriver) Connect(a ...interface{}) uint {
	if driver.ConnectFunc == nil {
		return driver.RPlidarDriver.Connect(a)
	}
	return driver.ConnectFunc(a)
}

func (driver *RPlidarDriver) Disconnect() {
	if driver.DisconnectFunc == nil {
		driver.RPlidarDriver.Disconnect()
	}
	driver.DisconnectFunc()
}

func (driver *RPlidarDriver) IsConnected() bool {
	if driver.IsConnectedFunc == nil {
		return driver.RPlidarDriver.IsConnected()
	}
	return driver.IsConnectedFunc()
}

func (driver *RPlidarDriver) Reset(a ...interface{}) uint {
	if driver.ResetFunc == nil {
		return driver.RPlidarDriver.Reset(a)
	}
	return driver.ResetFunc(a)
}

func (driver *RPlidarDriver) ClearNetSerialRxCache() uint {
	if driver.ClearNetSerialRxCacheFunc == nil {
		return driver.RPlidarDriver.ClearNetSerialRxCache()
	}
	return driver.ClearNetSerialRxCacheFunc()
}

func (driver *RPlidarDriver) GetAllSupportedScanModes(a ...interface{}) uint {
	if driver.GetAllSupportedScanModesFunc == nil {
		return driver.RPlidarDriver.GetAllSupportedScanModes(a)
	}
	return driver.GetAllSupportedScanModesFunc(a)
}

func (driver *RPlidarDriver) GetTypicalScanMode(a ...interface{}) uint {
	if driver.GetTypicalScanModeFunc == nil {
		return driver.RPlidarDriver.GetTypicalScanMode(a)
	}
	return driver.GetTypicalScanModeFunc(a)
}

func (driver *RPlidarDriver) StartScan(a ...interface{}) uint {
	if driver.StartScanFunc == nil {
		return driver.RPlidarDriver.StartScan(a)
	}
	return driver.StartScanFunc(a)
}

func (driver *RPlidarDriver) StartScanExpress(a ...interface{}) uint {
	if driver.StartScanExpressFunc == nil {
		return driver.RPlidarDriver.StartScanExpress(a)
	}
	return driver.StartScanExpressFunc(a)
}

func (driver *RPlidarDriver) GetHealth(a ...interface{}) uint {
	if driver.GetHealthFunc == nil {
		return driver.RPlidarDriver.GetHealth(a)
	}
	return driver.GetHealthFunc(a)
}

func (driver *RPlidarDriver) GetDeviceInfo(a ...interface{}) uint {
	if driver.GetDeviceInfoFunc == nil {
		return driver.RPlidarDriver.GetDeviceInfo(a)
	}
	return driver.GetDeviceInfoFunc(a)
}

func (driver *RPlidarDriver) GetSampleDuration_uS(a ...interface{}) uint {
	if driver.GetSampleDuration_uSFunc == nil {
		return driver.RPlidarDriver.GetSampleDuration_uS(a)
	}
	return driver.GetSampleDuration_uSFunc(a)
}

func (driver *RPlidarDriver) SetMotorPWM(arg2 uint16) uint {
	if driver.SetMotorPWMFunc == nil {
		return driver.RPlidarDriver.SetMotorPWM(arg2)
	}
	return driver.SetMotorPWMFunc(arg2)
}

func (driver *RPlidarDriver) SetLidarSpinSpeed(a ...interface{}) uint {
	if driver.SetLidarSpinSpeedFunc == nil {
		return driver.RPlidarDriver.SetLidarSpinSpeed(a)
	}
	return driver.SetLidarSpinSpeedFunc(a)
}

func (driver *RPlidarDriver) StartMotor() uint {
	if driver.StartMotorFunc == nil {
		return driver.RPlidarDriver.StartMotor()
	}
	return driver.StartMotorFunc()
}

func (driver *RPlidarDriver) StopMotor() uint {
	if driver.StopMotorFunc == nil {
		return driver.RPlidarDriver.StopMotor()
	}
	return driver.StopMotorFunc()
}

func (driver *RPlidarDriver) CheckMotorCtrlSupport(a ...interface{}) uint {
	if driver.CheckMotorCtrlSupportFunc == nil {
		return driver.RPlidarDriver.CheckMotorCtrlSupport(a)
	}
	return driver.CheckMotorCtrlSupportFunc(a)
}

func (driver *RPlidarDriver) CheckIfTofLidar(a ...interface{}) uint {
	if driver.CheckIfTofLidarFunc == nil {
		return driver.RPlidarDriver.CheckIfTofLidar(a)
	}
	return driver.CheckIfTofLidarFunc(a)
}

func (driver *RPlidarDriver) GetFrequency(a ...interface{}) uint {
	if driver.GetFrequencyFunc == nil {
		return driver.RPlidarDriver.GetFrequency(a)
	}
	return driver.GetFrequencyFunc(a)
}

func (driver *RPlidarDriver) StartScanNormal(a ...interface{}) uint {
	if driver.StartScanNormalFunc == nil {
		return driver.RPlidarDriver.StartScanNormal(a)
	}
	return driver.StartScanNormalFunc(a)
}

func (driver *RPlidarDriver) CheckExpressScanSupported(a ...interface{}) uint {
	if driver.CheckExpressScanSupportedFunc == nil {
		return driver.RPlidarDriver.CheckExpressScanSupported(a)
	}
	return driver.CheckExpressScanSupportedFunc(a)
}

func (driver *RPlidarDriver) Stop(a ...interface{}) uint {
	if driver.StopFunc == nil {
		return driver.RPlidarDriver.Stop(a)
	}
	return driver.StopFunc(a)
}

func (driver *RPlidarDriver) GrabScanData(a ...interface{}) uint {
	if driver.GrabScanDataFunc == nil {
		return driver.RPlidarDriver.GrabScanData(a)
	}
	return driver.GrabScanDataFunc(a)
}

func (driver *RPlidarDriver) GrabScanDataHq(a ...interface{}) uint {
	if driver.GrabScanDataHqFunc == nil {
		return driver.RPlidarDriver.GrabScanDataHq(a)
	}
	return driver.GrabScanDataHqFunc(a)
}

func (driver *RPlidarDriver) AscendScanData(a ...interface{}) uint {
	if driver.AscendScanDataFunc == nil {
		return driver.RPlidarDriver.AscendScanData(a)
	}
	return driver.AscendScanDataFunc(a)
}

func (driver *RPlidarDriver) GetScanDataWithInterval(arg2 gen.Rplidar_response_measurement_node_t, arg3 *int64) uint {
	if driver.GetScanDataWithIntervalFunc == nil {
		return driver.RPlidarDriver.GetScanDataWithInterval(arg2, arg3)
	}
	return driver.GetScanDataWithIntervalFunc(arg2, arg3)
}

func (driver *RPlidarDriver) GetScanDataWithIntervalHq(arg2 gen.Rplidar_response_measurement_node_hq_t, arg3 *int64) uint {
	if driver.GetScanDataWithIntervalHqFunc == nil {
		return driver.RPlidarDriver.GetScanDataWithIntervalHq(arg2, arg3)
	}
	return driver.GetScanDataWithIntervalHqFunc(arg2, arg3)
}

func (driver *RPlidarDriver) SetX_chanDev(arg2 gen.ChannelDevice) {
	if driver.SetX_chanDevFunc == nil {
		driver.RPlidarDriver.SetX_chanDev(arg2)
	}
	driver.SetX_chanDevFunc(arg2)
}

func (driver *RPlidarDriver) GetX_chanDev() gen.ChannelDevice {
	if driver.GetX_chanDevFunc == nil {
		return driver.RPlidarDriver.GetX_chanDev()
	}
	return driver.GetX_chanDevFunc()
}
