package auther

// func SendOperateEvent(req, resp interface{}, hd *event.Header, od *event.OperateEventData) error {
// 	if od == nil {
// 		return nil
// 	}

// 	reqd, err := json.Marshal(req)
// 	if err != nil {
// 		return fmt.Errorf("marshal req for event error, %s", err)
// 	}

// 	respd, err := json.Marshal(resp)
// 	if err != nil {
// 		return fmt.Errorf("marshal resp for event error, %s", err)
// 	}

// 	od.Request = string(reqd)
// 	od.Response = string(respd)
// 	od.Cost = ftime.Now().Timestamp() - hd.Time
// 	oe, err := event.NewProtoOperateEvent(od)
// 	if err != nil {
// 		return fmt.Errorf("new operate event error, %s", err)
// 	}
// 	oe.Header = hd

// 	if err := bus.Pub(oe); err != nil {
// 		return fmt.Errorf("pub audit log error, %s", err)
// 	}

// 	return nil
// }

// func newOperateEventData(e *httpb.Entry, tk *token.Token) *event.OperateEventData {
// 	od := &event.OperateEventData{
// 		Action:       e.GetLableValue("action"),
// 		FeaturePath:  e.Path,
// 		ResourceType: e.Resource,
// 		ServiceName:  version.ServiceName,
// 	}

// 	if tk != nil {
// 		// 补充审计的用户信息
// 		od.Account = tk.Account
// 		od.UserDomain = tk.Domain
// 		od.Session = tk.SessionId
// 		od.UserType = tk.UserType.String()
// 	}
// 	return od
// }

// func newEventHeaderFromHTTP(r *http.Request) *event.Header {
// 	hd := event.NewHeader()
// 	hd.IpAddress = request.GetRemoteIP(r)
// 	hd.UserAgent = r.UserAgent()
// 	hd.RequestId = r.Header.Get("RequestIdHeader")
// 	hd.Source = version.ServiceName
// 	hd.Meta["host"], _ = os.Hostname()
// 	return hd
// }
