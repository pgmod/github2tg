package handlers

type GitHubWebhookPayload struct {
	Ref          string       `json:"ref"`
	Before       string       `json:"before"`
	After        string       `json:"after"`
	Repository   Repository   `json:"repository"`
	Pusher       Pusher       `json:"pusher"`
	Organization Organization `json:"organization"`
	Sender       Sender       `json:"sender"`
	Created      bool         `json:"created"`
	Deleted      bool         `json:"deleted"`
	Forced       bool         `json:"forced"`
	BaseRef      interface{}  `json:"base_ref"`
	Compare      string       `json:"compare"`
	Commits      []Commit     `json:"commits"`
	HeadCommit   Commit       `json:"head_commit"`
	Action       string       `json:"action"`
	Release      Release      `json:"release"`
	Issue        Issue        `json:"issue"`
	Comment      Comment      `json:"comment"`
}

type Comment struct {
	URL                   string      `json:"url"`
	HTMLURL               string      `json:"html_url"`
	IssueURL              string      `json:"issue_url"`
	ID                    int64       `json:"id"`
	NodeID                string      `json:"node_id"`
	User                  Sender      `json:"user"`
	CreatedAt             string      `json:"created_at"`
	UpdatedAt             string      `json:"updated_at"`
	AuthorAssociation     string      `json:"author_association"`
	Body                  string      `json:"body"`
	Reactions             Reactions   `json:"reactions"`
	PerformedViaGithubApp interface{} `json:"performed_via_github_app"`
}

type Reactions struct {
	URL        string `json:"url"`
	TotalCount int64  `json:"total_count"`
	The1       int64  `json:"+1"`
	Reactions1 int64  `json:"-1"`
	Laugh      int64  `json:"laugh"`
	Hooray     int64  `json:"hooray"`
	Confused   int64  `json:"confused"`
	Heart      int64  `json:"heart"`
	Rocket     int64  `json:"rocket"`
	Eyes       int64  `json:"eyes"`
}

type Sender struct {
	Login             string      `json:"login"`
	ID                int64       `json:"id"`
	NodeID            string      `json:"node_id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	UserViewType      string      `json:"user_view_type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              *string     `json:"name,omitempty"`
	Email             interface{} `json:"email"`
}

type Commit struct {
	ID        string   `json:"id"`
	TreeID    string   `json:"tree_id"`
	Distinct  bool     `json:"distinct"`
	Message   string   `json:"message"`
	Timestamp string   `json:"timestamp"`
	URL       string   `json:"url"`
	Author    Author   `json:"author"`
	Committer Author   `json:"committer"`
	Added     []string `json:"added"`
	Removed   []string `json:"removed"`
	Modified  []string `json:"modified"`
}

type Author struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Issue struct {
	URL                   string        `json:"url"`
	RepositoryURL         string        `json:"repository_url"`
	LabelsURL             string        `json:"labels_url"`
	CommentsURL           string        `json:"comments_url"`
	EventsURL             string        `json:"events_url"`
	HTMLURL               string        `json:"html_url"`
	ID                    int64         `json:"id"`
	NodeID                string        `json:"node_id"`
	Number                int64         `json:"number"`
	Title                 string        `json:"title"`
	User                  Sender        `json:"user"`
	Labels                []interface{} `json:"labels"`
	State                 string        `json:"state"`
	Locked                bool          `json:"locked"`
	Assignee              interface{}   `json:"assignee"`
	Assignees             []interface{} `json:"assignees"`
	Milestone             interface{}   `json:"milestone"`
	Comments              int64         `json:"comments"`
	CreatedAt             string        `json:"created_at"`
	UpdatedAt             string        `json:"updated_at"`
	ClosedAt              string        `json:"closed_at"`
	AuthorAssociation     string        `json:"author_association"`
	ActiveLockReason      interface{}   `json:"active_lock_reason"`
	Body                  string        `json:"body"`
	Reactions             Reactions     `json:"reactions"`
	TimelineURL           string        `json:"timeline_url"`
	PerformedViaGithubApp interface{}   `json:"performed_via_github_app"`
	StateReason           string        `json:"state_reason"`
}

type Organization struct {
	Login            string `json:"login"`
	ID               int64  `json:"id"`
	NodeID           string `json:"node_id"`
	URL              string `json:"url"`
	ReposURL         string `json:"repos_url"`
	EventsURL        string `json:"events_url"`
	HooksURL         string `json:"hooks_url"`
	IssuesURL        string `json:"issues_url"`
	MembersURL       string `json:"members_url"`
	PublicMembersURL string `json:"public_members_url"`
	AvatarURL        string `json:"avatar_url"`
	Description      string `json:"description"`
}

type Pusher struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Release struct {
	URL             string        `json:"url"`
	AssetsURL       string        `json:"assets_url"`
	UploadURL       string        `json:"upload_url"`
	HTMLURL         string        `json:"html_url"`
	ID              int64         `json:"id"`
	Author          Sender        `json:"author"`
	NodeID          string        `json:"node_id"`
	TagName         string        `json:"tag_name"`
	TargetCommitish string        `json:"target_commitish"`
	Name            string        `json:"name"`
	Draft           bool          `json:"draft"`
	Prerelease      bool          `json:"prerelease"`
	CreatedAt       string        `json:"created_at"`
	PublishedAt     string        `json:"published_at"`
	Assets          []interface{} `json:"assets"`
	TarballURL      string        `json:"tarball_url"`
	ZipballURL      string        `json:"zipball_url"`
	Body            string        `json:"body"`
}

type Repository struct {
	ID                       int64            `json:"id"`
	NodeID                   string           `json:"node_id"`
	Name                     string           `json:"name"`
	FullName                 string           `json:"full_name"`
	Private                  bool             `json:"private"`
	Owner                    Sender           `json:"owner"`
	HTMLURL                  string           `json:"html_url"`
	Description              interface{}      `json:"description"`
	Fork                     bool             `json:"fork"`
	URL                      string           `json:"url"`
	ForksURL                 string           `json:"forks_url"`
	KeysURL                  string           `json:"keys_url"`
	CollaboratorsURL         string           `json:"collaborators_url"`
	TeamsURL                 string           `json:"teams_url"`
	HooksURL                 string           `json:"hooks_url"`
	IssueEventsURL           string           `json:"issue_events_url"`
	EventsURL                string           `json:"events_url"`
	AssigneesURL             string           `json:"assignees_url"`
	BranchesURL              string           `json:"branches_url"`
	TagsURL                  string           `json:"tags_url"`
	BlobsURL                 string           `json:"blobs_url"`
	GitTagsURL               string           `json:"git_tags_url"`
	GitRefsURL               string           `json:"git_refs_url"`
	TreesURL                 string           `json:"trees_url"`
	StatusesURL              string           `json:"statuses_url"`
	LanguagesURL             string           `json:"languages_url"`
	StargazersURL            string           `json:"stargazers_url"`
	ContributorsURL          string           `json:"contributors_url"`
	SubscribersURL           string           `json:"subscribers_url"`
	SubscriptionURL          string           `json:"subscription_url"`
	CommitsURL               string           `json:"commits_url"`
	GitCommitsURL            string           `json:"git_commits_url"`
	CommentsURL              string           `json:"comments_url"`
	IssueCommentURL          string           `json:"issue_comment_url"`
	ContentsURL              string           `json:"contents_url"`
	CompareURL               string           `json:"compare_url"`
	MergesURL                string           `json:"merges_url"`
	ArchiveURL               string           `json:"archive_url"`
	DownloadsURL             string           `json:"downloads_url"`
	IssuesURL                string           `json:"issues_url"`
	PullsURL                 string           `json:"pulls_url"`
	MilestonesURL            string           `json:"milestones_url"`
	NotificationsURL         string           `json:"notifications_url"`
	LabelsURL                string           `json:"labels_url"`
	ReleasesURL              string           `json:"releases_url"`
	DeploymentsURL           string           `json:"deployments_url"`
	UpdatedAt                string           `json:"updated_at"`
	GitURL                   string           `json:"git_url"`
	SSHURL                   string           `json:"ssh_url"`
	CloneURL                 string           `json:"clone_url"`
	SvnURL                   string           `json:"svn_url"`
	Homepage                 interface{}      `json:"homepage"`
	Size                     int64            `json:"size"`
	StargazersCount          int64            `json:"stargazers_count"`
	WatchersCount            int64            `json:"watchers_count"`
	Language                 string           `json:"language"`
	HasIssues                bool             `json:"has_issues"`
	HasProjects              bool             `json:"has_projects"`
	HasDownloads             bool             `json:"has_downloads"`
	HasWiki                  bool             `json:"has_wiki"`
	HasPages                 bool             `json:"has_pages"`
	HasDiscussions           bool             `json:"has_discussions"`
	ForksCount               int64            `json:"forks_count"`
	MirrorURL                interface{}      `json:"mirror_url"`
	Archived                 bool             `json:"archived"`
	Disabled                 bool             `json:"disabled"`
	OpenIssuesCount          int64            `json:"open_issues_count"`
	License                  interface{}      `json:"license"`
	AllowForking             bool             `json:"allow_forking"`
	IsTemplate               bool             `json:"is_template"`
	WebCommitSignoffRequired bool             `json:"web_commit_signoff_required"`
	Topics                   []interface{}    `json:"topics"`
	Visibility               string           `json:"visibility"`
	Forks                    int64            `json:"forks"`
	OpenIssues               int64            `json:"open_issues"`
	Watchers                 int64            `json:"watchers"`
	DefaultBranch            string           `json:"default_branch"`
	Stargazers               *int64           `json:"stargazers,omitempty"`
	MasterBranch             *string          `json:"master_branch,omitempty"`
	Organization             *string          `json:"organization,omitempty"`
	CustomProperties         CustomProperties `json:"custom_properties"`
}

type CustomProperties struct {
}
