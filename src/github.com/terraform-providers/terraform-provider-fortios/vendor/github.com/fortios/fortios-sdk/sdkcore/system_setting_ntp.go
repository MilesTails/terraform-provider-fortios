package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// "strconv"
)

// JSONSystemSettingNTP contains the parameters for Create and Update API function
type JSONSystemSettingNTP struct {
	Type      string        `json:"type"`
	Ntpserver NTPMultValues `json:"ntpserver"`
}

// JSONCreateSystemSettingNTPOutput contains the output results for Create API function
type JSONCreateSystemSettingNTPOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateSystemSettingNTPOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateSystemSettingNTPOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// NTPMultValue contains the output results for Read API function
type NTPMultValue struct {
	Server string `json:"server"`
}

// NTPMultValues contains the output results for Read API function
type NTPMultValues []NTPMultValue

// CreateSystemSettingNTP API operation for FortiOS
func (c *FortiSDKClient) CreateSystemSettingNTP(params *JSONSystemSettingNTP) (output *JSONCreateSystemSettingNTPOutput, err error) {
	// HTTPMethod := "POST"
	// path := "/api/v2/cmdb/system/ntp"
	// output = &JSONCreateSystemSettingNTPOutput{}
	// locJSON, err := json.Marshal(params)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// bytes := bytes.NewBuffer(locJSON)
	// req := c.NewRequest(HTTPMethod, path, nil, bytes)
	// err = req.Send()

	// body, err := ioutil.ReadAll(req.HTTPResponse.Body)

	// var result map[string]interface{}
	// json.Unmarshal([]byte(string(body)), &result)

	// req.HTTPResponse.Body.Close()

	// if result != nil {
	// 	if result["vdom"] != nil {
	// 		output.Vdom = result["vdom"].(string)
	// 	}
	// 	if result["mkey"] != nil {
	// 		output.Mkey = result["mkey"].(string)
	// 	}
	// 	if result["status"] != nil {
	//		if result["status"] != "success" {
	//			err = fmt.Errorf("cannot get the right response")
	//			return
	//		}
	// 		output.Status = result["status"].(string)
	// 	} else {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}
	// 	if result["http_status"] != nil {
	// 		output.HTTPStatus = result["http_status"].(float64)
	// 	}
	// } else {
	// 	err = fmt.Errorf("cannot get the right response")
	// 	return
	// }

	return
}

// UpdateSystemSettingNTP API operation for FortiOS configures Network Time Protocol (NTP) server.
// Returns the execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ntp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSettingNTP(params *JSONSystemSettingNTP, mkey string) (output *JSONUpdateSystemSettingNTPOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ntp"
	// path += "/" + mkey
	output = &JSONUpdateSystemSettingNTPOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		}
		if result["status"] != nil {
			if result["status"] != "success" {
				err = fmt.Errorf("cannot get the right response")
				return
			}
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get the right response")
			return
		}
		if result["http_status"] != nil {
			output.HTTPStatus = result["http_status"].(float64)
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// DeleteSystemSettingNTP API operation for FortiOS
func (c *FortiSDKClient) DeleteSystemSettingNTP(mkey string) (err error) {
	// HTTPMethod := "DELETE"
	// path := "/api/v2/cmdb/system/ntp"
	// // path += "/" + mkey

	// req := c.NewRequest(HTTPMethod, path, nil, nil)
	// err = req.Send()

	// body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	// log.Printf("FOS-fortios response: %s", string(body))

	// var result map[string]interface{}
	// json.Unmarshal([]byte(string(body)), &result)

	// req.HTTPResponse.Body.Close()

	// if result != nil {
	// 	if result["status"] == nil {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}
	//
	//	if result["status"] != "success" {
	//		err = fmt.Errorf("cannot get the right response")
	//		return
	//	}
	//
	// } else {
	// 	err = fmt.Errorf("cannot get the right response")
	// 	return
	// }

	return
}

// ReadSystemSettingNTP API operation for FortiOS gets the Network Time Protocol (NTP) server.
// Returns the requested NTP setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ntp chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSettingNTP(mkey string) (output *JSONSystemSettingNTP, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ntp"
	// path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("FOS-fortios reading response: %s", string(body))

	output = &JSONSystemSettingNTP{}
	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}
		
		if result["status"] != "success" {
			err = fmt.Errorf("cannot get the right response")
			return
		}

		mapTmp := (result["results"].(map[string]interface{}))

		if mapTmp == nil {
			return
		}

		if mapTmp["type"] != nil {
			output.Type = mapTmp["type"].(string)
		}
		if mapTmp["ntpserver"] != nil {
			member := mapTmp["ntpserver"].([]interface{})

			var members []NTPMultValue
			for _, v := range member {
				c := v.(map[string]interface{})

				members = append(members,
					NTPMultValue{
						Server: c["server"].(string),
					})
			}
			output.Ntpserver = members
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
