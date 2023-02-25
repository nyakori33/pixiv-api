package pixiv

type API struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Body    any    `json:"body,omitempty"`
}

type UserProfile struct {
	ID             int        `json:"userID"`
	Name           string     `json:"name"`
	Image          string     `json:"image"`
	ImageLarger    string     `json:"imageBig"`
	Premium        string     `json:"premium"`
	Followed       bool       `json:"isFollowed"`
	MyPixiv        bool       `json:"isMypixiv"`
	Blocking       bool       `json:"isBlocking"`
	Background     Background `json:"background"`
	AcceptRequest  any        `json:"acceptRequest,omitempty"`
	Partial        int        `json:"partial"`
	SketchLives    any        `json:"sketchLives,omitempty"`
	SketchLiveId   any        `json:"sketchLiveId,omitempty"`
	Following      int        `json:"following"`
	FollowedBack   bool       `json:"followedBack"`
	PixivCount     int        `json:"mypixivCount"`
	Comment        string     `json:"comment"`
	CommentHTML    string     `json:"commentHTML"`
	WebPage        string     `json:"webpage,omitempty"`
	CanSendMessage string     `json:"canSendMessage"`
	Official       bool       `json:"official"`
	Social         Social     `json:"social"`
	Region         Region     `json:"region"`
	Age            Privacy    `json:"age"`
	Birthday       Privacy    `json:"birthDay"`
	Gender         Privacy    `json:"gender"`
	Job            Privacy    `json:"job"`
	Workspace      Workspace  `json:"workspace"`
	Group          []Group    `json:"group"`
}

type UserProfileSimple struct {
	ID            int        `json:"userID"`
	Name          string     `json:"name"`
	Image         string     `json:"image"`
	ImageLarger   string     `json:"imageBig"`
	Premium       string     `json:"premium"`
	Followed      bool       `json:"isFollowed"`
	MyPixiv       bool       `json:"isMypixiv"`
	Blocking      bool       `json:"isBlocking"`
	Background    Background `json:"background"`
	AcceptRequest any        `json:"acceptRequest,omitempty"`
	Partial       int        `json:"partial"`
	SketchLives   any        `json:"sketchLives,omitempty"`
	SketchLiveId  any        `json:"sketchLiveId,omitempty"`
}

type Background struct {
	Repeat  any    `json:"repeat,omitempty"`
	Color   any    `json:"color,omitempty"`
	URL     string `json:"url"`
	Private bool   `json:"isPrivate"`
}

type Social struct {
	Twitter   SocialURL `json:"twitter,omitempty"`
	Facebook  SocialURL `json:"facebook,omitempty"`
	Instagram SocialURL `json:"instagram,omitempty"`
	Pawoo     SocialURL `json:"pawoo,omitempty"`
	Tumblr    SocialURL `json:"tumblr,omitempty"`
	Circlems  SocialURL `json:"circlems,omitempty"`
}

type SocialURL struct {
	URL string `json:"url"`
}

type Region struct {
	Name         string `json:"name,omitempty"`
	Code         string `json:"region,omitempty"`
	Prefecture   int    `json:"prefecture,omitempty"`
	PrivacyLevel int    `json:"privacyLevel,omitempty"`
}

type Privacy struct {
	Name         string `json:"name,omitempty"`
	PrivacyLevel int    `json:"privacyLevel,omitempty"`
}

type Workspace struct {
	PC          string `json:"userWorkspacePc"`
	Monitor     string `json:"userWorkspaceMonitor"`
	Tool        string `json:"userWorkspaceTool"`
	Scanner     string `json:"userWorkspaceScanner"`
	Tablet      string `json:"userWorkspaceTablet"`
	Mouse       string `json:"userWorkspaceMouse"`
	Printer     string `json:"userWorkspacePrinter"`
	Desktop     string `json:"userWorkspaceDesktop"`
	Music       string `json:"userWorkspaceMusic"`
	Desk        string `json:"userWorkspaceDesk"`
	Chair       string `json:"userWorkspaceChair"`
	Comment     string `json:"userWorkspaceComment"`
	Image       string `json:"wsUrl"`
	ImageLarger string `json:"wsBigUrl"`
}

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"title"`
	Icon string `json:"iconUrl"`
}
