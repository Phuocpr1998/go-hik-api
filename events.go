package hik_api

import "github.com/golang/glog"

func (device Device) EnableMotionEvent(isEnable bool) error {
	// create request
	soap := SOAP{
		User:          device.User,
		Password:      device.Password,
		DeviceAddress: device.DeviceAddress,
		Method:        "PUT",
		Uri:           `/ISAPI/System/Video/inputs/channels/1/motionDetection`,
		Body: `<MotionDetection xmlns="http://www.hikvision.com/ver20/XMLSchema" version="2.0">
					<enabled>` + boolToString(isEnable) + `</enabled>
					<enableHighlight>true</enableHighlight>
					<regionType>grid</regionType>
					<Grid>
						<rowGranularity>18</rowGranularity>
						<columnGranularity>22</columnGranularity>
					</Grid>
					<MotionDetectionLayout xmlns="http://www.hikvision.com/ver20/XMLSchema" version="2.0">
						<sensitivityLevel>60</sensitivityLevel>
						<layout>
							<gridMap/>
						</layout>
					</MotionDetectionLayout>
				</MotionDetection>`,
	}

	_, err := soap.SendRequest()
	if err != nil {
		glog.Info(err)
		return err
	}

	return nil
}
