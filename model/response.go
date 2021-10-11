package model

// LoginInfo 登录后系统返回的信息
// SessionId 登录成功后的 cookie
// LoginerId 登录成功后的 登录者的 id
type LoginInfo struct {
	Code string `json:"code"`
	Data struct {
		Activate     bool   `json:"activate"`
		Md5Uid       string `json:"md5Uid"`
		SessionId    string `json:"sessionId"`
		NeedComplete bool   `json:"needComplete"`
		LoginerId    int    `json:"loginerId"`
		Phone        string `json:"phone"`
	} `json:"data"`
	Msg  string `json:"msg"`
	Mstv struct {
		T int    `json:"t"`
		M string `json:"m"`
		S string `json:"s"`
	} `json:"mstv"`
}

type PersonInfo struct {
	Code string `json:"code"`
	Data struct {
		IntroductionCount int         `json:"introductionCount"`
		QuestionCount     int         `json:"questionCount"`
		Specialty         string      `json:"specialty"`
		FollowCount       int         `json:"followCount"`
		CareerTalkCount   int         `json:"careerTalkCount"`
		Schoolyear        string      `json:"schoolyear"`
		SexType           int         `json:"sexType"`
		HeadPic           string      `json:"headPic"`
		AnswerSheetCount  int         `json:"answerSheetCount"`
		Mbtitest          bool        `json:"mbtitest"`
		Faculty           string      `json:"faculty"`
		Loginer           string      `json:"loginer"`
		PicUrl            string      `json:"picUrl"`
		FanCount          int         `json:"fanCount"`
		MovementCount     int         `json:"movementCount"`
		AnswerCount       int         `json:"answerCount"`
		School            string      `json:"school"`
		Klass             string      `json:"klass"`
		SchoolId          int         `json:"schoolId"`
		Activate          bool        `json:"activate"`
		IsTeacher         bool        `json:"isTeacher"`
		FindProfessionId  interface{} `json:"findProfessionId"`
		CompletenessSum   int         `json:"completenessSum"`
	} `json:"data"`
	Msg  string `json:"msg"`
	Mstv struct {
		T int    `json:"t"`
		M string `json:"m"`
		S string `json:"s"`
	} `json:"mstv"`
}

type TraineeInfo struct {
	Code string `json:"code"`
	Data struct {
		ClockVo struct {
			DateName    interface{} `json:"dateName"`
			EndDate     string      `json:"endDate"`
			ModuleName  string      `json:"moduleName"`
			PlanName    string      `json:"planName"`
			ProjectName string      `json:"projectName"`
			StartDate   string      `json:"startDate"`
			TraineeId   int         `json:"traineeId"`
		} `json:"clockVo"`
		HasMore   bool   `json:"hasMore"`
		ClockType string `json:"clockType"`
	} `json:"data"`
	Msg  string `json:"msg"`
	Mstv struct {
		T int    `json:"t"`
		M string `json:"m"`
		S string `json:"s"`
	} `json:"mstv"`
}

type PositionInfo struct {
	Code string `json:"code"`
	Data struct {
		ClockRuleType int `json:"clockRuleType"`
		ClockInfo     struct {
			Date          string `json:"date"`
			InAddress     string `json:"inAddress"`
			InStatus      int    `json:"inStatus"`
			InStatusDesc  string `json:"inStatusDesc"`
			InTime        string `json:"inTime"`
			OutAddress    string `json:"outAddress"`
			OutStatus     int    `json:"outStatus"`
			OutStatusDesc string `json:"outStatusDesc"`
			OutTime       string `json:"outTime"`
			Status        int    `json:"status"`
			Week          string `json:"week"`
		} `json:"clockInfo"`
		PostInfo struct {
			Address  string  `json:"address"`
			Clock    int     `json:"clock"`
			Compare  int     `json:"compare"`
			Distance int     `json:"distance"`
			Lat      float64 `json:"lat"`
			Lng      float64 `json:"lng"`
			State    int     `json:"state"`
		} `json:"postInfo"`
	} `json:"data"`
	Msg  string `json:"msg"`
	Mstv struct {
		T int    `json:"t"`
		M string `json:"m"`
		S string `json:"s"`
	} `json:"mstv"`
}

type SignStatusInfo struct {
	Code string `json:"code"`
	Data struct {
		ClockRuleType int `json:"clockRuleType"`
		ClockInfo     struct {
			Date          string `json:"date"`
			InAddress     string `json:"inAddress"`
			InStatus      int    `json:"inStatus"`
			InStatusDesc  string `json:"inStatusDesc"`
			InTime        string `json:"inTime"`
			OutAddress    string `json:"outAddress"`
			OutStatus     int    `json:"outStatus"`
			OutStatusDesc string `json:"outStatusDesc"`
			OutTime       string `json:"outTime"`
			Status        int    `json:"status"`
			Week          string `json:"week"`
		} `json:"clockInfo"`
		PostInfo struct {
			Address  string  `json:"address"`
			Clock    int     `json:"clock"`
			Compare  int     `json:"compare"`
			Distance int     `json:"distance"`
			Lat      float64 `json:"lat"`
			Lng      float64 `json:"lng"`
			State    int     `json:"state"`
		} `json:"postInfo"`
	} `json:"data"`
	Msg  string `json:"msg"`
	Mstv struct {
		T int    `json:"t"`
		M string `json:"m"`
		S string `json:"s"`
	} `json:"mstv"`
}
