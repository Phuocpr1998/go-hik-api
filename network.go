package hik_api

func (device Device) EnableOnvif(isEnable bool) error {
	// create request
	soap := SOAP{
		User:          device.User,
		Password:      device.Password,
		DeviceAddress: device.DeviceAddress,
		Uri:           "/ISAPI/System/Network/Integrate",
		Method:        "PUT",
		Body: `<Integrate>
					<ONVIF>
						<enable>` + boolToString(isEnable) + `</enable>
						<certificateType/>
					</ONVIF>
				</Integrate>`,
	}

	_, err := soap.SendRequest()
	if err != nil {
		return err
	}

	return nil
}
