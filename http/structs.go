package get_pair

import "time"

type Tour []struct {
	Match struct {
		ID                        int         `json:"id"`
		TournamentID              int         `json:"tournament_id"`
		State                     string      `json:"state"`
		Player1ID                 int         `json:"player1_id"`
		Player2ID                 int         `json:"player2_id"`
		Player1PrereqMatchID      interface{} `json:"player1_prereq_match_id"`
		Player2PrereqMatchID      interface{} `json:"player2_prereq_match_id"`
		Player1IsPrereqMatchLoser bool        `json:"player1_is_prereq_match_loser"`
		Player2IsPrereqMatchLoser bool        `json:"player2_is_prereq_match_loser"`
		WinnerID                  interface{} `json:"winner_id"`
		LoserID                   interface{} `json:"loser_id"`
		StartedAt                 string      `json:"started_at"`
		CreatedAt                 string      `json:"created_at"`
		UpdatedAt                 string      `json:"updated_at"`
		Identifier                string      `json:"identifier"`
		HasAttachment             bool        `json:"has_attachment"`
		Round                     int         `json:"round"`
		Player1Votes              interface{} `json:"player1_votes"`
		Player2Votes              interface{} `json:"player2_votes"`
		GroupID                   interface{} `json:"group_id"`
		AttachmentCount           interface{} `json:"attachment_count"`
		ScheduledTime             interface{} `json:"scheduled_time"`
		Location                  interface{} `json:"location"`
		UnderwayAt                interface{} `json:"underway_at"`
		Optional                  interface{} `json:"optional"`
		RushbID                   interface{} `json:"rushb_id"`
		CompletedAt               interface{} `json:"completed_at"`
		SuggestedPlayOrder        int         `json:"suggested_play_order"`
		Forfeited                 interface{} `json:"forfeited"`
		OpenGraphImageFileName    interface{} `json:"open_graph_image_file_name"`
		OpenGraphImageContentType interface{} `json:"open_graph_image_content_type"`
		OpenGraphImageFileSize    interface{} `json:"open_graph_image_file_size"`
		PrerequisiteMatchIdsCsv   string      `json:"prerequisite_match_ids_csv"`
		ScoresCsv                 string      `json:"scores_csv"`
	} `json:"match"`
}
type ParticipantData struct {
	Participant struct {
		ID                                    int           `json:"id"`
		TournamentID                          int           `json:"tournament_id"`
		Name                                  string        `json:"name"`
		Seed                                  int           `json:"seed"`
		Active                                bool          `json:"active"`
		CreatedAt                             time.Time     `json:"created_at"`
		UpdatedAt                             time.Time     `json:"updated_at"`
		InviteEmail                           interface{}   `json:"invite_email"`
		FinalRank                             interface{}   `json:"final_rank"`
		Misc                                  interface{}   `json:"misc"`
		Icon                                  interface{}   `json:"icon"`
		OnWaitingList                         bool          `json:"on_waiting_list"`
		InvitationID                          interface{}   `json:"invitation_id"`
		GroupID                               interface{}   `json:"group_id"`
		CheckedInAt                           interface{}   `json:"checked_in_at"`
		RankedMemberID                        interface{}   `json:"ranked_member_id"`
		CustomFieldResponse                   interface{}   `json:"custom_field_response"`
		Clinch                                interface{}   `json:"clinch"`
		IntegrationUids                       interface{}   `json:"integration_uids"`
		ChallongeUsername                     interface{}   `json:"challonge_username"`
		ChallongeUserID                       interface{}   `json:"challonge_user_id"`
		ChallongeEmailAddressVerified         interface{}   `json:"challonge_email_address_verified"`
		Removable                             bool          `json:"removable"`
		ParticipatableOrInvitationAttached    bool          `json:"participatable_or_invitation_attached"`
		ConfirmRemove                         bool          `json:"confirm_remove"`
		InvitationPending                     bool          `json:"invitation_pending"`
		DisplayNameWithInvitationEmailAddress string        `json:"display_name_with_invitation_email_address"`
		EmailHash                             interface{}   `json:"email_hash"`
		Username                              interface{}   `json:"username"`
		DisplayName                           string        `json:"display_name"`
		AttachedParticipatablePortraitURL     interface{}   `json:"attached_participatable_portrait_url"`
		CanCheckIn                            bool          `json:"can_check_in"`
		CheckedIn                             bool          `json:"checked_in"`
		Reactivatable                         bool          `json:"reactivatable"`
		CheckInOpen                           bool          `json:"check_in_open"`
		GroupPlayerIds                        []interface{} `json:"group_player_ids"`
		HasIrrelevantSeed                     bool          `json:"has_irrelevant_seed"`
		OrdinalSeed                           string        `json:"ordinal_seed"`
	} `json:"participant"`
}
